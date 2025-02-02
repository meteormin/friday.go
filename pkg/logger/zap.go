package logger

import (
	"os"
	"path"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// ZapLoggerConfig is custom logger configuration
type ZapLoggerConfig struct {
	Name        string        `yaml:"name" json:"name"`             // logger name
	TimeFormat  string        `yaml:"timeFormat" json:"timeFormat"` // time format
	FilePath    string        `yaml:"filePath" json:"filePath"`     // log file path
	Filename    string        `yaml:"filename" json:"filename"`     // log file name
	MaxSize     int           `yaml:"maxSize" json:"maxSize"`       // max log file size
	MaxBackups  int           `yaml:"maxBackups" json:"maxBackups"` // max log file backups
	MaxAge      int           `yaml:"maxAge" json:"maxAge"`         // max log file age
	Compress    bool          `yaml:"compress" json:"compress"`     // compress log status
	TimeKey     string        `yaml:"timeKey" json:"timeKey"`       // time key
	TimeZone    string        `yaml:"timeZone" json:"timeZone"`     // time zone
	LogLevel    zapcore.Level `yaml:"logLevel" json:"logLevel"`     // log level
	WithOptions []zap.Option  `yaml:"-" json:"-"`                   // zap.Option
}

// defaultZapLoggerConfig is default logger configuration
var defaultZapLoggerConfig = ZapLoggerConfig{
	Name:       "default",
	TimeFormat: "2006-01-02 15:04:05",
	FilePath:   "",
	Filename:   "",
	MaxSize:    10,
	MaxBackups: 30,
	MaxAge:     1,
	Compress:   false,
	TimeKey:    "timestamp",
	TimeZone:   "",
	LogLevel:   zapcore.DebugLevel,
}

// resolveConfig returns the default configuration with optional overrides.
//
// It takes a variable number of Config parameters and returns a Config type.
func resolveZapLoggerConfig(config ...ZapLoggerConfig) ZapLoggerConfig {
	if len(config) < 1 {
		return defaultZapLoggerConfig
	}

	cfg := config[0]

	if cfg.TimeFormat == "" {
		cfg.TimeFormat = defaultZapLoggerConfig.TimeFormat
	}

	if cfg.TimeKey == "" {
		cfg.TimeKey = defaultZapLoggerConfig.TimeKey
	}

	if cfg.TimeZone == "" {
		cfg.TimeZone, _ = time.Now().Zone()
	}

	if cfg.MaxSize == 0 {
		cfg.MaxSize = defaultZapLoggerConfig.MaxSize
	}

	if cfg.MaxAge == 0 {
		cfg.MaxAge = defaultZapLoggerConfig.MaxAge
	}

	if cfg.MaxBackups == 0 {
		cfg.MaxBackups = defaultZapLoggerConfig.MaxBackups
	}

	return cfg
}

// Default is default logger name
const Default string = "default"

// gLoggers is global logger map
var gLoggers = make(map[string]*zap.SugaredLogger)

// GetLogger returns a *zap.SugaredLogger.
//
// It takes an optional loggerName string parameter and returns a *zap.SugaredLogger.
func GetLogger(loggerName ...string) *zap.SugaredLogger {
	var logger *zap.SugaredLogger
	if gLoggers == nil {
		return NewZapLogger()
	}

	if len(loggerName) == 0 {
		if gLoggers[Default] != nil {
			logger = gLoggers[Default]
		} else {
			for _, l := range gLoggers {
				logger = l
				break
			}
		}
	} else {
		if _, ok := gLoggers[loggerName[0]]; ok {
			logger = gLoggers[loggerName[0]].Named(loggerName[0])
		}
	}

	if logger == nil {
		logger = NewZapLogger()
	}

	return logger
}

// NewZapLogger initializes and returns a new sugared logger.
//
// It accepts a variable number of ZapLoggerConfig options.
// Returns a pointer to a zap.SugaredLogger.
func NewZapLogger(config ...ZapLoggerConfig) *zap.SugaredLogger {
	cfg := resolveZapLoggerConfig(config...)

	logFilename := path.Join(cfg.FilePath, cfg.Filename)

	ll := &lumberjack.Logger{
		Filename:   logFilename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}

	ws := zapcore.AddSync(ll)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = cfg.TimeKey
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.LevelKey = "level"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.CallerKey = "caller"
	encoderConfig.MessageKey = "msg"
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		loc, err := time.LoadLocation(cfg.TimeZone)
		if err != nil {
			loc = time.Local
		}

		t = t.In(loc)

		type appendTimeEncoder interface {
			// AppendTimeLayout description of the Go function.
			//
			// This function takes a time.Time and a string as parameters.
			AppendTimeLayout(time.Time, string)
		}

		if enc, ok := enc.(appendTimeEncoder); ok {
			enc.AppendTimeLayout(t, cfg.TimeFormat)
			return
		}

		enc.AppendString(t.Format(cfg.TimeFormat))
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, ws, cfg.LogLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), cfg.LogLevel),
	)

	zapLogger := zap.New(core, zap.AddCaller())
	logger := zapLogger.Named(cfg.Name).Sugar()

	if cfg.WithOptions != nil {
		logger = logger.WithOptions(cfg.WithOptions...)
	}

	gLoggers[cfg.Name] = logger

	logger.Debug("logger initialized")

	return logger
}
