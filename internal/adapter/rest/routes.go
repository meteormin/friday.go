package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/meteormin/friday.go/internal/adapter/repo"
	"github.com/meteormin/friday.go/internal/adapter/rest/handler"
	"github.com/meteormin/friday.go/internal/app/service"
	"github.com/meteormin/friday.go/internal/core"
	"github.com/meteormin/friday.go/internal/core/http"
	"github.com/meteormin/friday.go/internal/core/http/middleware"
	"github.com/meteormin/friday.go/internal/core/task"
)

func authHandler(router fiber.Router) {
	userRepo := repo.NewUserRepository(core.GetDB())
	userCommand := service.NewUserCommandService(userRepo)
	userQuery := service.NewUserQueryService(userRepo)
	auth := handler.NewAuthHandler(userCommand, userQuery)
	auth(router)
}

func userHandler(router fiber.Router) {
	userRepo := repo.NewUserRepository(core.GetDB())
	userCommand := service.NewUserCommandService(userRepo)
	userQuery := service.NewUserQueryService(userRepo)
	users := handler.NewUserHandler(userCommand, userQuery)
	users(router)
}

func uploadFileHandler(router fiber.Router) {
	fileRepo := repo.NewFileRepository("uploads", core.GetDB(), core.GetStorage())
	fileCommand := service.NewUploadFileService(fileRepo)
	uploadFile := handler.NewUploadFileHandler(fileCommand)
	uploadFile(router)
}

func siteHandler(router fiber.Router) {
	siteRepo := repo.NewSiteRepository(core.GetDB())
	siteCommand := service.NewSiteCommandService(siteRepo)
	siteQuery := service.NewSiteQueryService(siteRepo)
	sites := handler.NewSiteHandler(siteCommand, siteQuery)
	sites(router)
}

func postHandler(router fiber.Router) {
	postRepo := repo.NewPostRepository(core.GetDB())
	siteRepo := repo.NewSiteRepository(core.GetDB())
	postCommand := service.NewPostCommandService(postRepo, siteRepo)
	postQuery := service.NewPostQueryService(postRepo)
	posts := handler.NewPostHandler(postCommand, postQuery)
	posts(router)
}

func RegisterRoutes() {
	http.Route("/api-docs", func(router fiber.Router) {
		router.Get("/swagger/*", swagger.HandlerDefault)
	})

	http.Middleware(middleware.NewCommon)
	http.Route("/api", func(router fiber.Router) {
		task.Handler(router)
		authHandler(router)
		middleware.NewJwtGuard(router)
		userHandler(router)
		uploadFileHandler(router)
		siteHandler(router)
		postHandler(router)
	})

	http.Middleware(middleware.EmbedUI)
}
