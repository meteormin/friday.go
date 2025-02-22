package repo

import (
	"github.com/meteormin/friday.go/internal/core/db/entity"
	"github.com/meteormin/friday.go/pkg/database"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	log.Println("Setup test: repo")

	db = database.New(database.Config{
		Name: "test",
		Path: ":memory:",
	}).Debug()

	log.Println("Setup test: migrate database")

	err := db.AutoMigrate(
		&entity.File{},
		&entity.Site{},
		&entity.Post{},
		&entity.Tag{},
		&entity.User{},
	)

	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
