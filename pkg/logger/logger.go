package logger

import (
	"os"

	"github.com/dchebakov/tracker/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getLoggerLevel(cfg *config.Config) zapcore.Level {
	loggerLevelMap := map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}

	level, exists := loggerLevelMap[cfg.Api.LogLevel]
	if !exists {
		return zapcore.DebugLevel
	}

	return level
}

func NewLogger(cfg *config.Config) *zap.SugaredLogger {
	logWriter := zapcore.AddSync(os.Stderr)
	logLevel := getLoggerLevel(cfg)
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	if cfg.Api.Mode != "dev" {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	encoder := zapcore.NewJSONEncoder(encoderCfg)
	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	sugarLogger := logger.Sugar()
	defer sugarLogger.Sync()

	return sugarLogger
}
