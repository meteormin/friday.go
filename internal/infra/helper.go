package infra

import (
	"github.com/meteormin/friday.go/pkg/config"
	"github.com/meteormin/friday.go/pkg/database"
	"github.com/meteormin/friday.go/pkg/logger"
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
