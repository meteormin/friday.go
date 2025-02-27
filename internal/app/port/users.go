package port

import (
	"github.com/meteormin/friday.go/internal/app"
	"github.com/meteormin/friday.go/internal/domain"
)

type CreateUser struct {
	Name     string
	Username string
	Password string
}

func (u CreateUser) Valid() (*domain.User, error) {
	if u.Name == "" || len(u.Name) < 4 {
		return nil, app.ErrInvalidUserName
	}

	if u.Username == "" || len(u.Username) < 4 {
		return nil, app.ErrInvalidUserUsername
	}

	if u.Password == "" || len(u.Password) < 8 {
		return nil, app.ErrInvalidUserPassword
	}

	return &domain.User{Name: u.Name, Username: u.Username, Password: u.Password}, nil
}

type UpdateUser struct {
	Name     string
	Password string
}

func (u UpdateUser) Valid() (*domain.User, error) {
	if u.Name == "" || len(u.Name) < 4 {
		return nil, app.ErrInvalidUserName
	}

	if u.Password == "" || len(u.Password) < 8 {
		return nil, app.ErrInvalidUserPassword
	}

	return &domain.User{Name: u.Name, Password: u.Password}, nil
}

type UserCommandUseCase interface {
	CreateUser(user CreateUser) (*domain.User, error)

	UpdateUser(id uint, user UpdateUser) (*domain.User, error)

	DeleteUser(id uint) error
}

type UserQueryUseCase interface {
	FindUser(id uint) (*domain.User, error)

	FindUserByUsername(username string) (*domain.User, error)

	FetchUsers() []domain.User

	HasAdmin() bool
}

type UserRepository interface {
	ExistsByUsername(username string) bool

	FindByID(id uint) (*domain.User, error)

	Fetch() []domain.User

	FindByUsername(username string) (*domain.User, error)

	Create(user *domain.User) (*domain.User, error)

	Update(id uint, user *domain.User) (*domain.User, error)

	Delete(id uint) error

	ExistsByIsAdmin() bool
}
