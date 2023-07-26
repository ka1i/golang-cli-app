package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ka1i/cli/internal/pkg/system/service"
	"github.com/ka1i/cli/pkg/info"
	"github.com/ka1i/cli/pkg/timezone"
)

func HealthChecks(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"service":      info.MicroServiceName,
			"timestamp":    timezone.Format(time.Now()),
			"dependencies": service.Dependencies.Get(),
		},
	)
}
