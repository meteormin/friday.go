package app

import (
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/domain"
)

type UploadFileCommandService struct {
	repo port.FileRepository
}

func (u UploadFileCommandService) UploadFile(uploadFile port.UploadFile) (*domain.File, error) {
	file, err := uploadFile.Valid()
	if err != nil {
		return nil, err
	}

	return u.repo.CreateFile(file)
}

func NewUploadFileService(repo port.FileRepository) port.UploadFileUseCase {
	return &UploadFileCommandService{
		repo: repo,
	}
}
