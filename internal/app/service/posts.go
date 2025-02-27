package service

import (
	"github.com/meteormin/friday.go/internal/app"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/domain"
)

type PostCommandService struct {
	repo     port.PostRepository
	siteRepo port.SiteRepository
}

func (p PostCommandService) CreatePost(userId uint, post port.CreatePost) (*domain.Post, error) {
	model, err := post.Valid()

	if err != nil {
		return nil, err
	}

	if hasPerm, err := p.siteRepo.HasAccessPermission(userId, model.SiteID); err != nil {
		return nil, err
	} else if !hasPerm {
		return nil, app.ErrForbidden
	}

	if exists, err := p.repo.ExistsPostByPath(model.SiteID, model.Path); err != nil {
		return nil, err
	} else if exists {
		return nil, app.ErrDuplicatePostPath
	}

	return p.repo.CreatePost(model)
}

func (p PostCommandService) UpdatePost(userId, id uint, post port.UpdatePost) (*domain.Post, error) {
	model, err := post.Valid()
	if err != nil {
		return nil, err
	}

	if hasPerm, err := p.repo.HasAccessPermission(userId, id); err != nil {
		return nil, err
	} else if !hasPerm {
		return nil, app.ErrForbidden
	}

	findPost, err := p.repo.FindPost(id)
	if err != nil {
		return nil, err
	}

	findPost.Update(*model)

	updated, err := p.repo.UpdatePost(id, findPost)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (p PostCommandService) DeletePost(userId, id uint) error {
	if hasPerm, err := p.repo.HasAccessPermission(userId, id); err != nil {
		return err
	} else if !hasPerm {
		return app.ErrForbidden
	}

	return p.repo.DeletePost(id)
}

func NewPostCommandService(repo port.PostRepository, siteRepo port.SiteRepository) port.PostCommandUseCase {
	return &PostCommandService{repo: repo, siteRepo: siteRepo}
}

type PostQueryService struct {
	repo     port.PostRepository
	siteRepo port.SiteRepository
}

func (p PostQueryService) FindPost(userId, id uint) (*domain.Post, error) {

	if hasPerm, err := p.repo.HasAccessPermission(userId, id); err != nil {
		return nil, err
	} else if !hasPerm {
		return nil, app.ErrForbidden
	}

	return p.repo.FindPost(id)
}

func (p PostQueryService) RetrievePosts(userId uint, query string) ([]domain.Post, error) {
	return p.repo.RetrievePosts(userId, query)
}

func NewPostQueryService(repo port.PostRepository) port.PostQueryUseCase {
	return &PostQueryService{repo: repo}
}
