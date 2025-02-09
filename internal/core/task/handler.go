package task

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/meteormin/friday.go/internal/core"
	"github.com/meteormin/friday.go/internal/core/http"
)

func Handler(router fiber.Router) {
	jobRepo := NewJobRepository(core.GetDB())
	jobs := func(router fiber.Router) {
		router.Group("/tasks").
			Get("/", func(ctx *fiber.Ctx) error {
				jobs := jobRepo.All()
				return ctx.JSON(http.NewContentResource(jobs))
			}).
			Get("/:jobId", func(ctx *fiber.Ctx) error {
				id := ctx.Params("jobId")
				jobID, err := uuid.Parse(id)
				if err != nil {
					return fiber.NewError(fiber.StatusBadRequest, err.Error())
				}

				job, err := jobRepo.FindByJobID(jobID)
				if err != nil {
					return fiber.NewError(fiber.StatusNotFound, err.Error())
				}

				return ctx.JSON(job)
			})
	}

	jobs(router)
}
