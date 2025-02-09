package repo

import (
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/core/db/entity"
	"github.com/meteormin/friday.go/internal/domain"
	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	db *gorm.DB
}

func (p PostRepositoryImpl) CreatePost(post domain.Post) (*domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostRepositoryImpl) UpdatePost(id uint, post domain.Post) (*domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostRepositoryImpl) DeletePost(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (p PostRepositoryImpl) FindPost(id uint) (*domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostRepositoryImpl) RetrievePosts(query string) ([]domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func NewPostRepository(db *gorm.DB) port.PostRepository {
	return &PostRepositoryImpl{
		db: db,
	}
}

func mapToTagModel(ent entity.Tag) *domain.Tag {
	return &domain.Tag{
		ID:  ent.ID,
		Tag: ent.Tag,
	}
}

func mapToPostModel(ent entity.Post) *domain.Post {
	tags := make([]domain.Tag, 0)
	for _, tag := range ent.Tags {
		tags = append(tags, *mapToTagModel(tag))
	}

	return &domain.Post{
		ID:        ent.ID,
		Title:     ent.Title,
		Content:   ent.Content,
		FileID:    ent.FileID,
		File:      mapToFileModel(*ent.File),
		Tags:      tags,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
	}
}
