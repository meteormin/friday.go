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

type PostQueryService struct {
	repo port.PostRepository
}

func (p PostQueryService) FindPost(id uint) (*domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostQueryService) RetrievePosts(query string) ([]domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func NewPostQueryService(repo port.PostRepository) port.PostQueryUseCase {
	return &PostQueryService{repo: repo}
}
