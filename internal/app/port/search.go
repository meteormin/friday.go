package port

import "github.com/meteormin/friday.go/internal/domain"

type SearchByKeyword struct {
	Keyword string
	Tags    []string
}

type SearchBySite struct {
	SiteID  uint
	Keyword string
	Tags    []string
}

type SearchUseCase interface {
	KeywordSearch(search SearchByKeyword) ([]domain.Post, error)

	SiteSearch(search SearchBySite) ([]domain.Post, error)
}
