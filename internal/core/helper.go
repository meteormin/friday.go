package core

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/go-co-op/gocron/v2"
	"github.com/meteormin/friday.go/internal/core/config"
	"github.com/meteormin/friday.go/pkg/database"
	"github.com/meteormin/friday.go/pkg/logger"
	"github.com/meteormin/friday.go/pkg/scheduler"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var gCfg *config.Config

func SetConfig(cfg *config.Config) {
	gCfg = cfg
}

func GetConfig() *config.Config {
	return gCfg
}

func GetDB() *gorm.DB {
	if gCfg != nil {
		return database.GetDB(gCfg.Database.Name)
	}
	return database.GetDB()
}

func GetLogger() *zap.SugaredLogger {
	if gCfg != nil {
		return logger.GetLogger(gCfg.Logging.Name)
	}
	return logger.GetLogger()
}

func GetScheduler() gocron.Scheduler {
	return scheduler.GetScheduler()
}

var storage *badger.DB

func SetStorage(s *badger.DB) {
	storage = s
}

func GetStorage() *badger.DB {
	return storage
}
