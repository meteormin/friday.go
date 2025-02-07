package repo

import (
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/core/db/entity"
	"github.com/meteormin/friday.go/internal/domain"
	"gorm.io/gorm"
)

type FileRepositoryImpl struct {
	db       *gorm.DB
	basePath string
}

func (f FileRepositoryImpl) CreateFile(file *domain.File) (*domain.File, error) {
	ent := mapToFileEntity(file)
	f.db.Create(&ent)
	return mapToFileModel(ent), nil
}

func NewFileRepository(db *gorm.DB, basePath string) port.FileRepository {
	return &FileRepositoryImpl{db: db, basePath: basePath}
}

func mapToFileEntity(model *domain.File) entity.File {
	return entity.File{
		OriginFilename: model.OriginName,
		ConvFilename:   model.ConvName,
		Path:           model.FilePath,
		MimeType:       model.MimeType,
		Size:           model.Size,
	}
}

func mapToFileModel(ent entity.File) *domain.File {
	return &domain.File{
		ID:         ent.ID,
		OriginName: ent.OriginFilename,
		ConvName:   ent.ConvFilename,
		MimeType:   ent.MimeType,
		Size:       ent.Size,
		FilePath:   ent.Path,
	}
}
