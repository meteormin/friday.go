package http

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/meteormin/friday.go/internal/core"
	"time"
)

// GenerateToken JWT 토큰 생성 함수
func GenerateToken(username string, exp time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(exp).Unix(), // 24시간 유효
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(core.GetConfig().Server.Jwt.Secret)
}
