package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/core/http"
)

type SiteResource struct {
	ID   uint   `json:"id"`
	Host string `json:"host"`
	Name string `json:"name"`
}

type CreateSiteRequest struct {
	Host string `json:"host"`
	Name string `json:"name"`
}

type UpdateSiteRequest struct {
	Name string `json:"name"`
}

type SiteHandler struct {
	useCase port.SiteCommandUseCase
	query   port.SiteQueryUseCase
}

func (handler *SiteHandler) Retrieve(ctx *fiber.Ctx) error {
	return nil
}

func (handler *SiteHandler) Find(ctx *fiber.Ctx) error {
	return nil
}

func (handler *SiteHandler) RetrievePosts(ctx *fiber.Ctx) error {
	return nil
}

func (handler *SiteHandler) Create(ctx *fiber.Ctx) error {
	return nil
}

func (handler *SiteHandler) Update(ctx *fiber.Ctx) error {
	return nil
}

func (handler *SiteHandler) Delete(ctx *fiber.Ctx) error {
	return nil
}

func NewSiteHandler(useCase port.SiteCommandUseCase, query port.SiteQueryUseCase) http.AddRouteFunc {

	handler := &SiteHandler{
		useCase: useCase,
		query:   query,
	}

	return func(router fiber.Router) {
		group := router.Group("/sites")
		group.Get("/", handler.Retrieve)
		group.Get("/:id", handler.Find)
		group.Get("/:id/posts", handler.RetrievePosts)
		group.Post("/", handler.Create)
		group.Put("/:id", handler.Update)
		group.Delete("/:id", handler.Delete)
	}
}
