package app

import (
	apperrors "github.com/meteormin/friday.go/internal/app/errors"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/domain"
)

type PostCommandService struct {
	repo port.PostRepository
}

func (p PostCommandService) CreatePost(post port.CreatePost) (*domain.Post, error) {
	model, err := post.Valid()

	if err != nil {
		return nil, err
	}

	if exists, err := p.repo.ExistsPostByPath(model.Path); err != nil {
		return nil, err
	} else if exists {
		return nil, apperrors.ErrDuplicatePostPath
	}

	return p.repo.CreatePost(model)
}

func (p PostCommandService) UpdatePost(id uint, post port.UpdatePost) (*domain.Post, error) {
	model, err := post.Valid()

	if err != nil {
		return nil, err
	}

	return p.repo.UpdatePost(id, model)
}

func (p PostCommandService) DeletePost(id uint) error {
	return p.repo.DeletePost(id)
}

func NewPostCommandService(repo port.PostRepository) port.PostCommandUseCase {
	return &PostCommandService{repo: repo}
}

type PostQueryService struct {
	repo port.PostRepository
}

func (p PostQueryService) FindPost(id uint) (*domain.Post, error) {
	return p.repo.FindPost(id)
}

func (p PostQueryService) RetrievePosts(query string) ([]domain.Post, error) {
	return p.repo.RetrievePosts(query)
}

func NewPostQueryService(repo port.PostRepository) port.PostQueryUseCase {
	return &PostQueryService{repo: repo}
}
