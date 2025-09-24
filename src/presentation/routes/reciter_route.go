package routes

import (
	"github.com/achmdndy/safa-life-api/src/presentation/handlers/reciter"
	"github.com/gin-gonic/gin"
)

// ReciterRoutes sets up all routes for the Reciter module.
func ReciterRoutes(router *gin.RouterGroup, handler *reciter.Handler) {
	group := router.Group("/reciters")
	{
		group.GET("", handler.GetAllReciters)
		group.POST("", handler.CreateReciter)
		group.GET("/:id", handler.GetReciterByID)
		group.PUT("/:id", handler.UpdateReciter)
		group.DELETE("/:id", handler.DeleteReciter)
	}
}