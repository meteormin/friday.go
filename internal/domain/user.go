package domain

import (
	"time"
)

type User struct {
	ID        uint
	Name      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) CheckPassword(inputPassword string) bool {
	return !checkPasswordHash(inputPassword, u.Password)
}

func (u *User) HashPassword() error {
	hashed, err := hashPassword(u.Password)
	u.Password = hashed

	return err
}

func (u *User) Update(update *User) error {
	if update.Name != "" {
		u.Name = update.Name
	}

	if update.Username != "" {
		u.Username = update.Username
	}

	if update.Password != "" {

		u.Password = update.Password

		err := u.HashPassword()
		if err != nil {
			return err
		}
	}

	return nil
}
