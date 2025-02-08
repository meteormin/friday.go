package config

import (
	"github.com/meteormin/friday.go/pkg/database"
	"github.com/meteormin/friday.go/pkg/logger"
	"github.com/meteormin/friday.go/pkg/scheduler"
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
	Data   string `yaml:"data"`
	Log    string `yaml:"log"`
	Secret string `yaml:"secret"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Jwt  struct {
		Secret string `yaml:"secret"`
		Exp    int    `yaml:"exp"`
	} `yaml:"jwt"`
}

type Config struct {
	Env       Env                    `yaml:"env"`
	TZ        string                 `yaml:"timeZone"`
	Path      Path                   `yaml:"path"`
	App       App                    `yaml:"app"`
	Server    Server                 `yaml:"server"`
	Database  database.Config        `yaml:"database"`
	Badger    database.BadgerConfig  `yaml:"badger"`
	Logging   logger.ZapLoggerConfig `yaml:"logging"`
	Scheduler scheduler.Config       `yaml:"scheduler"`
}
