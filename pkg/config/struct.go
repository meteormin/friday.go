package config

import (
	"friday.go/pkg/database"
	"friday.go/pkg/logger"
)

type App struct {
	Name    string `yaml:"-"`
	Version string `yaml:"-"`
	Mod     string `yaml:"-"`
}

type Path struct {
	Base string `yaml:"base"`
	Data string `yaml:"data"`
	Log  string `yaml:"log"`
}

type Config struct {
	App      App                    `yaml:"-"`
	Database database.Config        `yaml:"database"`
	Logging  logger.ZapLoggerConfig `yaml:"logging"`
	Path     Path                   `yaml:"path"`
}
