/*
Copyright © 2023 Mardan https://www.mardan.wiki
*/
package main

import (
	"os"

	"github.com/ka1i/cli/internal/app"
)

func main() {
	// bootstrap
	service := app.App
	service.Usage = app.Cli.Usage()
	service.Flags = app.Cli.Flags()
	service.Service = app.Cli.Service()
	code, err := service.Run()
	defer os.Exit(code)
	if err != nil {
		panic(err)
	}
}
