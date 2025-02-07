package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/meteormin/friday.go/internal/adapter/repo"
	"github.com/meteormin/friday.go/internal/adapter/rest"
	"github.com/meteormin/friday.go/internal/app"
	"github.com/meteormin/friday.go/internal/boot"
	"github.com/meteormin/friday.go/internal/core"
	"github.com/meteormin/friday.go/internal/core/http"
	"github.com/meteormin/friday.go/internal/core/http/middleware"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func init() {
	boot.Boot()
}

func main() {

	cfg := core.GetConfig()
	l := core.GetLogger()

	if cfg.Server.Port <= 0 {
		cfg.Server.Port = 8080
	}

	fiberApp := http.Fiber()
	taskScheduler := core.GetScheduler()

	routes()

	// Listen from a different goroutine
	go func() {
		taskScheduler.Start()
		if err := fiberApp.Listen(":" + strconv.Itoa(cfg.Server.Port)); err != nil {
			l.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = fiberApp.Shutdown()

	fmt.Println("Running cleanup tasks...")

	err := core.GetScheduler().Shutdown()
	if err != nil {
		l.Fatal(err)
	}

	err = core.GetStorage().Close()
	if err != nil {
		l.Fatal(err)
	}

	fmt.Println("Fiber was successful shutdown.")
}

func authHandler() http.AddRouteFunc {
	userRepo := repo.NewUserRepository(core.GetDB())
	userCommand := app.NewUserCommandService(userRepo)
	userQuery := app.NewUserQueryService(userRepo)
	return rest.NewAuthHandler(userCommand, userQuery)
}

func userHandler() http.AddRouteFunc {
	userRepo := repo.NewUserRepository(core.GetDB())
	userCommand := app.NewUserCommandService(userRepo)
	userQuery := app.NewUserQueryService(userRepo)
	return rest.NewUserHandler(userCommand, userQuery)
}

func uploadFileHandler() http.AddRouteFunc {
	fileRepo := repo.NewFileRepository(core.GetDB(), "uploads")
	fileCommand := app.NewUploadFileService(fileRepo)
	return rest.NewUploadFileHandler(fileCommand)
}

func routes() {
	http.Middleware(middleware.NewCommon, "/api")
	http.Route("/api", func(router fiber.Router) {
		auth := authHandler()
		user := userHandler()
		uploadFile := uploadFileHandler()

		auth(router)
		middleware.NewJwtGuard(router)
		user(router)
		uploadFile(router)
	})
}
