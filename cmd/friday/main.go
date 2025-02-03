package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/meteormin/friday.go/internal/adapter/repo"
	"github.com/meteormin/friday.go/internal/adapter/rest"
	"github.com/meteormin/friday.go/internal/app"
	"github.com/meteormin/friday.go/internal/infra"
	"github.com/meteormin/friday.go/internal/infra/db"
	"github.com/meteormin/friday.go/internal/infra/http"
	"github.com/meteormin/friday.go/internal/infra/http/middleware"
	"github.com/meteormin/friday.go/internal/infra/task"
	"github.com/meteormin/friday.go/pkg/config"
	"github.com/meteormin/friday.go/pkg/logger"
	"github.com/meteormin/friday.go/pkg/scheduler"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	appName    = "Friday.go"
	appVersion = "0.0.1"
)

func init() {
	cfg := config.LoadWithViper("config.yml", config.App{
		Name:    appName,
		Version: appVersion,
	})

	infra.SetConfig(cfg)

	l := logger.NewZapLogger(cfg.Logging)
	l.Info("Initializing application...")

	if err := db.New(cfg.Database); err != nil {
		l.Fatal(err)
	}

	l.Info("Database connection established.")

	loc, err := time.LoadLocation(cfg.TZ)
	if err != nil {
		l.Fatal(err)
	}

	cfg.Scheduler.Location = loc

	jobRepo := task.NewJobRepository(infra.GetDB())
	cfg.Scheduler.Monitor = jobRepo
	cfg.Scheduler.MonitorStatus = jobRepo

	err = scheduler.New(cfg.Scheduler)
	if err != nil {
		l.Fatal(err)
	}

	http.NewFiber(fiber.Config{
		CaseSensitive:     true,
		AppName:           appName + " v" + appVersion,
		ErrorHandler:      http.NewErrorHandler(),
		EnablePrintRoutes: true,
	})
}

func authHandler() http.AddRouteFunc {
	userRepo := repo.NewUserRepository(infra.GetDB())
	userCommand := app.NewAccountCommandService(userRepo)
	userQuery := app.NewAccountQueryService(userRepo)
	return rest.NewAuthHandler(userCommand, userQuery)
}

func userHandler() http.AddRouteFunc {
	userRepo := repo.NewUserRepository(infra.GetDB())
	userCommand := app.NewAccountCommandService(userRepo)
	userQuery := app.NewAccountQueryService(userRepo)
	return rest.NewUserHandler(userCommand, userQuery)
}

func main() {

	cfg := infra.GetConfig()
	l := infra.GetLogger()

	http.Middleware(middleware.NewCommon, "/api")
	http.Route("/api", func(router fiber.Router) {
		auth := authHandler()
		user := userHandler()

		auth(router)
		middleware.NewJwtGuard(router)
		user(router)
	})

	if cfg.Server.Port <= 0 {
		cfg.Server.Port = 8080
	}

	fiberApp := http.Fiber()
	taskScheduler := infra.GetScheduler()

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

	err := infra.GetScheduler().Shutdown()
	if err != nil {
		l.Fatal(err)
	}

	fmt.Println("Fiber was successful shutdown.")
}
