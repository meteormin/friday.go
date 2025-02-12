package boot

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/meteormin/friday.go/internal/adapter/repo"
	"github.com/meteormin/friday.go/internal/adapter/rest"
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
	auth := rest.NewAuthHandler(userCommand, userQuery)
	auth(router)
}

func userHandler(router fiber.Router) {
	userRepo := repo.NewUserRepository(core.GetDB())
	userCommand := service.NewUserCommandService(userRepo)
	userQuery := service.NewUserQueryService(userRepo)
	users := rest.NewUserHandler(userCommand, userQuery)
	users(router)
}

func uploadFileHandler(router fiber.Router) {
	fileRepo := repo.NewFileRepository("uploads", core.GetDB(), core.GetStorage())
	fileCommand := service.NewUploadFileService(fileRepo)
	uploadFile := rest.NewUploadFileHandler(fileCommand)
	uploadFile(router)
}

func siteHandler(router fiber.Router) {
	siteRepo := repo.NewSiteRepository(core.GetDB())
	siteCommand := service.NewSiteCommandService(siteRepo)
	siteQuery := service.NewSiteQueryService(siteRepo)
	sites := rest.NewSiteHandler(siteCommand, siteQuery)
	sites(router)
}

func postHandler(router fiber.Router) {
	postRepo := repo.NewPostRepository(core.GetDB())
	postCommand := service.NewPostCommandService(postRepo)
	postQuery := service.NewPostQueryService(postRepo)
	posts := rest.NewPostHandler(postCommand, postQuery)
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
