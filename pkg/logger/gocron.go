package logger

import "github.com/go-co-op/gocron/v2"

func NewGoCronLogger(level gocron.LogLevel) gocron.Logger {
	return gocron.NewLogger(level)
}
