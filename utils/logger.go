package utils

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger returns the zap logger
func Logger() *zap.Logger {
	env := os.Getenv("ENV")
	config := zap.NewDevelopmentConfig()
	config.DisableStacktrace = true
	if env == "prod" || env == "staging" {
		config = zap.NewProductionConfig()
	}
	if env == "test" {
		config.Level = zap.NewAtomicLevelAt(zapcore.Level(zapcore.FatalLevel))
	}

	logger, err := config.Build()
	if err != nil {
		fmt.Println("Failed to create logger")
		os.Exit(1)
	}
	return logger
}

// GetLogger return the zap sugared logger
func GetLogger() *zap.SugaredLogger {
	return Logger().Sugar()
}
