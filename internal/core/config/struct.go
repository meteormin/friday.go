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
	Name    string `json:"name" yaml:"name"`
	Version string `json:"version" yaml:"version"`
}

type Path struct {
	Data   string `json:"data" yaml:"data"`
	Log    string `json:"log" yaml:"log"`
	Secret string `json:"secret" yaml:"secret"`
}

type Server struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	Jwt  struct {
		Secret string `json:"secret" yaml:"secret"`
		Exp    int    `json:"exp" yaml:"exp"`
	} `json:"jwt" yaml:"jwt"`
}

type Config struct {
	Env       Env                    `json:"env" yaml:"env"`
	TZ        string                 `json:"tz" yaml:"timeZone"`
	Path      Path                   `json:"path" yaml:"path"`
	App       App                    `json:"app" yaml:"app"`
	Server    Server                 `json:"server" yaml:"server"`
	Database  database.Config        `json:"database" yaml:"database"`
	Badger    database.BadgerConfig  `json:"badger" yaml:"badger"`
	Logging   logger.ZapLoggerConfig `json:"logging" yaml:"logging"`
	Scheduler scheduler.Config       `json:"scheduler" yaml:"scheduler"`
}
