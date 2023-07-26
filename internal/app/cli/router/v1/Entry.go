package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ka1i/cli/internal/app/cli/controller"
)

func Entry(engine *gin.RouterGroup) {
	apiRoutes := engine.Group(prefix + "/entry")
	apiRoutes.GET("/query", controller.EntryQuery)
}
