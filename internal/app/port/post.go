package port

import "github.com/meteormin/friday.go/internal/domain"

type CreatePost struct {
	Title   string
	Content string
	FileID  uint
	SiteID  uint
}

type UpdatePost struct {
	Title   string
	Content string
	FileID  uint
}

type PostCommandUseCase interface {
	CreatePost(post CreatePost) (*domain.Post, error)
	UpdatePost(id uint, post UpdatePost) (*domain.Post, error)
	DeletePost(id uint) error
}

type PostQueryUseCase interface {
	FindPost(id uint) (*domain.Post, error)

	RetrievePosts(query string) []domain.Post
}

type PostRepository interface {
	CreatePost(post domain.Post) (*domain.Post, error)

	UpdatePost(id uint, post domain.Post) (*domain.Post, error)

	DeletePost(id uint) error

	FindPost(id uint) (*domain.Post, error)

	RetrievePosts(query string) []domain.Post
}
