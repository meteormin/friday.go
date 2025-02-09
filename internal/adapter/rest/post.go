package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/core/http"
	"strconv"
)

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

func (h PostHandler) Retrieve(ctx *fiber.Ctx) error {
	posts, err := h.query.RetrievePosts(ctx.Query("query"))
	if err != nil {
		return err
	}
	return ctx.JSON(http.NewContentResource(posts))
}

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

	return ctx.JSON(post)
}

func (h PostHandler) Update(ctx *fiber.Ctx) error {

	id, err := strconv.Atoi(ctx.Params("id"))

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

func (h PostHandler) Delete(ctx *fiber.Ctx) error {

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return h.command.DeletePost(uint(id))
}

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
