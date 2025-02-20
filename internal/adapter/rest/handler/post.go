package handler

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/meteormin/friday.go/internal/app/errors"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/core/http"
	"strconv"
)

// CreatePostRequest
// @Description 생성 요청
type CreatePostRequest struct {
	SiteID  uint     `json:"siteId"`
	FileID  uint     `json:"fileId"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type UpdatePostRequest struct {
	FileID  uint     `json:"fileId"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type PostResource struct {
	ID        uint     `json:"id"`
	URI       string   `json:"uri"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type PostHandler struct {
	command port.PostCommandUseCase
	query   port.PostQueryUseCase
}

// Retrieve
// @Summary 포스트 리스트 조회
// @Description 포스트 리스트 조회 API
// @ID posts.retrieve
// @Accept json
// @Produce json
// @Success 200 {array} PostResource "포스트 리스트 조회 성공"
// @Failure 500 {object} errors.Error "서버 오류"
// @Router /api/posts [get]
// @Tags posts
func (h PostHandler) Retrieve(ctx *fiber.Ctx) error {
	posts, err := h.query.RetrievePosts(ctx.Query("query"))
	if err != nil {
		return err
	}
	return ctx.JSON(http.NewContentResource(posts))
}

// Find
// @Summary 포스트 조회
// @Description 포스트 조회 API
// @ID posts.find
// @Accept json
// @Produce json
// @Param id path int true "포스트 ID"
// @Success 200 {object} PostResource "포스트 조회 성공"
// @Failure 400 {object} errors.Error "잘못된 요청"
// @Failure 500 {object} errors.Error "서버 오류"
// @Router /api/posts/{id} [get]
// @Tags posts
func (h PostHandler) Find(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	post, err := h.query.FindPost(uint(id))
	if err != nil {
		return err
	}

	return ctx.JSON(post)
}

// Create
// @Summary 포스트 생성
// @Description 포스트 생성 API
// @ID posts.create
// @Accept json
// @Produce json
// @Param req body CreatePostRequest true "포스트 생성 정보"
// @Success 201 {object} PostResource "포스트 생성 성공"
// @Failure 400 {object} errors.Error "잘못된 요청"
// @Failure 409 {object} errors.Error "이메일 중복"
// @Failure 500 {object} errors.Error "서버 오류"
// @Router /api/posts [post]
// @Tags posts
func (h PostHandler) Create(ctx *fiber.Ctx) error {
	var req CreatePostRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	post, err := h.command.CreatePost(port.CreatePost{
		SiteID:  req.SiteID,
		FileID:  req.FileID,
		Title:   req.Title,
		Content: req.Content,
		Tags:    req.Tags,
	})

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(post)
}

// Update
// @Summary 포스트 수정
// @Description 포스트 수정 API
// @ID posts.update
// @Accept json
// @Produce json
// @Param id path int true "포스트 ID"
// @Param req body UpdatePostRequest true "포스트 수정 정보"
// @Success 200 {object} PostResource "포스트 수정 성공"
// @Failure 400 {object} errors.Error "잘못된 요청"
// @Failure 500 {object} errors.Error "서버 오류"
// @Router /api/posts/{id} [put]
// @Tags posts
func (h PostHandler) Update(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var update UpdatePostRequest
	err = ctx.BodyParser(&update)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	post, err := h.command.UpdatePost(uint(id), port.UpdatePost{
		FileID:  update.FileID,
		Title:   update.Title,
		Content: update.Content,
		Tags:    update.Tags,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(post)
}

// Delete
// @Summary 포스트 삭제
// @Description 포스트 삭제 API
// @ID posts.delete
// @Accept json
// @Produce json
// @Param id path int true "포스트 ID"
// @Success 204 {object} PostResource "포스트 삭제 성공"
// @Failure 404 {object} errors.Error "존재하지 않는 포스트"
// @Failure 500 {object} errors.Error "서버 오류"
// @Router /api/posts/{id} [delete]
// @Tags posts
func (h PostHandler) Delete(ctx *fiber.Ctx) error {

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = h.command.DeletePost(uint(id))
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// NewPostHandler 포스트 관리 Handler 생성
func NewPostHandler(command port.PostCommandUseCase, query port.PostQueryUseCase) http.AddRouteFunc {

	handler := &PostHandler{
		command: command,
		query:   query,
	}

	return func(router fiber.Router) {
		group := router.Group("/posts")
		group.Get("/", handler.Retrieve)
		group.Get("/:id", handler.Find)
		group.Post("/", handler.Create)
		group.Put("/:id", handler.Update)
		group.Delete("/:id", handler.Delete)
	}
}
