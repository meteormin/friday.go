package logger

import (
	"context"
	gLogger "gorm.io/gorm/logger"
	"time"
)

type gLoggerProxy struct {
	base gLogger.Interface
}

// LogMode description of the Go function.
//
// It takes a parameter of type logger.LogLevel and returns a logger.Interface.
func (l *gLoggerProxy) LogMode(level gLogger.LogLevel) gLogger.Interface {
	newLogger := *l
	newLogger.base = l.base.LogMode(level)
	return &newLogger
}

// Info logs information messages.
//
// ctx is the context for the logger.
// `s` is the message string.
// `i` is an optional list of interface{} for additional information.
func (l *gLoggerProxy) Info(ctx context.Context, s string, i ...interface{}) {
	l.base.Info(ctx, s, i...)
}

// Warn description of the Go function.
//
// ctx context.Context, s string, i ...interface{}
func (l *gLoggerProxy) Warn(ctx context.Context, s string, i ...interface{}) {
	l.base.Warn(ctx, s, i...)
}

// Error describes the Go function.
//
// It takes a context and a string, and a variadic list of interfaces as parameters.
// It does not return anything.
func (l *gLoggerProxy) Error(ctx context.Context, s string, i ...interface{}) {
	l.base.Error(ctx, s, i...)
}

// Trace description of the Go function.
//
// ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error.
func (l *gLoggerProxy) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	l.base.Trace(ctx, begin, fc, err)
}

// NewGormLoggerProxy creates a new logger proxy.
//
// Returns a logger.Interface.
func NewGormLoggerProxy(writer gLogger.Writer, config gLogger.Config) gLogger.Interface {
	baseLogger := gLogger.New(writer, config)
	return &gLoggerProxy{
		base: baseLogger,
	}
}
