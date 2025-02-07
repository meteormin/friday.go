package http

import "github.com/gofiber/fiber/v2"

var app *fiber.App

type AddRouteFunc func(router fiber.Router)

type AddMiddlewareFunc func(router fiber.Router)

func NewFiber(cfg fiber.Config) {
	app = fiber.New(cfg)
}

func Fiber() *fiber.App {
	return app
}

func Route(prefix string, addRoute AddRouteFunc) {
	addRoute(app.Group(prefix))
}

func Middleware(addMiddleware AddMiddlewareFunc, scopes ...string) {
	if len(scopes) == 0 {
		addMiddleware(app)
		return
	}

	for _, scope := range scopes {
		addMiddleware(app.Group(scope))
	}
}
