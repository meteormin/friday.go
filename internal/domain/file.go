package domain

import (
	"github.com/google/uuid"
	"github.com/meteormin/friday.go/internal/infra"
	"io"
	"path"
	"path/filepath"
)

var (
	ErrInvalidFileName = NewError(400, "InvalidFileName", "invalid file name")
	ErrInvalidFile     = NewError(400, "InvalidFile", "invalid file")
)

type File struct {
	ID         uint
	OriginName string
	Extension  string
	Size       uint
	FilePath   string
}

type UploadFile struct {
	FileName string
	Size     uint
	Reader   io.Reader
}

func (u *UploadFile) Valid() (*File, error) {
	if u.FileName == "" {
		return nil, ErrInvalidFileName
	}

	if u.Reader == nil {
		return nil, ErrInvalidFile
	}

	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	filePath := path.Join(infra.GetConfig().Path.Data, "tmp", newUUID.String())

	return &File{
		OriginName: u.FileName,
		Extension:  filepath.Ext(u.FileName),
		Size:       u.Size,
		FilePath:   filePath,
	}, nil
}

type UploadFileCommand interface {
	Upload(file UploadFile) File
}
