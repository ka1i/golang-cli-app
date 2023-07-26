package handlers

import (
	zaplog "github.com/ka1i/cli/pkg/logger"
	"gorm.io/gorm/logger"
)

func GormLogger() logger.Interface {
	logger := zaplog.NewGormLogger()
	logger.IgnoreRecordNotFoundError = true
	return logger
}
