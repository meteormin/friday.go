package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/core/http"
	"path"
)

type UploadFileResponse struct {
	ID  uint   `json:"id"`
	URL string `json:"url"`
}

type UploadFileHandler struct {
	useCase port.UploadFileUseCase
}

func (handler *UploadFileHandler) uploadFile(ctx *fiber.Ctx) error {

	uploadFile, err := ctx.FormFile("file")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	fileUUID, err := uuid.NewUUID()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	convFilename := fileUUID.String()

	data, err := uploadFile.Open()
	if err != nil {
		return err
	}

	bytes := make([]byte, uploadFile.Size)
	_, err = data.Read(bytes)
	if err != nil {
		return err
	}

	file, err := handler.useCase.UploadFile(port.UploadFile{
		FileName: uploadFile.Filename,
		Size:     uint(uploadFile.Size),
		Data:     bytes,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(UploadFileResponse{
		ID:  file.ID,
		URL: path.Join("/files", convFilename),
	})
}

func NewUploadFileHandler(useCase port.UploadFileUseCase) http.AddRouteFunc {

	handler := &UploadFileHandler{
		useCase: useCase,
	}

	return func(router fiber.Router) {
		group := router.Group("/upload-file")
		group.Post("/", handler.uploadFile)
	}
}
