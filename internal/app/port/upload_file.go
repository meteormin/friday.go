package port

import (
	"github.com/meteormin/friday.go/internal/app/errors"
	"github.com/meteormin/friday.go/internal/domain"
	"io"
	"path/filepath"
)

type UploadFile struct {
	FileName string
	Size     uint
	Reader   io.Reader
}

func (u UploadFile) Valid() (*domain.File, error) {
	if u.FileName == "" {
		return nil, errors.ErrInvalidFileName
	}

	if u.Reader == nil {
		return nil, errors.ErrInvalidFile
	}

	return &domain.File{
		OriginName: u.FileName,
		Extension:  filepath.Ext(u.FileName),
		Size:       u.Size,
		FilePath:   "",
	}, nil
}

type UploadFileUseCase interface {
	UploadFile(file UploadFile) (*domain.File, error)
}

type FileRepository interface {
	CreateFile(file *domain.File) (*domain.File, error)
}
