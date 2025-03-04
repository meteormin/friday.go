package repo

import (
	"errors"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/core/db/entity"
	"github.com/meteormin/friday.go/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostRepositoryImpl struct {
	db *gorm.DB
}

func (p *PostRepositoryImpl) HasAccessPermission(userId, id uint) (bool, error) {
	var count int64
	if err := p.db.Preload("Site", "user_id = ?", userId).
		Where("id = ?", id).
		Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return count > 0, nil
}

func (p *PostRepositoryImpl) ExistsPostByPath(siteId uint, path string) (bool, error) {
	var ent entity.Post
	if err := p.db.Where("site_id = ? AND path = ?", siteId, path).First(&ent).Error; err != nil {
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

	err := p.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload(clause.Associations).First(&ent, id).Error; err != nil {
			return err
		}

		if err := tx.Model(&ent).Updates(entity.Post{
			Title:   post.Title,
			Content: post.Content,
			FileID:  post.FileID,
			SiteID:  post.SiteID,
		}).Error; err != nil {
			return err
		}

		if err := updatePostTags(tx, id, post.Tags); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	p.db.Preload(clause.Associations).First(&ent, id)

	return mapToPostModel(ent), nil
}

func updatePostTags(tx *gorm.DB, id uint, tags []domain.Tag) error {
	var tagEntities []entity.Tag
	if err := tx.Where("post_id = ?", id).Find(&tagEntities).Error; err != nil {
		return err
	}

	for _, tag := range tags {
		tagEntities = append(tagEntities, entity.Tag{
			Tag: tag.Tag,
		})
	}

	if err := tx.Model(&entity.Tag{}).
		Where("post_id = ?", id).
		Updates(tagEntities).Error; err != nil {
		return err
	}

	return nil
}

func (p *PostRepositoryImpl) DeletePost(id uint) error {
	return p.db.Delete(&entity.Post{}, id).Error
}

func (p *PostRepositoryImpl) FindPost(id uint) (*domain.Post, error) {
	var ent entity.Post

	if err := p.db.Preload(clause.Associations).First(&ent, id).Error; err != nil {
		return nil, err
	}

	return mapToPostModel(ent), nil
}

func (p *PostRepositoryImpl) RetrievePosts(userId uint, query string) ([]domain.Post, error) {
	var posts []entity.Post

	tx := p.db.Preload("Site", "user_id = ?", userId).
		Preload("File").
		Preload("Tags").
		Where("title LIKE ? OR content LIKE ?", "%"+query+"%", "%"+query+"%").
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
