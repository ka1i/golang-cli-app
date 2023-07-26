package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ka1i/cli/internal/app/cli/router/v1"
	"github.com/ka1i/cli/internal/pkg/controller"
)

// define root routes
func rootRoutes(engine *gin.Engine, context string) {
	apiRoutes := engine.Group("/")
	apiRoutes.GET("health", controller.HealthChecks)
}

// registry all routes
func SetRouter(engine *gin.Engine, context string) {
	rootRoutes(engine, context)

	// service context
	subApiRoutes := engine.Group(context)

	// api ver v1
	v1.Entry(subApiRoutes)
}
