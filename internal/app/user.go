package app

import (
	"github.com/meteormin/friday.go/internal/app/errors"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/domain"
)

type UserCommandService struct {
	repo port.UserRepository
}

func (a *UserCommandService) CreateUser(user port.CreateUser) (*domain.User, error) {
	model, err := user.Valid()
	if err != nil {
		return nil, err
	}

	if a.repo.ExistsByUsername(model.Username) {
		return nil, errors.ErrDuplicateUserUsername
	}

	err = model.HashPassword()
	if err != nil {
		return nil, err
	}

	return a.repo.Create(model)
}

func (a *UserCommandService) UpdateUser(id uint, user port.UpdateUser) (*domain.User, error) {
	updateModel, err := user.Valid()
	if err != nil {
		return nil, err
	}

	model, err := a.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = model.Update(updateModel)
	if err != nil {
		return nil, err
	}

	return a.repo.Update(id, model)
}

func (a *UserCommandService) DeleteUser(id uint) error {
	return a.repo.Delete(id)
}

func NewUserCommandService(repo port.UserRepository) port.UserCommandUseCase {
	return &UserCommandService{
		repo: repo,
	}
}

type UserQueryService struct {
	repo port.UserRepository
}

func (a UserQueryService) FindUser(id uint) (*domain.User, error) {
	return a.repo.FindByID(id)
}

func (a UserQueryService) FindUserByUsername(username string) (*domain.User, error) {
	return a.repo.FindByUsername(username)
}

func (a UserQueryService) FetchUsers() []domain.User {
	return a.repo.Fetch()
}

func NewUserQueryService(repo port.UserRepository) port.UserQueryUseCase {
	return &UserQueryService{
		repo: repo,
	}
}
