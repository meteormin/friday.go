package repo

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/core/db/entity"
	"github.com/meteormin/friday.go/internal/domain"
	"github.com/meteormin/friday.go/pkg/database"
	"gorm.io/gorm"
)

type FileRepositoryImpl struct {
	basePath string
	db       *gorm.DB
	storage  *badger.DB
}

func (f FileRepositoryImpl) CreateFile(file *domain.File, data []byte) (*domain.File, error) {
	ent := mapToFileEntity(file)

	err := database.PutFile(f.storage, ent.ConvFilename, data)
	if err != nil {
		return nil, err
	}

	f.db.Create(&ent)
	return mapToFileModel(ent), nil
}

func NewFileRepository(basePath string, db *gorm.DB, storage *badger.DB) port.FileRepository {
	return &FileRepositoryImpl{
		basePath: basePath,
		db:       db,
		storage:  storage,
	}
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
