/*
Copyright © 2023 Mardan https://www.mardan.wiki
*/
package app

import (
	"github.com/ka1i/cli/internal/app/cli"
	"github.com/ka1i/cli/pkg/info"
)

func init() {
	info.ShowBanner()
}

type service string

const (
	Cli service = "cli"
)

func (service service) Usage() string {
	switch service {
	case Cli:
		return cli.Usage()
	}
	return ""
}

func (service service) Flags() func() (bool, error) {
	switch service {
	case Cli:
		return cli.Flags
	}
	return func() (bool, error) {
		return true, nil
	}
}

func (service service) Service() func() error {
	switch service {
	case Cli:
		return cli.Execute
	}
	return nil
}
