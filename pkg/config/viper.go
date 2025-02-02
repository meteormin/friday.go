package config

import (
	"bytes"
	"github.com/spf13/viper"
	"log"
	"os"
)

func LoadWithViper(in string, appInfo App) *Config {

	if _, err := os.Stat(in); err == nil {
		viper.SetConfigFile(in)
		if err = viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
	} else {
		err = viper.ReadConfig(bytes.NewBufferString(in))
		if err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	cfg.App = appInfo

	return &cfg
}
