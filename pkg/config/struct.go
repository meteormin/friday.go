package config

import (
	"github.com/meteormin/friday.go/pkg/database"
	"github.com/meteormin/friday.go/pkg/logger"
)

type Env string

const (
	Test    Env = "test"
	Dev     Env = "dev"
	Release Env = "release"
)

type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Path struct {
	Base string `yaml:"base"`
	Data string `yaml:"data"`
	Log  string `yaml:"log"`
}

type Server struct {
	Port int `yaml:"port"`
	Jwt  struct {
		Secret string `yaml:"secret"`
		Exp    int    `yaml:"exp"`
	} `yaml:"jwt"`
}

type Config struct {
	Env      Env                    `yaml:"env"`
	App      App                    `yaml:"app"`
	Server   Server                 `yaml:"server"`
	Database database.Config        `yaml:"database"`
	Logging  logger.ZapLoggerConfig `yaml:"logging"`
	Path     Path                   `yaml:"path"`
}
