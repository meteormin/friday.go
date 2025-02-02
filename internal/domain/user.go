package domain

import (
	"errors"
	"time"
)

var (
	ErrNotFoundUser        = errors.New("user not found")
	ErrInvalidUserName     = errors.New("invalid name")
	ErrInvalidUserUsername = errors.New("invalid username")
	ErrInvalidUserPassword = errors.New("invalid password")
)

type User struct {
	ID        uint
	Name      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) CheckPassword(inputPassword string) error {
	if !checkPasswordHash(inputPassword, u.Password) {
		return ErrInvalidUserPassword
	}

	return nil
}

func (u *User) Update(update *User) {
	if update.Name != "" {
		u.Name = update.Name
	}

	if update.Username != "" {
		u.Username = update.Username
	}

	if update.Password != "" {
		u.Password = update.Password
	}
}

type CreateUser struct {
	Name     string
	Username string
	Password string
}

func (u CreateUser) Valid() (*User, error) {
	if u.Name == "" || len(u.Name) < 4 {
		return nil, ErrInvalidUserName
	}

	if u.Username == "" || len(u.Username) < 4 {
		return nil, ErrInvalidUserUsername
	}

	if u.Password == "" || len(u.Password) < 8 {
		return nil, ErrInvalidUserPassword
	}

	hashedPassword, err := hashPassword(u.Password)
	if err != nil {
		return nil, err
	}

	return &User{Name: u.Name, Username: u.Username, Password: hashedPassword}, nil
}

type UpdateUser struct {
	Name     string
	Password string
}

func (u UpdateUser) Valid() (*User, error) {
	if u.Name == "" || len(u.Name) < 4 {
		return nil, ErrInvalidUserName
	}

	if u.Password == "" || len(u.Password) < 8 {
		return nil, ErrInvalidUserPassword
	}

	return &User{Name: u.Name, Password: u.Password}, nil
}

type UserCommand interface {
	CreateUser(user CreateUser) (*User, error)

	UpdateUser(id uint, user UpdateUser) (*User, error)

	DeleteUser(id uint) error
}

type UserQuery interface {
	FindUser(id uint) (*User, error)

	FindUserByUsername(username string) (*User, error)

	FetchUsers() []User
}
