package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/meteormin/friday.go/internal/core"
	"github.com/meteormin/friday.go/internal/core/db/entity"
	"time"
)

type ContentResource[T interface{}] struct {
	Content []T `json:"content"`
	Length  int `json:"length"`
}

func NewContentResource[T interface{}](content []T) ContentResource[T] {
	return ContentResource[T]{
		Content: content,
		Length:  len(content),
	}
}

// GenerateToken JWT 토큰 생성 함수
func GenerateToken(username string, exp time.Duration, isAdmin bool) (string, error) {
	var user entity.User
	if err := core.GetDB().Where("username = ?", username).First(&user).Error; err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(exp).Unix(), // 24시간 유효
		"admin":    isAdmin,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(core.GetConfig().Server.Jwt.Secret)
}

func ExtractTokenClaims(ctx *fiber.Ctx) jwt.MapClaims {
	user := ctx.Locals("user").(*jwt.Token)
	return user.Claims.(jwt.MapClaims)
}
