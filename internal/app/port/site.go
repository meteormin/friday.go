package port

import (
	"github.com/meteormin/friday.go/internal/app/errors"
	"github.com/meteormin/friday.go/internal/domain"
)

type CreateSite struct {
	Host string
	Name string
}

func (c CreateSite) Valid() (*domain.Site, error) {
	if c.Host == "" {
		return nil, errors.ErrInvalidSiteHost
	}

	if c.Name == "" {
		return nil, errors.ErrInvalidSiteName
	}

	return &domain.Site{
		Host: c.Host,
		Name: c.Name,
	}, nil
}

type UpdateSite struct {
	Name string
}

func (u UpdateSite) Valid() (*domain.Site, error) {
	if u.Name == "" {
		return nil, errors.ErrInvalidSiteName
	}

	return &domain.Site{
		Name: u.Name,
	}, nil
}

type SiteCommandUseCase interface {
	CreateSite(site CreateSite) (*domain.Site, error)
	UpdateSite(id uint, site UpdateSite) (*domain.Site, error)
	DeleteSite(id uint) error
}

type SiteQueryUseCase interface {
	FindSite(id uint) (*domain.Site, error)

	RetrieveSite(query string) ([]domain.Site, error)
}

type SiteRepository interface {
	ExistsSiteByHost(host string) (bool, error)

	ExistsSiteByName(name string) (bool, error)

	CreateSite(site *domain.Site) (*domain.Site, error)

	UpdateSite(id uint, site *domain.Site) (*domain.Site, error)

	DeleteSite(id uint) error

	FindSite(id uint) (*domain.Site, error)

	RetrieveSite(query string) ([]domain.Site, error)

	RetrievePostBySite(siteID uint) ([]domain.Post, error)
}
