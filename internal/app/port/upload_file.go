package port

import (
	"github.com/google/uuid"
	"github.com/meteormin/friday.go/internal/app/errors"
	"github.com/meteormin/friday.go/internal/domain"
	"mime"
	"path"
)

type UploadFile struct {
	FileName string
	Size     uint
	Data     []byte
}

func (u UploadFile) Valid() (*domain.File, error) {
	if u.FileName == "" {
		return nil, errors.ErrInvalidFileName
	}

	if u.Data == nil {
		return nil, errors.ErrInvalidFile
	}

	convName, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &domain.File{
		OriginName: u.FileName,
		ConvName:   convName.String(),
		MimeType:   mime.TypeByExtension(u.FileName),
		Size:       uint64(u.Size),
		FilePath:   path.Join("tmp", convName.String()),
	}, nil
}

type UploadFileUseCase interface {
	UploadFile(file UploadFile) (*domain.File, error)
}

type FileRepository interface {
	CreateFile(file *domain.File) (*domain.File, error)
}
