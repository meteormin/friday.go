package http

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	app_errors "github.com/meteormin/friday.go/internal/app/errors"
)

type ErrorResponse struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		var fiberError *fiber.Error
		var domainError *app_errors.Error
		if errors.As(err, &fiberError) {
			return c.Status(fiberError.Code).JSON(ErrorResponse{
				Title:   fiberError.Message,
				Message: fiberError.Message,
			})
		} else if errors.As(err, &domainError) {
			return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
				Title:   domainError.Title,
				Message: domainError.Message,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Title:   "InternalServerError",
			Message: err.Error(),
		})
	}
}
