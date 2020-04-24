package logger

import (
	"github.com/hoangvantuan/go-base/infra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

// SetupZapLogger will setting zap logger
func initZapLogger() {
	if logger != nil {
		return
	}

	if infra.IsProduction() {
		logger, _ = zap.NewProduction()
	}

	logger, _ = zap.NewDevelopment()
}

// Info level
func Info(msg string, fields ...zapcore.Field) {
	initZapLogger()
	logger.Info(msg, fields...)
}

// Warn level
func Warn(msg string, fields ...zapcore.Field) {
	initZapLogger()
	logger.Warn(msg, fields...)
}

// Error levele
func Error(msg string, fields ...zapcore.Field) {
	initZapLogger()
	logger.Error(msg, fields...)
}

// Fatal level
func Fatal(msg string, fields ...zapcore.Field) {
	initZapLogger()
	logger.Fatal(msg, fields...)
}

// Debug level
func Debug(msg string, fields ...zapcore.Field) {
	initZapLogger()
	logger.Debug(msg, fields...)
}
