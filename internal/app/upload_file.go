package app

import (
	"github.com/google/uuid"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/domain"
	_ "github.com/meteormin/friday.go/internal/infra"
	"path"
)

type UploadFileCommandService struct {
	repo port.FileRepository
}

func (u UploadFileCommandService) UploadFile(uploadFile port.UploadFile) (*domain.File, error) {
	file, err := uploadFile.Valid()
	if err != nil {
		return nil, err
	}

	fileUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	file.FilePath = path.Join("tmp", fileUUID.String())

	return u.repo.CreateFile(file)
}

func NewUploadFileCommandService(repo port.FileRepository) port.UploadFileUseCase {
	return &UploadFileCommandService{
		repo: repo,
	}
}
