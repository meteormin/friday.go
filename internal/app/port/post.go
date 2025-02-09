package port

import (
	"github.com/meteormin/friday.go/internal/app/errors"
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
		return nil, errors.ErrInvalidPostTitle
	}

	if c.Path == "" {
		return nil, errors.ErrInvalidPostPath
	}

	if c.Content == "" {
		return nil, errors.ErrInvalidPostContent
	}

	if c.FileID <= 0 {
		return nil, errors.ErrInvalidPostFile
	}

	if c.SiteID <= 0 {
		return nil, errors.ErrInvalidPostSite
	}

	tags := make([]domain.Tag, 0)
	if len(c.Tags) != 0 {
		for _, tag := range c.Tags {
			if tag == "" {
				return nil, errors.ErrInvalidPostTags
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
		return nil, errors.ErrInvalidPostTitle
	}

	if u.Content == "" {
		return nil, errors.ErrInvalidPostContent
	}

	if u.Path == "" {
		return nil, errors.ErrInvalidPostPath
	}

	if u.FileID <= 0 {
		return nil, errors.ErrInvalidPostFile
	}

	tags := make([]domain.Tag, 0)
	if len(u.Tags) != 0 {
		for _, tag := range u.Tags {
			if tag == "" {
				return nil, errors.ErrInvalidPostTags
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
	CreatePost(post CreatePost) (*domain.Post, error)
	UpdatePost(id uint, post UpdatePost) (*domain.Post, error)
	DeletePost(id uint) error
}

type PostQueryUseCase interface {
	FindPost(id uint) (*domain.Post, error)

	RetrievePosts(query string) ([]domain.Post, error)
}

type PostRepository interface {
	ExistsPostByPath(path string) (bool, error)

	CreatePost(post *domain.Post) (*domain.Post, error)

	UpdatePost(id uint, post *domain.Post) (*domain.Post, error)

	DeletePost(id uint) error

	FindPost(id uint) (*domain.Post, error)

	RetrievePosts(query string) ([]domain.Post, error)
}
