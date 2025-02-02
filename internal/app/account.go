package app

import "github.com/meteormin/friday.go/internal/domain"

type UserRepository interface {
	ExistsByUsername(username string) bool

	FindByID(id uint) (*domain.User, error)

	Fetch() []domain.User

	FindByUsername(username string) (*domain.User, error)

	Create(user *domain.User) (*domain.User, error)

	Update(id uint, user *domain.User) (*domain.User, error)

	Delete(id uint) error
}

type AccountCommandService struct {
	repo UserRepository
}

func (a *AccountCommandService) CreateUser(user domain.CreateUser) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountCommandService) UpdateUser(id uint, user domain.UpdateUser) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountCommandService) DeleteUser(id uint) error {
	//TODO implement me
	panic("implement me")
}

func NewAccountCommandService(repo UserRepository) domain.UserCommand {
	return &AccountCommandService{
		repo: repo,
	}
}

type AccountQueryService struct {
	repo UserRepository
}

func (a AccountQueryService) FindUser(id uint) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountQueryService) FindUserByUsername(username string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountQueryService) FetchUsers() []domain.User {
	//TODO implement me
	panic("implement me")
}

func NewAccountQueryService(repo UserRepository) domain.UserQuery {
	return &AccountQueryService{
		repo: repo,
	}
}
