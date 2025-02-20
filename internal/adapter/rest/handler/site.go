package handler

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/meteormin/friday.go/internal/app/errors"
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

// Find
// @Summary 사이트 조회
// @Description 사이트 조회 API
// @ID sites.find
// @Accept json
// @Produce json
// @Param id path int true "사이트 ID"
// @Success 200 {object} SiteResource "사이트 조회 성공"
// @Failure 400 {object} errors.Error "잘못된 요청"
// @Failure 404 {object} errors.Error "사이트 없음"
// @Router /api/sites/{id} [get]
// @Tags sites
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

// RetrievePosts
// @Summary 사이트 포스트 리스트 조회
// @Description 사이트 포스트 리스트 조회 API
// @ID sites.retrievePosts
// @Accept json
// @Produce json
// @Param id path int true "사이트 ID"
// @Success 200 {array} PostResource "사이트 포스트 리스트 조회 성공"
// @Failure 400 {object} errors.Error "잘못된 요청"
// @Failure 404 {object} errors.Error "사이트 없음"
// @Router /api/sites/{id}/posts [get]
// @Tags sites
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

// Create
// @Summary 사이트 생성
// @Description 사이트 생성 API
// @ID sites.create
// @Accept json
// @Produce json
// @Param req body CreateSiteRequest true "사이트 생성 정보"
// @Success 201 {object} SiteResource "사이트 생성 성공"
// @Failure 400 {object} errors.Error "잘못된 요청"
// @Failure 409 {object} errors.Error "이메일 중복"
// @Failure 500 {object} errors.Error "서버 오류"
// @Router /api/sites [post]
// @Tags sites
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

	return ctx.Status(fiber.StatusCreated).JSON(site)
}

// Update
// @Summary 사이트 수정
// @Description 사이트 수정 API
// @ID sites.update
// @Accept json
// @Produce json
// @Param id path int true "사이트 ID"
// @Param req body UpdateSiteRequest true "사이트 수정 정보"
// @Success 200 {object} SiteResource "사이트 수정 성공"
// @Failure 400 {object} errors.Error "잘못된 요청"
// @Failure 404 {object} errors.Error "사이트 없음"
// @Failure 500 {object} errors.Error "서버 오류"
// @Router /api/sites/{id} [put]
// @Tags sites
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

// Delete
// @Summary 사이트 삭제
// @Description 사이트 삭제 API
// @ID sites.delete
// @Accept json
// @Produce json
// @Param id path int true "사이트 ID"
// @Success 204 {object} SiteResource "사이트 삭제 성공"
// @Failure 400 {object} errors.Error "잘못된 요청"
// @Failure 404 {object} errors.Error "사이트 없음"
// @Failure 500 {object} errors.Error "서버 오류"
// @Router /api/sites/{id} [delete]
// @Tags sites
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
