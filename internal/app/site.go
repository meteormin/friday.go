package app

import (
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

	return s.repo.CreateSite(model)
}

func (s *SiteCommandService) UpdateSite(id uint, site port.UpdateSite) (*domain.Site, error) {
	model, err := site.Valid()
	if err != nil {
		return nil, err
	}

	return s.repo.UpdateSite(id, model)
}

func (s *SiteCommandService) DeleteSite(id uint) error {
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

func (s *SiteQueryService) FindSite(id uint) (*domain.Site, error) {
	return s.repo.FindSite(id)
}

func (s *SiteQueryService) RetrieveSite(query string) ([]domain.Site, error) {
	return s.repo.RetrieveSite(query)
}

func NewSiteQueryService(repo port.SiteRepository) port.SiteQueryUseCase {
	return &SiteQueryService{
		repo: repo,
	}
}
