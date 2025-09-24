package routes

import (
	"github.com/gin-gonic/gin"

	healthHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/health"
)

func SetupHealthRoutes(router *gin.RouterGroup, handler *healthHandler.Handler) {
	health := router.Group("/health")
	{
		health.GET("", handler.GetHealth)
		health.GET("/simple", handler.GetHealthSimple)
		health.GET("/db", handler.GetHealthDB)
		health.GET("/redis", handler.GetHealthRedis)
	}
}