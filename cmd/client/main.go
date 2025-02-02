package main

import (
	"friday.go/pkg/config"
)

const (
	APP_NAME    = "friday-client"
	APP_VERSION = "0.0.1"
)

func main() {
	cfg := config.LoadWithViper(config.App{
		Name:    APP_NAME,
		Version: APP_VERSION,
		Mod:     "client",
	})

}
