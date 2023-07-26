package prepare

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ka1i/cli/internal/pkg/system/service"
	"github.com/ka1i/cli/pkg/info"
	"github.com/ka1i/cli/pkg/logger"
	"github.com/ka1i/cli/pkg/timezone"
)

func Environment() {
	app()
	appMode()
	appTimeZone()
}

func app() {
	logger.Printf("micro service: %s\n", info.MicroService)
}

func appMode() {
	var mode string
	mode = os.Getenv("AOK_APP_MODE")
	if len(mode) == 0 {
		mode = gin.TestMode
	}

	if (mode != gin.DebugMode) && (mode != gin.TestMode) && (mode != gin.ReleaseMode) {
		logger.Printf("Invalid Mode: %s\n", mode)
		mode = gin.TestMode
	}

	service.MODE = mode
	gin.SetMode(service.MODE)
	logger.Printf("System Running Mode: %s\n", service.MODE)
}

func appTimeZone() {
	tz := timezone.TZ()
	os.Setenv("TZ", tz)
	logger.Printf("System Running TimeZone: %s\n", tz)
}
