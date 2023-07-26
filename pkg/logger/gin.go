package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var gin *zap.Logger

func GinLog() {
	core := zapcore.NewTee(
		*Console(),
		*LocalFile("access.log"),
	)
	gin = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Development())
}

func GinDebug(message string, fields ...zap.Field) {
	gin.Debug(message, fields...)
}

func GinInfo(message string, fields ...zap.Field) {
	gin.Info(message, fields...)
}

func GinWarn(message string, fields ...zap.Field) {
	gin.Warn(message, fields...)
}

func GinError(message string, fields ...zap.Field) {
	gin.Error(message, fields...)
}
