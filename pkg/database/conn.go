package database

import (
	"fmt"
	"github.com/meteormin/friday.go/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"path"
)

// db is a map of database connections.
var db = make(map[string]*gorm.DB)

type Config struct {
	gorm.Config `yaml:"-" json:"-"`
	Name        string       `yaml:"name" json:"name"`
	Path        string       `yaml:"path" json:"path"`
	Debug       bool         `yaml:"debug" json:"debug"`
	Logger      LoggerConfig `yaml:"logger" json:"logger"`
}

type LoggerConfig struct {
	gLogger.Config
	LogPath string
}

type HasWithMigrate interface {
	WithMigrate(tx *gorm.DB) error
}

func newGormLogger(name string, cfg LoggerConfig) gLogger.Interface {
	var writer io.Writer
	if cfg.LogPath != "" {
		logPath := path.Join(cfg.LogPath,
			fmt.Sprintf("%s-%s.%s", "gorm", name, "log"))
		file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
		if err != nil {
			log.Printf("Failed to open log file: %s\n", err)
			log.Println("Gorm logging only stdout.")
			writer = os.Stdout
		} else {
			writer = io.MultiWriter(os.Stdout, file)
		}

		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				log.Printf("Failed to close log file: %s\n", err)
			}
		}(file)

	} else {
		writer = os.Stdout
	}

	return logger.NewGormLoggerProxy(log.New(writer, "\r\n", log.LstdFlags), cfg.Config)
}

var defaultConfig = Config{
	Name:   "default",
	Path:   "file::memory:?cache=shared",
	Debug:  false,
	Logger: LoggerConfig{},
}

func resolveConfig(cfg Config) Config {
	if cfg.Name == "" {
		cfg.Name = defaultConfig.Name
	}

	if cfg.Path == "" {
		cfg.Path = defaultConfig.Path
	}

	return cfg
}

// New initializes the database with the given name and returns a pointer to the gorm.DB object.
//
// dbname: The name of the database to initialize.
// Returns:
//   - *gorm.DB: A pointer to the gorm.DB object representing the initialized database.
func New(cfg Config) *gorm.DB {
	cfg = resolveConfig(cfg)

	if exists, ok := db[cfg.Name]; ok {
		return exists
	}

	if cfg.Config.Logger == nil {
		cfg.Config.Logger = newGormLogger(cfg.Name, cfg.Logger)
	}

	d, err := gorm.Open(sqlite.Open(cfg.Path), &gorm.Config{
		Logger:      cfg.Config.Logger,
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}

	if cfg.Debug {
		d = d.Debug()
	}

	db[cfg.Name] = d

	return db[cfg.Name]
}

// GetDB returns the gorm.DB object associated with the given database name.
//
// dbname: the name of the database
// returns: the gorm.DB object
func GetDB(dbname ...string) *gorm.DB {
	if len(dbname) > 0 {
		if exists, ok := db[dbname[0]]; ok {
			return exists
		}

		return nil
	}

	for _, conn := range db {
		return conn
	}

	return nil
}

// Migrate migrates the given models in the database.
//
// db *gorm.DB, models ...interface{}
func Migrate(db *gorm.DB, models ...interface{}) error {
	err := db.AutoMigrate(models...)
	if err != nil {
		return err
	}

	for _, model := range models {
		if m, ok := model.(HasWithMigrate); ok {
			err = m.WithMigrate(db)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Drop drops the specified table from the database.
//
// db: a pointer to the gorm.DB instance.
// There is no return value for this function.
func Drop(db *gorm.DB, models ...interface{}) error {
	return db.Migrator().DropTable(models...)
}
