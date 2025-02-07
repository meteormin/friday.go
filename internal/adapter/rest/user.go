package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/core"
	"github.com/meteormin/friday.go/internal/core/http"
	"time"
)

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserResource struct {
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type AuthHandler struct {
	useCase port.UserCommandUseCase
	query   port.UserQueryUseCase
}

func (auth *AuthHandler) signUp(ctx *fiber.Ctx) error {
	var req SignupRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, err := auth.useCase.CreateUser(port.CreateUser{
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
		return fiber.NewError(fiber.StatusUnauthorized, "Access Denied")
	}

	if exists.CheckPassword(req.Password) {
		return fiber.NewError(fiber.StatusUnauthorized, "Access Denied")
	}

	exp := core.GetConfig().Server.Jwt.Exp
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

func NewAuthHandler(useCase port.UserCommandUseCase, query port.UserQueryUseCase) http.AddRouteFunc {

	handler := &AuthHandler{
		useCase: useCase,
		query:   query,
	}

	return func(router fiber.Router) {
		group := router.Group("/auth")
		group.Post("/sign-up", handler.signUp)
		group.Post("/sign-in", handler.signIn)
		group.Get("/me", handler.me)
	}
}

type UserHandler struct {
	command port.UserCommandUseCase
	query   port.UserQueryUseCase
}

func (handler *UserHandler) getAll(ctx *fiber.Ctx) error {
	users := handler.query.FetchUsers()

	return ctx.JSON(users)
}

func NewUserHandler(command port.UserCommandUseCase, query port.UserQueryUseCase) http.AddRouteFunc {

	handler := &UserHandler{
		command: command,
		query:   query,
	}

	return func(router fiber.Router) {
		group := router.Group("/users")
		group.Get("/", handler.getAll)
	}
}
