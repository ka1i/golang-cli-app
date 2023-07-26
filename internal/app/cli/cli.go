package cli

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ka1i/cli/internal/app/cli/router"
	"github.com/ka1i/cli/internal/pkg/datasource"
	"github.com/ka1i/cli/internal/pkg/handlers"
	"github.com/ka1i/cli/internal/pkg/system/prepare"
	"github.com/ka1i/cli/pkg/logger"
	"github.com/ka1i/cli/pkg/utils"
)

func Usage() string {
	var buffer strings.Builder

	return buffer.String()
}

func Flags() (bool, error) {
	var err error
	var argv = os.Args[1:]
	var isBreak bool = false

	switch argv[0] {
	default:
		err = errors.New("please check usage")
	}
	return isBreak, err
}

func Execute() error {
	// service prepare
	prepare.Environment()

	// init data source
	datasource.Datasource.Init()

	// system initialize
	prepare.Configure()

	// configure web server
	engine := gin.New()
	handlers.SetHandlers(engine)
	router.SetRouter(engine, "/cli-server")

	var addr string = utils.Resolver()

	server := &http.Server{
		Addr:           addr,
		Handler:        engine,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: 1 << 20,
	}

	// start web server
	logger.Printf("Server Listening %s Success...\n", addr)
	err := server.ListenAndServe()
	return err
}
