package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type Error struct {
	Code    int
	Title   string
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code int, title, message string) *Error {
	return &Error{
		Code:    code,
		Title:   title,
		Message: message,
	}
}

// 비밀번호 해시화
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// 비밀번호 검증
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
