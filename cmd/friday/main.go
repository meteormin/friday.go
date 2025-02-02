package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meteormin/friday.go/internal/adapter/http/handler"
	"github.com/meteormin/friday.go/internal/adapter/repo"
	"github.com/meteormin/friday.go/internal/app"
	"github.com/meteormin/friday.go/internal/infra"
	"github.com/meteormin/friday.go/internal/infra/db"
	"github.com/meteormin/friday.go/internal/infra/http"
	"github.com/meteormin/friday.go/internal/infra/http/middleware"
	"github.com/meteormin/friday.go/pkg/config"
	"github.com/meteormin/friday.go/pkg/logger"
	"strconv"
)

const (
	AppName    = "Friday.go"
	AppVersion = "0.0.1"
)

func init() {
	cfg := config.LoadWithViper("config.yml", config.App{
		Name:    AppName,
		Version: AppVersion,
	})

	infra.SetConfig(cfg)

	l := logger.NewZapLogger(cfg.Logging)
	if err := db.New(cfg.Database); err != nil {
		l.Fatal(err)
	}

	http.NewFiber(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		AppName:       AppName + " v" + AppVersion,
	})
}

func AuthHandler() http.AddRouteFunc {
	userRepo := repo.NewUserRepository(infra.GetDB())
	userCommand := app.NewAccountCommandService(userRepo)
	userQuery := app.NewAccountQueryService(userRepo)
	return handler.NewAuthHandler(userCommand, userQuery)
}

func main() {

	cfg := infra.GetConfig()
	l := infra.GetLogger()

	http.Middleware(middleware.NewCommon, "/api")
	http.Route("/api", AuthHandler())

	if cfg.Server.Port <= 0 {
		cfg.Server.Port = 8080
	}

	l.Fatal(http.Fiber().Listen(":" + strconv.Itoa(cfg.Server.Port)))
}
