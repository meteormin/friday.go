package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/meteormin/friday.go/internal/app/errors"
	"github.com/meteormin/friday.go/internal/app/port"
	"github.com/meteormin/friday.go/internal/core"
	"github.com/meteormin/friday.go/internal/core/http"
	"time"
)

// SignInRequest
// @Description 로그인 요청
type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignupRequest
// @Description 가입 요청
type SignupRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// UpdateUserRequest
// @Description 회원 정보 수정 요청
type UpdateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// UserResource
// @Description 회원 정보 리소스
type UserResource struct {
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// TokenResource
// @Description 토큰 정보 리소스
type TokenResource struct {
	Token string `json:"token"`
	Exp   int    `json:"exp"`
}

type AuthHandler struct {
	useCase port.UserCommandUseCase
	query   port.UserQueryUseCase
}

// @Summary 회원 가입
// @Description 회원 가입 API
// @ID sign-up
// @Accept json
// @Produce json
// @Param req body SignupRequest true "회원 가입 정보"
// @Success 201 {object} UserResource "회원 가입 성공"
// @Failure 400 {object} errors.Error "잘못된 요청" errors.ErrInvalidUserName
// @Failure 400 {object} errors.Error "잘못된 요청" errors.ErrInvalidUserUsername
// @Failure 400 {object} errors.Error "잘못된 요청" errors.ErrInvalidUserPassword
// @Failure 409 {object} errors.Error "이메일 중복" errors.ErrDuplicateUserUsername
// @Failure 500 {object} errors.Error "서버 오류"
// @Router /auth/sign-up [post]
// @Tags auth
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

	return ctx.Status(fiber.StatusCreated).
		JSON(UserResource{
			Name:      user.Name,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
}

// @Summary 회원 로그인
// @Description 회원 로그인 API
// @ID sign-in
// @Accept json
// @Produce json
// @Param req body SignInRequest true "회원 로그인 정보"
// @Success 200 {object} TokenResource "회원 로그인 성공"
// @Failure 400 {object} errors.Error "잘못된 요청"
// @Failure 401 {object} errors.Error "로그인 실험"
// @Failure 500 {object} errors.Error "서버 오류"
// @Router /auth/sign-in [post]
// @Tags auth
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

	return ctx.JSON(TokenResource{
		Token: token,
		Exp:   exp,
	})
}

// @Summary 회원 정보 조회
// @Description 회원 정보 조회 API
// @ID me
// @Accept json
// @Produce json
// @Success 200 {object} UserResource "회원 정보 조회 성공"
// @Failure 401 {object} errors.Error "로그인 정보 없음"
// @Failure 500 {object} errors.Error "서버 오류"
// @Router /auth/me [get]
// @Tags auth
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

// @Summary 회원 리스트 조회
// @Description 회원 리스트 조회 API
// @ID users
// @Accept json
// @Produce json
// @Success 200 {array} UserResource "회원 리스트 조회 성공"
// @Failure 500 {object} errors.Error "서버 오류"
// @Router /users [get]
// @Tags users
func (handler *UserHandler) fetchUsers(ctx *fiber.Ctx) error {
	users := handler.query.FetchUsers()
	return ctx.JSON(users)
}

// NewUserHandler 회원 관리 Handler 생성
func NewUserHandler(command port.UserCommandUseCase, query port.UserQueryUseCase) http.AddRouteFunc {

	handler := &UserHandler{
		command: command,
		query:   query,
	}

	return func(router fiber.Router) {
		group := router.Group("/users")
		group.Get("/", handler.fetchUsers)
	}
}
