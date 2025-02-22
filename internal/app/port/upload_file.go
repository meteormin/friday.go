package port

import (
	"github.com/google/uuid"
	"github.com/meteormin/friday.go/internal/app"
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
		return nil, app.ErrInvalidFileName
	}

	if u.Data == nil {
		return nil, app.ErrInvalidFile
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
	UploadFile(userId uint, file UploadFile) (*domain.File, error)

	DownloadFIle(userId, fileID uint) ([]byte, *domain.File, error)
}

type FileRepository interface {
	HasAccessPermission(userId, fileID uint) (bool, error)

	CreateFile(file *domain.File, data []byte) (*domain.File, error)

	FindFile(id uint) ([]byte, *domain.File, error)
}
