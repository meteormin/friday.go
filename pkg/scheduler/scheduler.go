package scheduler

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/meteormin/friday.go/pkg/logger"
	"time"
)

var scheduler gocron.Scheduler

type Config struct {
	Location      *time.Location       `yaml:"-" json:"-"`
	Monitor       gocron.Monitor       `yaml:"-" json:"-"`
	MonitorStatus gocron.MonitorStatus `yaml:"-" json:"-"`

	LogLevel gocron.LogLevel `yaml:"logLevel" json:"logLevel"`
}

func (cfg Config) options() []gocron.SchedulerOption {
	return []gocron.SchedulerOption{
		gocron.WithLocation(cfg.Location),
		gocron.WithLogger(logger.NewGoCronLogger(cfg.LogLevel)),
		gocron.WithMonitor(cfg.Monitor),
		gocron.WithMonitorStatus(cfg.MonitorStatus),
	}
}

func New(cfg Config) error {
	s, err := gocron.NewScheduler(cfg.options()...)
	if err != nil {
		return err
	}

	scheduler = s

	return nil
}

func GetScheduler() gocron.Scheduler {
	return scheduler
}
