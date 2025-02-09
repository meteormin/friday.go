package app

import (
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/domain"
)

type SiteCommandService struct {
	repo port.SiteRepository
}

func (s *SiteCommandService) CreateSite(site port.CreateSite) (*domain.Site, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SiteCommandService) UpdateSite(id uint, site port.UpdateSite) (*domain.Site, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SiteCommandService) DeleteSite(id uint) error {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

func (s *SiteQueryService) RetrieveSite(query string) []domain.Site {
	//TODO implement me
	panic("implement me")
}

func NewSiteQueryService(repo port.SiteRepository) port.SiteQueryUseCase {
	return &SiteQueryService{
		repo: repo,
	}
}
