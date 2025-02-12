package boot

import (
	"crypto/rand"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/meteormin/friday.go/internal/core"
	"github.com/meteormin/friday.go/internal/core/config"
	"github.com/meteormin/friday.go/internal/core/db"
	"github.com/meteormin/friday.go/internal/core/http"
	"github.com/meteormin/friday.go/internal/core/task"
	"github.com/meteormin/friday.go/pkg/database"
	"github.com/meteormin/friday.go/pkg/logger"
	"github.com/meteormin/friday.go/pkg/scheduler"
	"os"
	"path"
	"time"
)

const (
	appName    = "Friday.go"
	appVersion = "0.0.1"
)

func Boot() {
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
	if err != nil {
		l.Fatal(err)
	}

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
		EnablePrintRoutes: cfg.Env != config.Release,
	})
}

// generateEncryptionKey 32바이트(256비트)의 랜덤 암호화 키를 생성합니다.
func generateEncryptionKey(secretPath string) ([]byte, error) {
	secretFile := path.Join(secretPath, ".secret")
	if _, err := os.Stat(secretFile); err == nil {
		key, err := os.ReadFile(secretFile)
		if err == nil {
			core.GetLogger().Debug("Exists Encryption Key...")
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
