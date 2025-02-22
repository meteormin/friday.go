package entity

import "gorm.io/gorm"

type File struct {
	gorm.Model
	OriginFilename string
	ConvFilename   string
	Path           string
	MimeType       string
	Size           uint64
	UserID         uint
	User           User
}
