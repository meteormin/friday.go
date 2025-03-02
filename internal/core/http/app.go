package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meteormin/friday.go/internal/core/config"
)

var app *fiber.App

func NewApp(cfg config.App) *fiber.App {
	app = fiber.New(fiber.Config{
		AppName:           cfg.Name + " v" + cfg.Version,
		CaseSensitive:     cfg.CaseSensitive,
		EnablePrintRoutes: cfg.EnablePrintRouts,
		ErrorHandler:      NewErrorHandler(),
	})
	return app
}

func App() *fiber.App {
	return app
}

type AddRouteFunc func(router fiber.Router)
