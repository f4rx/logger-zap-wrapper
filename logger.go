package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewLogger() *zap.Logger {
	encoding := "console"
	encodeLevel := zapcore.CapitalColorLevelEncoder
	logType, exists := os.LookupEnv("LOG_TYPE")
	if exists && logType == "json" {
		encoding = "json"
		encodeLevel = zapcore.CapitalLevelEncoder
	}

	level := zapcore.InfoLevel
	_, exists = os.LookupEnv("DEBUG")
	if exists {
		level = zapcore.DebugLevel
	}

	cfg := zap.Config{
		Encoding:         encoding,
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    encodeLevel,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(fmt.Sprintf("Error during setting logger: %v", err))
	}
	logger.Info("Logger enabled: ", zap.String("level", level.String()), zap.String("format", encoding))

	return logger
}

func NewSugaredLogger() *zap.SugaredLogger {
	logger := NewLogger()
	return logger.Sugar()
}
