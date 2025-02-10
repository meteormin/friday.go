package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/envvar"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/meteormin/friday.go/internal/core"
	"github.com/meteormin/friday.go/internal/core/config"
	"github.com/meteormin/friday.go/internal/core/http"
	"time"
)
import "github.com/gofiber/fiber/v2/middleware/cors"
import "github.com/gofiber/jwt/v3"

func NewCommon(router fiber.Router) {
	router.Use(logger.New())
	router.Use(requestid.New())
	router.Use(cors.New())
	router.Use(healthcheck.New())

	if core.GetConfig().Env != config.Release {
		router.Use("/expose/envvars", envvar.New())
		router.Use("/metrics", monitor.New())
		router.Use("/routes", func(ctx *fiber.Ctx) error {
			return ctx.JSON(http.Fiber().GetRoutes())
		})
		router.Use("/configs", func(ctx *fiber.Ctx) error { return ctx.JSON(core.GetConfig()) })
	}

	router.Use(cache.New(cache.Config{
		Expiration: time.Minute * 30,
	}))

	router.Use(etag.New())
}

func NewJwtGuard(router fiber.Router) {
	router.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(core.GetConfig().Server.Jwt.Secret),
	}))
}
