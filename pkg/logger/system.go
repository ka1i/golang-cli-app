package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var system *zap.Logger

func SystemLog() {
	core := zapcore.NewTee(
		*Console(),
		*LocalFile("system.log"),
	)
	system = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Development())
}

func Debug(message string, fields ...zap.Field) {
	system.Debug(message, fields...)
}

func Debugf(format string, v ...any) {
	msg := validPrintf(format, v...)
	system.Debug(msg)
}

func Info(message string, fields ...zap.Field) {
	system.Info(message, fields...)
}

func Infof(format string, v ...any) {
	msg := validPrintf(format, v...)
	system.Info(msg)
}

func Warn(message string, fields ...zap.Field) {
	system.Warn(message, fields...)
}

func Warnf(format string, v ...any) {
	msg := validPrintf(format, v...)
	system.Warn(msg)
}

func Error(message string, fields ...zap.Field) {
	system.Error(message, fields...)
}

func Errorf(format string, v ...any) {
	msg := validPrintf(format, v...)
	system.Error(msg)
}

func Printf(format string, v ...any) {
	msg := validPrintf(format, v...)
	system.Info(msg)
}

func Println(v ...string) {
	system.Info(strings.Join(v, " "))
}
