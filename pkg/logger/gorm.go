package logger

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormorigin "gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var gorm *zap.Logger

func GormLog() {
	core := zapcore.NewTee(
		*Console(),
		*LocalFile("sql.log"),
	)
	gorm = zap.New(core, zap.AddCaller(), zap.Development())
}

type GormLogger struct {
	ZapLogger                 *zap.Logger
	LogLevel                  gormlogger.LogLevel
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
}

func NewGormLogger() *GormLogger {
	return &GormLogger{
		ZapLogger:                 gorm,
		LogLevel:                  gormlogger.Info,
		SlowThreshold:             1 * time.Second,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: false,
	}
}

func (gl *GormLogger) SetAsDefault() {
	gormlogger.Default = gl
}

func (gl GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return GormLogger{
		ZapLogger:                 gl.ZapLogger,
		SlowThreshold:             gl.SlowThreshold,
		LogLevel:                  level,
		SkipCallerLookup:          gl.SkipCallerLookup,
		IgnoreRecordNotFoundError: gl.IgnoreRecordNotFoundError,
	}
}

func (gl GormLogger) Print(v ...interface{}) {
	if gl.LogLevel < gormlogger.Error {
		return
	}
	gl.logger(context.Background()).Sugar().Errorf(fmt.Sprint(v...))
}

func (gl GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	if gl.LogLevel < gormlogger.Info {
		return
	}
	gl.logger(ctx).Sugar().Debugf(str, args...)
}

func (gl GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	if gl.LogLevel < gormlogger.Warn {
		return
	}
	gl.logger(ctx).Sugar().Warnf(str, args...)
}

func (gl GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	if gl.LogLevel < gormlogger.Error {
		return
	}
	gl.logger(ctx).Sugar().Errorf(str, args...)
}

func (gl GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if gl.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	logger := gl.logger(ctx)
	switch {
	case err != nil && gl.LogLevel >= gormlogger.Error && (!gl.IgnoreRecordNotFoundError || !errors.Is(err, gormorigin.ErrRecordNotFound)):
		sql, rows := fc()
		logger.Warn("trace", zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	case gl.SlowThreshold != 0 && elapsed > gl.SlowThreshold && gl.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		logger.Warn("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	case gl.LogLevel >= gormlogger.Info:
		sql, rows := fc()
		logger.Debug("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	}
}

var (
	gormPackage    = filepath.Join("gorm.io", "gorm")
	zapgormPackage = filepath.Join("moul.io", "zapgorm2")
)

func (gl GormLogger) logger(ctx context.Context) *zap.Logger {
	logger := gl.ZapLogger

	if gl.SkipCallerLookup {
		return logger
	}

	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, zapgormPackage):
		default:
			return logger.WithOptions(zap.AddCallerSkip(i - 1))
		}
	}
	return logger
}
