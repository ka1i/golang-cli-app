/*
Copyright © 2023 Mardan https://www.mardan.wiki
*/
package app

import (
	"errors"
	"os"

	"github.com/ka1i/cli/internal/pkg/config"
	"github.com/ka1i/cli/internal/pkg/usage"
	"github.com/ka1i/cli/pkg/info"
	"github.com/ka1i/cli/pkg/logger"
)

type app struct {
	Usage   string
	Flags   func() (bool, error)
	Service func() error

	failed  int
	success int
}

func GetApp() *app {
	return &app{
		failed:  137,
		success: 0,
	}
}

func (app *app) Run() (int, error) {
	var err error
	var isBreak bool = false
	var isCapture bool = false

	// check service instance
	if app.Service == nil {
		return app.failed, errors.New("exception: Service Instance NOT DEFINE")
	}

	// process app flags
	if len(os.Args) > 1 {
		var argv = os.Args[1:]
		var argc = len(os.Args) - 1
		isBreak, isCapture, err = app.flags(argc, argv)
		if err != nil {
			return app.success, err
		}
	}

	// service entry
	if isBreak {
		return app.success, err
	}

	// config parser
	config.Configure.Init()
	err = config.Configure.Error
	if err != nil {
		panic(err)
	}

	// init logger
	logger.Init()

	if isCapture {
		isBreak, err = app.Flags()
		if err != nil {
			return app.failed, err
		}
		if isBreak {
			return app.success, err
		}
	}

	// run service
	err = app.Service()
	return app.success, err
}

func (app *app) flags(argc int, argv []string) (bool, bool, error) {
	var err error = nil
	var isBreak bool = false
	var isCapture bool = false

	switch argv[0] {
	case "-h", "--help", "help":
		isBreak = true
		usage.Usage(app.Usage)
	case "-v", "--version", "version":
		isBreak = true
		info.Version.Print()
	default:
		isCapture = true
	}
	return isBreak, isCapture, err
}

var App = GetApp()
