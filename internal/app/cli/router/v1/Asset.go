package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ka1i/cli/internal/app/cli/controller"
)

func Asset(engine *gin.RouterGroup) {
	apiRoutes := engine.Group(prefix + "/asset")
	apiRoutes.GET("/*filepath", controller.EmbedAsset)
}
