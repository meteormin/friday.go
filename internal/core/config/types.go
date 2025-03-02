package config

import (
	"github.com/meteormin/friday.go/pkg/database"
	"github.com/meteormin/friday.go/pkg/logger"
	"github.com/meteormin/friday.go/pkg/scheduler"
	"os"
)

type Env string

const (
	Test    Env = "test"
	Dev     Env = "dev"
	Release Env = "release"
)

type App struct {
	Name             string `json:"name" yaml:"name"`
	Version          string `json:"version" yaml:"version"`
	CaseSensitive    bool   `json:"caseSensitive" yaml:"caseSensitive"`
	EnablePrintRouts bool   `json:"enablePrintRouts" yaml:"enablePrintRouts"`
}

type Path struct {
	Data   string `json:"data" yaml:"data"`
	Log    string `json:"log" yaml:"log"`
	Secret string `json:"secret" yaml:"secret"`
}

func (p Path) mkdir() error {
	if p.Data == "" {
		p.Data = "./data"
	}

	if p.Log == "" {
		p.Log = "./logs"
	}

	if p.Secret == "" {
		p.Secret = "./secret"
	}

	if err := mkdirAll(p.Data); err != nil {
		return err
	}

	if err := mkdirAll(p.Log); err != nil {
		return err
	}

	if err := mkdirAll(p.Secret); err != nil {
		return err
	}

	return nil
}

func mkdirAll(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(path, 0755); err != nil {
				return err
			}
		}
	}
	return nil
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
	Env          Env                    `json:"env" yaml:"env"`
	TZ           string                 `json:"tz" yaml:"timeZone"`
	Path         Path                   `json:"path" yaml:"path"`
	App          App                    `json:"app" yaml:"app"`
	Server       Server                 `json:"server" yaml:"server"`
	Database     database.Config        `json:"database" yaml:"database"`
	Badger       database.BadgerConfig  `json:"badger" yaml:"badger"`
	Logging      logger.ZapLoggerConfig `json:"logging" yaml:"logging"`
	Scheduler    scheduler.Config       `json:"scheduler" yaml:"scheduler"`
	InMemoryMode bool                   `json:"inMemoryMode" yaml:"inMemoryMode"`
}
