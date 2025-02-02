package domain

import "io"

type File struct {
	ID         uint
	OriginName string
	Extension  string
	Size       uint
	FilePath   string
}

type UploadFile struct {
	FileName string
	Reader   io.Reader
}

type UploadFileCommand interface {
	Upload(file UploadFile) File
}
