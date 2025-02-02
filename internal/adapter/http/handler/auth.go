package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/meteormin/friday.go/internal/domain"
	"github.com/meteormin/friday.go/internal/infra"
	"github.com/meteormin/friday.go/internal/infra/http"
	"time"
)

type AuthHandler struct {
	command domain.UserCommand
	query   domain.UserQuery
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResource struct {
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (auth *AuthHandler) signUp(ctx *fiber.Ctx) error {
	var req SignupRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, err := auth.command.CreateUser(domain.CreateUser{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(UserResource{
		Name:      user.Name,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

func (auth *AuthHandler) signIn(ctx *fiber.Ctx) error {
	var req SignInRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	exists, err := auth.query.FindUserByUsername(req.Username)
	if err != nil {
		return err
	}

	err = exists.CheckPassword(req.Password)
	if err != nil {
		return err
	}

	exp := infra.GetConfig().Server.Jwt.Exp
	token, err := http.GenerateToken(req.Username,
		time.Second*time.Duration(exp))

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"token": token,
		"exp":   exp,
	})
}

func (auth *AuthHandler) me(ctx *fiber.Ctx) error {
	token, ok := ctx.Locals("user").(*jwt.Token)
	if token == nil || !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}
	username := token.Claims.(jwt.MapClaims)["username"].(string)

	user, err := auth.query.FindUserByUsername(username)
	if err != nil {
		return err
	}

	return ctx.JSON(UserResource{
		Name:      user.Name,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

func NewAuthHandler(command domain.UserCommand, query domain.UserQuery) http.AddRouteFunc {

	handler := &AuthHandler{
		command: command,
		query:   query,
	}

	return func(router fiber.Router) {
		group := router.Group("/auth")
		group.Post("/sign-in", handler.signIn)
		group.Get("/me", handler.me)
	}
}
