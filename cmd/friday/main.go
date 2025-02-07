package main

import (
	"crypto/rand"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/meteormin/friday.go/internal/adapter/repo"
	"github.com/meteormin/friday.go/internal/adapter/rest"
	"github.com/meteormin/friday.go/internal/app"
	"github.com/meteormin/friday.go/internal/core"
	"github.com/meteormin/friday.go/internal/core/db"
	"github.com/meteormin/friday.go/internal/core/http"
	"github.com/meteormin/friday.go/internal/core/http/middleware"
	"github.com/meteormin/friday.go/internal/core/task"
	"github.com/meteormin/friday.go/pkg/config"
	"github.com/meteormin/friday.go/pkg/database"
	"github.com/meteormin/friday.go/pkg/logger"
	"github.com/meteormin/friday.go/pkg/scheduler"
	"os"
	"os/signal"
	"path"
	"strconv"
	"syscall"
	"time"
)

const (
	appName    = "Friday.go"
	appVersion = "0.0.1"
)

func init() {
	cfg := config.LoadWithViper("config.yml", config.App{
		Name:    appName,
		Version: appVersion,
	})

	core.SetConfig(cfg)

	l := logger.NewZapLogger(cfg.Logging)
	l.Info("Initializing application...")

	if err := db.New(cfg.Database); err != nil {
		l.Fatal(err)
	}

	l.Info("Database connection established.")

	if len(cfg.Badger.EncryptKey) == 0 {
		l.Info("Generating Encryption Key...")
		encKey, err := generateEncryptionKey(cfg.Path.Secret)
		if err != nil {
			l.Fatal(err)
		}
		cfg.Badger.EncryptKey = encKey
	}

	storage, err := database.NewBadger(cfg.Badger)

	core.SetStorage(storage)

	l.Info("Storage connection established.")

	loc, err := time.LoadLocation(cfg.TZ)
	if err != nil {
		l.Fatal(err)
	}

	cfg.Scheduler.Location = loc

	jobRepo := task.NewJobRepository(core.GetDB())
	cfg.Scheduler.Monitor = jobRepo
	cfg.Scheduler.MonitorStatus = jobRepo

	err = scheduler.New(cfg.Scheduler)
	if err != nil {
		l.Fatal(err)
	}

	http.NewFiber(fiber.Config{
		CaseSensitive:     true,
		AppName:           appName + " v" + appVersion,
		ErrorHandler:      http.NewErrorHandler(),
		EnablePrintRoutes: true,
	})
}

func authHandler() http.AddRouteFunc {
	userRepo := repo.NewUserRepository(core.GetDB())
	userCommand := app.NewUserCommandService(userRepo)
	userQuery := app.NewUserQueryService(userRepo)
	return rest.NewAuthHandler(userCommand, userQuery)
}

func userHandler() http.AddRouteFunc {
	userRepo := repo.NewUserRepository(core.GetDB())
	userCommand := app.NewUserCommandService(userRepo)
	userQuery := app.NewUserQueryService(userRepo)
	return rest.NewUserHandler(userCommand, userQuery)
}

// generateEncryptionKey 32바이트(256비트)의 랜덤 암호화 키를 생성합니다.
func generateEncryptionKey(secretPath string) ([]byte, error) {
	secretFile := path.Join(secretPath, ".secret")
	if _, err := os.Stat(secretFile); err == nil {
		key, err := os.ReadFile(secretFile)
		if err == nil {
			core.GetLogger().Debug("Exists Encryption Key")
			return key, nil
		}
	}

	key := make([]byte, 32) // AES-256에 필요한 32바이트 키
	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("failed generating encrypt ley: %w", err)
	}

	err = os.MkdirAll(secretPath, 0755)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(secretFile, key, 0644)
	if err != nil {
		return nil, err
	}

	core.GetLogger().Debug("Generated Encryption Key")
	return key, nil
}

func main() {

	cfg := core.GetConfig()
	l := core.GetLogger()

	http.Middleware(middleware.NewCommon, "/api")
	http.Route("/api", func(router fiber.Router) {
		auth := authHandler()
		user := userHandler()

		auth(router)
		middleware.NewJwtGuard(router)
		user(router)
	})

	if cfg.Server.Port <= 0 {
		cfg.Server.Port = 8080
	}

	fiberApp := http.Fiber()
	taskScheduler := core.GetScheduler()

	// Listen from a different goroutine
	go func() {
		taskScheduler.Start()
		if err := fiberApp.Listen(":" + strconv.Itoa(cfg.Server.Port)); err != nil {
			l.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = fiberApp.Shutdown()

	fmt.Println("Running cleanup tasks...")

	err := core.GetScheduler().Shutdown()
	if err != nil {
		l.Fatal(err)
	}

	fmt.Println("Fiber was successful shutdown.")
}
