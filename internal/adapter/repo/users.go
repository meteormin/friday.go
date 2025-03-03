package repo

import (
	"errors"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/core/db/entity"
	"github.com/meteormin/friday.go/internal/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u UserRepositoryImpl) ExistsByIsAdmin() bool {
	var count int64
	u.db.Model(&entity.User{}).Where("is_admin = ?", true).
		Count(&count)

	return count > 0
}

func (u UserRepositoryImpl) ExistsByUsername(username string) bool {

	var user entity.User

	tx := u.db.Where("username = ?", username).First(&user)

	return !errors.Is(tx.Error, gorm.ErrRecordNotFound)
}

func (u UserRepositoryImpl) Create(user *domain.User) (*domain.User, error) {

	ent := mapToUserEntity(user)

	u.db.Create(&ent)

	return mapToUserModel(ent), nil
}

func (u UserRepositoryImpl) Update(id uint, user *domain.User) (*domain.User, error) {

	var ent entity.User

	tx := u.db.First(&ent, id)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, tx.Error
	}

	ent.Name = user.Name
	ent.Password = user.Password

	if err := u.db.Model(&ent).Updates(ent).Error; err != nil {
		return nil, err
	}

	u.db.First(&ent, id)

	return mapToUserModel(ent), nil
}

func (u UserRepositoryImpl) Delete(id uint) error {

	var ent entity.User

	tx := u.db.First(&ent, id)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return tx.Error
	}

	return u.db.Delete(&ent).Error
}

func (u UserRepositoryImpl) FindByID(id uint) (*domain.User, error) {

	var user domain.User

	tx := u.db.First(&user, id)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, tx.Error
	}

	return &user, nil
}

func (u UserRepositoryImpl) Fetch() []domain.User {
	users := make([]entity.User, 0)

	u.db.Find(&users)

	models := make([]domain.User, len(users))
	for _, model := range users {
		models = append(models, *mapToUserModel(model))
	}

	return models
}

func (u UserRepositoryImpl) FindByUsername(username string) (*domain.User, error) {

	var user entity.User

	tx := u.db.Where("username = ?", username).First(&user)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, tx.Error
	}

	return mapToUserModel(user), nil
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func mapToUserEntity(user *domain.User) entity.User {
	return entity.User{
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
	}
}

func mapToUserModel(user entity.User) *domain.User {
	return &domain.User{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Password:  user.Password,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
