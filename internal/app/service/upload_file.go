package service

import (
	"github.com/meteormin/friday.go/internal/app"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/domain"
)

type UploadFileCommandService struct {
	repo port.FileRepository
}

func (u UploadFileCommandService) DownloadFIle(userID uint, fileID uint) ([]byte, *domain.File, error) {
	if hasPerm, err := u.repo.HasAccessPermission(userID, fileID); err != nil {
		return nil, nil, err
	} else if !hasPerm {
		return nil, nil, app.ErrForbidden
	}

	return u.repo.FindFile(fileID)
}

func (u UploadFileCommandService) UploadFile(userID uint, uploadFile port.UploadFile) (*domain.File, error) {
	file, err := uploadFile.Valid()
	if err != nil {
		return nil, err
	}

	file.User.ID = userID

	return u.repo.CreateFile(file, uploadFile.Data)
}

func NewUploadFileService(repo port.FileRepository) port.UploadFileUseCase {
	return &UploadFileCommandService{
		repo: repo,
	}
}
