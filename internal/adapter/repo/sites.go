package repo

import (
	"errors"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/core/db/entity"
	"github.com/meteormin/friday.go/internal/domain"
	"gorm.io/gorm"
)

type SiteRepositoryImpl struct {
	db *gorm.DB
}

func (s SiteRepositoryImpl) HasAccessPermission(userID, siteID uint) (bool, error) {
	var count int64
	if err := s.db.Where("user_id = ? AND id = ?", userID, siteID).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return count > 0, nil
}

func (s SiteRepositoryImpl) ExistsSiteByHost(host string) (bool, error) {
	var site entity.Site
	if err := s.db.Where("host = ?", host).First(&site).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (s SiteRepositoryImpl) ExistsSiteByName(name string) (bool, error) {
	var site entity.Site
	if err := s.db.Where("name = ?", name).First(&site).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (s SiteRepositoryImpl) CreateSite(site *domain.Site) (*domain.Site, error) {
	ent := mapToSiteEntity(site)
	if err := s.db.Create(&ent).Error; err != nil {
		return nil, err
	}

	return mapToSiteModel(ent), nil
}

func (s SiteRepositoryImpl) UpdateSite(id uint, site *domain.Site) (*domain.Site, error) {
	var ent entity.Site

	if err := s.db.First(&ent, id).Error; err != nil {
		return nil, err
	}

	ent.Name = site.Name
	ent.UserID = site.User.ID

	if err := s.db.Save(&ent).Error; err != nil {
		return nil, err
	}

	return mapToSiteModel(ent), nil
}

func (s SiteRepositoryImpl) DeleteSite(id uint) error {
	return s.db.Delete(&entity.Site{}, id).Error
}

func (s SiteRepositoryImpl) FindSite(id uint) (*domain.Site, error) {
	var ent entity.Site

	if err := s.db.First(&ent, id).Error; err != nil {
		return nil, err
	}

	return mapToSiteModel(ent), nil
}

func (s SiteRepositoryImpl) RetrieveSite(userID uint, query string) ([]domain.Site, error) {
	var sites []entity.Site

	tx := s.db.Preload("User", "user_id = ?", userID).
		Where("name LIKE ?", "%"+query+"% OR host LIKE ?", "%"+query+"%").
		Find(&sites)

	if err := tx.Error; err != nil {
		return make([]domain.Site, 0), err
	}

	var results []domain.Site
	for _, site := range sites {
		results = append(results, *mapToSiteModel(site))
	}

	return results, nil
}

func (s SiteRepositoryImpl) RetrievePostBySite(userID, siteID uint) ([]domain.Post, error) {
	var posts []entity.Post

	tx := s.db.Preload("User", "user_id = ?", userID).
		Preload("Posts").
		Preload("Posts.Tags").
		Where("site_id = ?", siteID).Find(&posts)

	if err := tx.Error; err != nil {
		return make([]domain.Post, 0), err
	}

	var results []domain.Post
	for _, post := range posts {
		results = append(results, *mapToPostModel(post))
	}

	return results, nil
}

func NewSiteRepository(db *gorm.DB) port.SiteRepository {
	return &SiteRepositoryImpl{db: db}
}

func mapToSiteModel(ent entity.Site) *domain.Site {
	return &domain.Site{
		ID:   ent.ID,
		Name: ent.Name,
		Host: ent.Host,
	}
}

func mapToSiteEntity(site *domain.Site) entity.Site {
	return entity.Site{
		Name:   site.Name,
		Host:   site.Host,
		UserID: site.User.ID,
	}
}
