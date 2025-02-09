package repo

import (
	"errors"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/core/db/entity"
	"github.com/meteormin/friday.go/internal/domain"
	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	db *gorm.DB
}

func (p *PostRepositoryImpl) ExistsPostByPath(path string) (bool, error) {
	var ent entity.Post
	if err := p.db.Where("path = ?", path).First(&ent).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (p *PostRepositoryImpl) CreatePost(post *domain.Post) (*domain.Post, error) {
	ent := mapToPostEntity(post)

	if err := p.db.Create(&ent).Error; err != nil {
		return nil, err
	}

	return mapToPostModel(ent), nil
}

func (p *PostRepositoryImpl) UpdatePost(id uint, post *domain.Post) (*domain.Post, error) {
	var ent entity.Post

	if err := p.db.First(&ent, id).Error; err != nil {
		return nil, err
	}

	ent.Title = post.Title
	ent.Content = post.Content
	ent.FileID = post.FileID

	if err := p.db.Save(&ent).Error; err != nil {
		return nil, err
	}

	return mapToPostModel(ent), nil
}

func (p *PostRepositoryImpl) DeletePost(id uint) error {
	return p.db.Delete(&entity.Post{}, id).Error
}

func (p *PostRepositoryImpl) FindPost(id uint) (*domain.Post, error) {
	var ent entity.Post

	if err := p.db.First(&ent, id).Error; err != nil {
		return nil, err
	}

	return mapToPostModel(ent), nil
}

func (p *PostRepositoryImpl) RetrievePosts(query string) ([]domain.Post, error) {
	var posts []entity.Post

	tx := p.db.Where("title LIKE ?", "%"+query+"% OR content LIKE ?", "%"+query+"%").
		Find(&posts)

	if err := tx.Error; err != nil {
		return make([]domain.Post, 0), err
	}

	var results []domain.Post
	for _, post := range posts {
		results = append(results, *mapToPostModel(post))
	}

	return results, nil
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

func mapToPostEntity(post *domain.Post) entity.Post {
	tags := make([]entity.Tag, 0)
	for _, tag := range post.Tags {
		tags = append(tags, entity.Tag{
			Tag: tag.Tag,
		})
	}

	return entity.Post{
		Title:   post.Title,
		Content: post.Content,
		Path:    post.Path,
		FileID:  post.FileID,
		SiteID:  post.SiteID,
		Tags:    tags,
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
		Path:      ent.Path,
		FileID:    ent.FileID,
		File:      mapToFileModel(*ent.File),
		Tags:      tags,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
	}
}
