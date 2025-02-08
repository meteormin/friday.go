package app

import (
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/domain"
)

type PostCommandService struct {
	repo port.PostRepository
}

func (p PostCommandService) CreatePost(post port.CreatePost) (*domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostCommandService) UpdatePost(id uint, post port.UpdatePost) (*domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostCommandService) DeletePost(id uint) error {
	//TODO implement me
	panic("implement me")
}

func NewPostCommandService(repo port.PostRepository) port.PostCommandUseCase {
	return &PostCommandService{repo: repo}
}
