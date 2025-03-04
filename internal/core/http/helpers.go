package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/meteormin/friday.go/internal/core"
	"github.com/meteormin/friday.go/internal/core/db/entity"
	"strconv"
	"time"
)

// GenerateToken JWT 토큰 생성 함수
func GenerateToken(username string, exp time.Duration, isAdmin bool) (string, error) {
	var user entity.User
	if err := core.DB().Where("username = ?", username).First(&user).Error; err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"id":       strconv.Itoa(int(user.ID)),
		"username": username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(exp).Unix(), // 24시간 유효
		"admin":    isAdmin,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(core.Config().Server.Jwt.Secret))
}

func ExtractTokenClaims(ctx *fiber.Ctx) jwt.MapClaims {
	user := ctx.Locals("user").(*jwt.Token)
	return user.Claims.(jwt.MapClaims)
}
