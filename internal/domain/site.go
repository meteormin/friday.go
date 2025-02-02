package domain

type Site struct {
	ID    uint
	Host  string
	Alias string
}

type CreateSite struct {
	Host  string
	Alias string
}

type UpdateSite struct {
	Alias string
}

type SiteCommand interface {
	CreateSite(site CreateSite) Site
	UpdateSite(id uint, site UpdateSite) Site
	DeleteSite(id uint) Site
}

type SiteQueryDTO struct {
	Host  string
	Alias string
}

type SiteQuery interface {
	FindSite(id uint) Site

	RetrieveSite(query SiteQueryDTO) Site
}
