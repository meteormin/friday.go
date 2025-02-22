package service

import (
	"github.com/meteormin/friday.go/internal/app"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/domain"
)

type SiteCommandService struct {
	repo port.SiteRepository
}

func (s *SiteCommandService) CreateSite(site port.CreateSite) (*domain.Site, error) {
	model, err := site.Valid()
	if err != nil {
		return nil, err
	}

	if exists, err := s.repo.ExistsSiteByHost(model.Host); err != nil {
		return nil, err
	} else if exists {
		return nil, app.ErrDuplicateSiteHost
	}

	if exists, err := s.repo.ExistsSiteByName(model.Name); err != nil {
		return nil, err
	} else if exists {
		return nil, app.ErrDuplicateSiteName
	}

	return s.repo.CreateSite(model)
}

func (s *SiteCommandService) UpdateSite(userId, id uint, site port.UpdateSite) (*domain.Site, error) {
	model, err := site.Valid()
	if err != nil {
		return nil, err
	}

	if hasPerm, err := s.repo.HasAccessPermission(userId, id); err != nil {
		return nil, err
	} else if !hasPerm {
		return nil, app.ErrForbidden
	}

	findModel, err := s.repo.FindSite(id)
	if err != nil {
		return nil, err
	}

	findModel.Update(*model)

	return s.repo.UpdateSite(id, findModel)
}

func (s *SiteCommandService) DeleteSite(userID, id uint) error {
	if hasPerm, err := s.repo.HasAccessPermission(userID, id); err != nil {
		return err
	} else if !hasPerm {
		return app.ErrForbidden
	}

	return s.repo.DeleteSite(id)
}

func NewSiteCommandService(repo port.SiteRepository) port.SiteCommandUseCase {
	return &SiteCommandService{
		repo: repo,
	}
}

type SiteQueryService struct {
	repo port.SiteRepository
}

func (s *SiteQueryService) FindSite(userId, id uint) (*domain.Site, error) {
	if hasPerm, err := s.repo.HasAccessPermission(userId, id); err != nil {
		return nil, err
	} else if !hasPerm {
		return nil, app.ErrForbidden
	}

	return s.repo.FindSite(id)
}

func (s *SiteQueryService) RetrieveSite(userId uint, query string) ([]domain.Site, error) {
	return s.repo.RetrieveSite(userId, query)
}

func NewSiteQueryService(repo port.SiteRepository) port.SiteQueryUseCase {
	return &SiteQueryService{
		repo: repo,
	}
}
