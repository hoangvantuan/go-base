package logger

import (
	"os"

	"github.com/hoangvantuan/go-base/infra"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

// logrotation config
var logPath = "logs/gobase.log"
var maxSize = 10
var maxBackups = 5
var maxAge = 30
var compress = false

// SetupZapLogger will setting zap logger
func initZapLogger() {
	if logger != nil && sugar != nil {
		return
	}

	core := zapcore.NewCore(getEncoder(), getLogWriter(), getLogLevel())

	logger = zap.New(core, zap.AddCaller())
	sugar = logger.Sugar()
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

func getEncoder() zapcore.Encoder {
	var encoderCfg zapcore.EncoderConfig

	if infra.IsProduction() {
		encoderCfg = zap.NewProductionEncoderConfig()
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	} else {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	return zapcore.NewConsoleEncoder(encoderCfg)

}

func getLogWriter() zapcore.WriteSyncer {
	if infra.IsProduction() {
		lumberJackLogger := &lumberjack.Logger{
			Filename:   logPath,
			MaxSize:    maxSize,
			MaxBackups: maxBackups,
			MaxAge:     maxAge,
			Compress:   compress,
		}

		return zapcore.AddSync(lumberJackLogger)
	}

	return zapcore.AddSync(os.Stdout)
}

func getLogLevel() zapcore.Level {
	if infra.IsProduction() {
		return zapcore.ErrorLevel
	}

	return zapcore.DebugLevel
}
