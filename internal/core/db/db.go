package db

import (
	"github.com/meteormin/friday.go/internal/core/db/entity"
	"github.com/meteormin/friday.go/pkg/database"
	"github.com/meteormin/friday.go/pkg/logger"
)

func New(cfg database.Config) error {
	db := database.New(cfg)

	logger.GetLogger().Debug("Migrating database...")

	err := database.Migrate(db, &entity.User{})
	if err != nil {
		return err
	}

	return nil
}
