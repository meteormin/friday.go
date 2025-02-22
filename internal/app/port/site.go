package port

import (
	"github.com/meteormin/friday.go/internal/app"
	"github.com/meteormin/friday.go/internal/domain"
)

type CreateSite struct {
	Host   string
	Name   string
	UserID uint
}

func (c CreateSite) Valid() (*domain.Site, error) {
	if c.Host == "" {
		return nil, app.ErrInvalidSiteHost
	}

	if c.Name == "" {
		return nil, app.ErrInvalidSiteName
	}

	return &domain.Site{
		Host: c.Host,
		Name: c.Name,
		User: domain.User{
			ID: c.UserID,
		},
	}, nil
}

type UpdateSite struct {
	Name string
}

func (u UpdateSite) Valid() (*domain.Site, error) {
	if u.Name == "" {
		return nil, app.ErrInvalidSiteName
	}

	return &domain.Site{
		Name: u.Name,
	}, nil
}

type SiteCommandUseCase interface {
	CreateSite(site CreateSite) (*domain.Site, error)
	UpdateSite(userId, id uint, site UpdateSite) (*domain.Site, error)
	DeleteSite(userId, id uint) error
}

type SiteQueryUseCase interface {
	FindSite(userId, id uint) (*domain.Site, error)

	RetrieveSite(userId uint, query string) ([]domain.Site, error)
}

type SiteRepository interface {
	ExistsSiteByHost(host string) (bool, error)

	ExistsSiteByName(name string) (bool, error)

	CreateSite(site *domain.Site) (*domain.Site, error)

	UpdateSite(id uint, site *domain.Site) (*domain.Site, error)

	DeleteSite(id uint) error

	FindSite(id uint) (*domain.Site, error)

	RetrieveSite(userID uint, query string) ([]domain.Site, error)

	RetrievePostBySite(userID, siteID uint) ([]domain.Post, error)

	HasAccessPermission(userID, siteID uint) (bool, error)
}
