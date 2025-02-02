package db

import (
	"github.com/meteormin/friday.go/internal/infra/db/entity"
	"github.com/meteormin/friday.go/pkg/database"
)

func New(cfg database.Config) error {
	db := database.New(cfg)
	err := database.Migrate(db, &entity.User{})
	if err != nil {
		return err
	}

	return nil
}
