package domain

type User struct {
	ID       uint
	Name     string
	Username string
	Password string
}

type CreateUser struct {
	Name     string
	Username string
	Password string
}

type UpdateUser struct {
	Name     string
	Password string
}

type UserCommand interface {
	CreateUser(user CreateUser) User

	UpdateUser(id uint, user UpdateUser) User

	DeleteUser(id uint)
}

type UserQuery interface {
	FindUser(id uint) User

	FetchUsers() []User
}
