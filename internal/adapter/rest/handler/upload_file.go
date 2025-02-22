package handler

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	_ "github.com/meteormin/friday.go/internal/app"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/core/http"
	"path"
	"strconv"
)

type UploadFileResponse struct {
	ID  uint   `json:"id"`
	URI string `json:"uri"`
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
// @Failure 400 {object} app.Error "잘못된 요청"
// @Failure 500 {object} app.Error "서버 오류"
// @Router /api/upload-file [post]
// @Tags upload-file
func (handler *UploadFileHandler) uploadFile(ctx *fiber.Ctx) error {
	userId := http.ExtractTokenClaims(ctx)["id"].(uint)

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

	fBytes := make([]byte, uploadFile.Size)
	_, err = data.Read(fBytes)
	if err != nil {
		return err
	}

	file, err := handler.useCase.UploadFile(userId, port.UploadFile{
		FileName: uploadFile.Filename,
		Size:     uint(uploadFile.Size),
		Data:     fBytes,
	})

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(UploadFileResponse{
		ID:  file.ID,
		URI: path.Join("/upload-file", convFilename),
	})
}

// DownloadFile
// @Summary 파일 다운로드
// @Description 파일 다운로드 API
// @ID files.download
// @Produce octet-stream
// @Param id path string true "파일 ID"
// @Success 200 {object} []byte "파일 다운로드 성공"
// @Failure 400 {object} app.Error "잘못된 요청"
// @Failure 404 {object} app.Error "파일 없음"
// @Failure 500 {object} app.Error "서버 오류"
// @Router /api/upload-file/{id} [get]
// @Tags upload-file
func (handler *UploadFileHandler) downloadFile(ctx *fiber.Ctx) error {
	userId := http.ExtractTokenClaims(ctx)["id"].(uint)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	file, domainModel, err := handler.useCase.DownloadFIle(userId, uint(id))
	if err != nil {
		return err
	}

	return ctx.Type(domainModel.MimeType).SendStream(bytes.NewReader(file))
}

func NewUploadFileHandler(useCase port.UploadFileUseCase) http.AddRouteFunc {

	handler := &UploadFileHandler{
		useCase: useCase,
	}

	return func(router fiber.Router) {
		group := router.Group("/upload-file")
		group.Post("/", handler.uploadFile)
		group.Get("/:id", handler.downloadFile)
	}
}
