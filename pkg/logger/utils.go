package logger

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ka1i/cli/internal/pkg/config"
	"github.com/ka1i/cli/pkg/info"
	"github.com/ka1i/cli/pkg/utils"
	"go.uber.org/zap/zapcore"
)

var EncoderConfig = zapcore.EncoderConfig{
	TimeKey:        "time",
	LevelKey:       "level",
	NameKey:        "logger",
	CallerKey:      "caller",
	FunctionKey:    zapcore.OmitKey,
	MessageKey:     "msg",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.CapitalColorLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
	EncodeName:     zapcore.FullNameEncoder,
}

func Console() *zapcore.Core {
	// 控制台日志句柄
	encoderConfig := EncoderConfig
	consoleCore := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	return &consoleCore
}

func LocalFile(filename string) *zapcore.Core {
	dir := config.Cfg.Get().Logdir

	logdir := path.Join(dir, info.MicroService)

	// 持久化系统日志
	if !utils.IsExist(logdir) {
		err := os.MkdirAll(logdir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	file_info, _ := os.OpenFile(
		filepath.Join(logdir, filename),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644,
	)

	fileWriteSyncer := zapcore.AddSync(file_info)

	// 文件日志句柄
	encoderConfig := EncoderConfig
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	fileCore := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), fileWriteSyncer, zapcore.DebugLevel)
	return &fileCore
}

func validPrintf(format string, v ...any) string {
	if len(format) != 0 || format[len(format)-1] == '\n' {
		format = strings.Replace(format, "\n", "", -1)
	}
	msg := fmt.Sprintf(format, v...)
	return msg
}
