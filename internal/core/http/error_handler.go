package http

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	apperrors "github.com/meteormin/friday.go/internal/app/errors"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		var fiberError *fiber.Error
		var domainError *apperrors.Error
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
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
				Title:   "NotFound",
				Message: err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Title:   "InternalServerError",
			Message: err.Error(),
		})
	}
}
