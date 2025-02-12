package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	_ "github.com/meteormin/friday.go/internal/app/errors"
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

// UploadFile
// @Summary 파일 업로드
// @Description 파일 업로드 API
// @ID files.upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "파일"
// @Success 201 {object} UploadFileResponse "파일 업로드 성공"
// @Failure 400 {object} errors.Error "잘못된 요청"
// @Failure 500 {object} errors.Error "서버 오류"
// @Router /api/upload-file [post]
// @Tags upload-file
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

	return ctx.Status(fiber.StatusCreated).JSON(UploadFileResponse{
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
