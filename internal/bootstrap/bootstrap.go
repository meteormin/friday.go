package bootstrap

import (
	"github.com/meteormin/friday.go/internal/core"
	"github.com/meteormin/friday.go/internal/core/config"
	"github.com/meteormin/friday.go/internal/core/db"
	"github.com/meteormin/friday.go/internal/core/task"
	"github.com/meteormin/friday.go/pkg/database"
	"github.com/meteormin/friday.go/pkg/logger"
	"github.com/meteormin/friday.go/pkg/scheduler"
	"time"
)

const (
	appName    = "Friday.go"
	appVersion = "0.0.1"
)

func Initialize(cfgPath string) *config.Config {

	cfg := config.LoadWithViper(cfgPath, config.App{
		Name:    appName,
		Version: appVersion,
	})

	core.SetConfig(cfg)

	if cfg.Env == config.Test {
		cfg.InMemoryMode = true
	}

	l := logger.NewZapLogger(cfg.Logging)
	l.Info("Initializing application...")

	if err := db.New(cfg.Database); err != nil {
		l.Fatal(err)
	}

	l.Info("Database connection established.")

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

	jobRepo := task.NewJobRepository(core.DB())
	cfg.Scheduler.Monitor = jobRepo
	cfg.Scheduler.MonitorStatus = jobRepo

	err = scheduler.New(cfg.Scheduler)
	if err != nil {
		l.Fatal(err)
	}

	if cfg.App.EnablePrintRouts {
		cfg.App.EnablePrintRouts = cfg.Env != config.Release
	}

	return cfg
}
