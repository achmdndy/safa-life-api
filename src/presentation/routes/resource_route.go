package routes

import (
	"github.com/achmdndy/safa-life-api/src/presentation/handlers/resource"
	"github.com/gin-gonic/gin"
)

// ResourceRoutes sets up all routes for the Resource module.
func ResourceRoutes(router *gin.RouterGroup, handler *resource.Handler) {
	group := router.Group("/resources")
	{
		group.GET("", handler.GetAllResources)
		group.POST("", handler.CreateResource)
		group.GET("/:id", handler.GetResourceByID)
		group.PUT("/:id", handler.UpdateResource)
		group.DELETE("/:id", handler.DeleteResource)
		group.PATCH("/:id/status", handler.UpdateInstallStatus)
	}
}
