package port

import "github.com/meteormin/friday.go/internal/domain"

type CreateSite struct {
	Host  string
	Alias string
}

type UpdateSite struct {
	Alias string
}

type SiteCommandUseCase interface {
	CreateSite(site CreateSite) (*domain.Site, error)
	UpdateSite(id uint, site UpdateSite) (*domain.Site, error)
	DeleteSite(id uint) error
}

type SiteQueryUseCase interface {
	FindSite(id uint) (*domain.Site, error)

	RetrieveSite(query string) []domain.Site
}

type SiteRepository interface {
	CreateSite(site domain.Site) (*domain.Site, error)

	UpdateSite(id uint, site domain.Site) (*domain.Site, error)

	DeleteSite(id uint) error

	FindSite(id uint) (*domain.Site, error)

	RetrieveSite(query string) []domain.Site

	RetrievePostBySite(siteID uint) []domain.Post
}
