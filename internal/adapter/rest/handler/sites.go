package handler

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/meteormin/friday.go/internal/app"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/core/http"
	"github.com/meteormin/friday.go/internal/domain"
	"strconv"
)

type SiteResource struct {
	ID        uint           `json:"id"`
	Host      string         `json:"host"`
	Name      string         `json:"name"`
	CreatedAt string         `json:"createdAt"`
	UpdatedAt string         `json:"updatedAt"`
	Posts     []PostResource `json:"posts"`
}

func mapToSiteResource(site domain.Site) SiteResource {
	posts := make([]PostResource, 0)
	for _, post := range site.Posts {
		posts = append(posts, mapToPostResource(post))
	}

	return SiteResource{
		ID:        site.ID,
		Host:      site.Host,
		Name:      site.Name,
		CreatedAt: site.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: site.UpdatedAt.Format("2006-01-02 15:04:05"),
		Posts:     posts,
	}
}

type CreateSiteRequest struct {
	Host string `json:"host"`
	Name string `json:"name"`
}

func (r CreateSiteRequest) ToDomain() port.CreateSite {
	return port.CreateSite{
		Host: r.Host,
		Name: r.Name,
	}
}

type UpdateSiteRequest struct {
	Name string `json:"name"`
}

type SiteHandler struct {
	useCase port.SiteCommandUseCase
	query   port.SiteQueryUseCase
}

// Retrieve
// @Summary 사이트 리스트 조회
// @Description 사이트 리스트 조회 API
// @ID sites.retrieve
// @Accept json
// @Produce json
// @Success 200 {array} SiteResource "사이트 리스트 조회 성공"
// @Failure 400 {object} app.Error "잘못된 요청"
// @Failure 500 {object} app.Error "서버 오류"
// @Router /api/sites [get]
// @Tags sites
// @Security BearerAuth
func (handler *SiteHandler) Retrieve(ctx *fiber.Ctx) error {
	userId := http.ExtractTokenClaims(ctx)["id"].(uint)

	sites, err := handler.query.RetrieveSite(userId, ctx.Query("query"))
	if err != nil {
		return err
	}

	resources := make([]SiteResource, 0)
	for _, site := range sites {
		resources = append(resources, mapToSiteResource(site))
	}

	return ctx.JSON(http.NewContentResource(resources))
}

// Find
// @Summary 사이트 조회
// @Description 사이트 조회 API
// @ID sites.find
// @Accept json
// @Produce json
// @Param id path int true "사이트 ID"
// @Success 200 {object} SiteResource "사이트 조회 성공"
// @Failure 400 {object} app.Error "잘못된 요청"
// @Failure 404 {object} app.Error "사이트 없음"
// @Router /api/sites/{id} [get]
// @Tags sites
// @Security BearerAuth
func (handler *SiteHandler) Find(ctx *fiber.Ctx) error {
	userId := http.ExtractTokenClaims(ctx)["id"].(uint)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	site, err := handler.query.FindSite(userId, uint(id))
	if err != nil {
		return err
	}

	return ctx.JSON(mapToSiteResource(*site))
}

// RetrievePosts
// @Summary 사이트 포스트 리스트 조회
// @Description 사이트 포스트 리스트 조회 API
// @ID sites.retrievePosts
// @Accept json
// @Produce json
// @Param id path int true "사이트 ID"
// @Success 200 {array} PostResource "사이트 포스트 리스트 조회 성공"
// @Failure 400 {object} app.Error "잘못된 요청"
// @Failure 404 {object} app.Error "사이트 없음"
// @Router /api/sites/{id}/posts [get]
// @Tags sites
// @Security BearerAuth
func (handler *SiteHandler) RetrievePosts(ctx *fiber.Ctx) error {
	userId := http.ExtractTokenClaims(ctx)["id"].(uint)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	site, err := handler.query.FindSite(userId, uint(id))
	if err != nil {
		return err
	}

	return ctx.JSON(http.NewContentResource(mapToSiteResource(*site).Posts))
}

// Create
// @Summary 사이트 생성
// @Description 사이트 생성 API
// @ID sites.create
// @Accept json
// @Produce json
// @Param req body CreateSiteRequest true "사이트 생성 정보"
// @Success 201 {object} SiteResource "사이트 생성 성공"
// @Failure 400 {object} app.Error "잘못된 요청"
// @Failure 409 {object} app.Error "이메일 중복"
// @Failure 500 {object} app.Error "서버 오류"
// @Router /api/sites [post]
// @Tags sites
// @Security BearerAuth
func (handler *SiteHandler) Create(ctx *fiber.Ctx) error {
	userId := http.ExtractTokenClaims(ctx)["id"].(uint)

	var req CreateSiteRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	site, err := handler.useCase.CreateSite(port.CreateSite{
		Host:   req.Host,
		Name:   req.Name,
		UserID: userId,
	})

	if err != nil {
		return err
	}

	resource := mapToSiteResource(*site)

	return ctx.Status(fiber.StatusCreated).JSON(resource)
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
// @Failure 400 {object} app.Error "잘못된 요청"
// @Failure 404 {object} app.Error "사이트 없음"
// @Failure 500 {object} app.Error "서버 오류"
// @Router /api/sites/{id} [put]
// @Tags sites
// @Security BearerAuth
func (handler *SiteHandler) Update(ctx *fiber.Ctx) error {
	userId := http.ExtractTokenClaims(ctx)["id"].(uint)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var req UpdateSiteRequest
	err = ctx.BodyParser(&req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	site, err := handler.useCase.UpdateSite(userId, uint(id), port.UpdateSite{
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
// @Failure 400 {object} app.Error "잘못된 요청"
// @Failure 404 {object} app.Error "사이트 없음"
// @Failure 500 {object} app.Error "서버 오류"
// @Router /api/sites/{id} [delete]
// @Tags sites
// @Security BearerAuth
func (handler *SiteHandler) Delete(ctx *fiber.Ctx) error {
	userId := http.ExtractTokenClaims(ctx)["id"].(uint)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return handler.useCase.DeleteSite(userId, uint(id))
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
