package utils

import (
	"os"
	"strings"

	"github.com/ka1i/cli/internal/pkg/config"
)

func Resolver() string {
	options := config.Cfg.Get().App
	var addr []string

	if options.Addr != "" {
		addr = append(addr, options.Addr)
	} else {
		addr = append(addr, "localhost")
	}

	if options.Port != "" {
		addr = append(addr, options.Port)
	} else {
		addr = append(addr, "80")
	}

	port := os.Getenv("GCA_APP_PORT")
	if len(port) != 0 {
		addr[1] = port
	}

	return strings.Join(addr, ":")
}
