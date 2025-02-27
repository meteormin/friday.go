package port

import (
	"github.com/meteormin/friday.go/internal/app"
	"github.com/meteormin/friday.go/internal/domain"
)

type CreatePost struct {
	Title   string
	Content string
	Path    string
	FileID  uint
	SiteID  uint
	Tags    []string
}

func (c CreatePost) Valid() (*domain.Post, error) {

	if c.Title == "" {
		return nil, app.ErrInvalidPostTitle
	}

	if c.Path == "" {
		return nil, app.ErrInvalidPostPath
	}

	if c.Content == "" {
		return nil, app.ErrInvalidPostContent
	}

	if c.FileID <= 0 {
		return nil, app.ErrInvalidPostFile
	}

	if c.SiteID <= 0 {
		return nil, app.ErrInvalidPostSite
	}

	tags := make([]domain.Tag, 0)
	if len(c.Tags) != 0 {
		for _, tag := range c.Tags {
			if tag == "" {
				return nil, app.ErrInvalidPostTags
			}

			tags = append(tags, domain.Tag{
				Tag: tag,
			})
		}
	}

	return &domain.Post{
		Title:   c.Title,
		Content: c.Content,
		Path:    c.Path,
		FileID:  c.FileID,
		SiteID:  c.SiteID,
		Tags:    tags,
	}, nil
}

type UpdatePost struct {
	Title   string
	Content string
	Path    string
	FileID  uint
	Tags    []string
}

func (u UpdatePost) Valid() (*domain.Post, error) {
	if u.Title == "" {
		return nil, app.ErrInvalidPostTitle
	}

	if u.Content == "" {
		return nil, app.ErrInvalidPostContent
	}

	if u.Path == "" {
		return nil, app.ErrInvalidPostPath
	}

	if u.FileID <= 0 {
		return nil, app.ErrInvalidPostFile
	}

	tags := make([]domain.Tag, 0)
	if len(u.Tags) != 0 {
		for _, tag := range u.Tags {
			if tag == "" {
				return nil, app.ErrInvalidPostTags
			}

			tags = append(tags, domain.Tag{
				Tag: tag,
			})
		}
	}

	return &domain.Post{
		Title:   u.Title,
		Content: u.Content,
		Path:    u.Path,
		FileID:  u.FileID,
		Tags:    tags,
	}, nil
}

type PostCommandUseCase interface {
	CreatePost(userId uint, post CreatePost) (*domain.Post, error)
	UpdatePost(userId, id uint, post UpdatePost) (*domain.Post, error)
	DeletePost(userId, id uint) error
}

type PostQueryUseCase interface {
	FindPost(userId, id uint) (*domain.Post, error)

	RetrievePosts(userId uint, query string) ([]domain.Post, error)
}

type PostRepository interface {
	ExistsPostByPath(siteId uint, path string) (bool, error)

	HasAccessPermission(userId, id uint) (bool, error)

	CreatePost(post *domain.Post) (*domain.Post, error)

	UpdatePost(id uint, post *domain.Post) (*domain.Post, error)

	DeletePost(id uint) error

	FindPost(id uint) (*domain.Post, error)

	RetrievePosts(userId uint, query string) ([]domain.Post, error)
}
