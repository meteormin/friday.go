package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // swagger handler
	"github.com/meteormin/friday.go/api"
	"github.com/meteormin/friday.go/internal/adapter/repo"
	"github.com/meteormin/friday.go/internal/adapter/rest"
	"github.com/meteormin/friday.go/internal/app"
	"github.com/meteormin/friday.go/internal/boot"
	"github.com/meteormin/friday.go/internal/core"
	"github.com/meteormin/friday.go/internal/core/http"
	"github.com/meteormin/friday.go/internal/core/http/middleware"
	"github.com/meteormin/friday.go/internal/core/task"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

// init is a function that is called when the program starts.
func init() {
	boot.Boot()
}

// apiInfo sets the API information based on the given port.
//
// It takes a string parameter 'port' and does not return anything.
func apiInfo(host string, version string, port string) {

	if host == "" {
		host = fmt.Sprintf("%s:%s", "localhost", port)
	}

	schemaHost := strings.Split(host, "://")
	if len(schemaHost) > 1 {
		host = schemaHost[1]
	}

	api.SwaggerInfo.Title = "Friday.go API"
	api.SwaggerInfo.Version = version
	api.SwaggerInfo.Schemes = []string{"http", "https"}
	api.SwaggerInfo.Host = host
	api.SwaggerInfo.BasePath = "/"
	api.SwaggerInfo.SwaggerTemplate = api.OpenAPITemplate
}

// Friday.go API
// @title Friday.go API
// @version {{.Version}}
// @description Friday.go API
// @contact.name meteormin
// @contact.url https://github.com/meteormin/friday.go
// @contact.email miniyu97@gmail.com
// @schemes http https
// @host {{.Host}}
// @BasePath /
// @tag.name api
// @tag.name auth
// @tag.name users
// @tag.name upload-file
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := core.GetConfig()

	l := core.GetLogger()

	if cfg.Server.Port <= 0 {
		cfg.Server.Port = 8080
	}

	fiberApp := http.Fiber()
	taskScheduler := core.GetScheduler()

	// set swagger info
	apiInfo(cfg.Server.Host, cfg.App.Version, strconv.Itoa(cfg.Server.Port))

	// set routes
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

	<-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	err := fiberApp.Shutdown()
	if err != nil {
		l.Error("Graceful shutdown failed", err)
	} else {
		l.Info("Fiber was successful shutdown.")
	}

	fmt.Println("Running cleanup tasks...")

	err = core.GetScheduler().Shutdown()
	if err != nil {
		l.Error("Scheduler shutdown failed", err)
	} else {
		l.Info("Scheduler was successful shutdown.")
	}

	err = core.GetStorage().Close()
	if err != nil {
		l.Error("Storage shutdown failed", err)
	} else {
		l.Error("Storage was successful shutdown.")
	}
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

func siteHandler() http.AddRouteFunc {
	siteRepo := repo.NewSiteRepository(core.GetDB())
	siteCommand := app.NewSiteCommandService(siteRepo)
	siteQuery := app.NewSiteQueryService(siteRepo)
	return rest.NewSiteHandler(siteCommand, siteQuery)
}

func postHandler() http.AddRouteFunc {
	postRepo := repo.NewPostRepository(core.GetDB())
	postCommand := app.NewPostCommandService(postRepo)
	postQuery := app.NewPostQueryService(postRepo)
	return rest.NewPostHandler(postCommand, postQuery)
}

func routes() {
	http.Route("/api-docs", func(router fiber.Router) {
		router.Get("/swagger/*", swagger.HandlerDefault)
	})

	http.Middleware(middleware.NewCommon, "/api")
	http.Route("/api", func(router fiber.Router) {
		tasks := task.Handler()
		auth := authHandler()
		user := userHandler()
		uploadFile := uploadFileHandler()
		sites := siteHandler()
		posts := postHandler()

		tasks(router)
		auth(router)
		middleware.NewJwtGuard(router)
		user(router)
		uploadFile(router)
		sites(router)
		posts(router)

	})
}
