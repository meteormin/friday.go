package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/core/http"
	"strconv"
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
	sites, err := handler.query.RetrieveSite(ctx.Query("query"))
	if err != nil {
		return err
	}
	return ctx.JSON(http.NewContentResource(sites))
}

func (handler *SiteHandler) Find(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	site, err := handler.query.FindSite(uint(id))
	if err != nil {
		return err
	}

	return ctx.JSON(site)
}

func (handler *SiteHandler) RetrievePosts(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	site, err := handler.query.FindSite(uint(id))
	if err != nil {
		return err
	}

	return ctx.JSON(http.NewContentResource(site.Posts))
}

func (handler *SiteHandler) Create(ctx *fiber.Ctx) error {
	var req CreateSiteRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	site, err := handler.useCase.CreateSite(port.CreateSite{
		Host: req.Host,
		Name: req.Name,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(site)
}

func (handler *SiteHandler) Update(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var req UpdateSiteRequest
	err = ctx.BodyParser(&req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	site, err := handler.useCase.UpdateSite(uint(id), port.UpdateSite{
		Name: req.Name,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(site)
}

func (handler *SiteHandler) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return handler.useCase.DeleteSite(uint(id))
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
