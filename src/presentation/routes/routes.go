package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/safalife/core-api/src/presentation/container"
	"github.com/safalife/core-api/src/presentation/core"
	"github.com/safalife/core-api/src/presentation/middlewares"
)

// SetupRoutes configures all application routes
func SetupRoutes(router *gin.Engine, presentationContainer *container.PresentationContainer) {
	// CORS middleware - should be first to handle preflight requests
	router.Use(middlewares.CORS())

	// Attach monitoring middleware
	router.Use(middlewares.TracingMiddleware(presentationContainer.GetMonitoringService()))
	router.Use(middlewares.MetricsMiddleware(presentationContainer.GetMonitoringService()))
	router.Use(middlewares.MonitoringMiddleware(presentationContainer.GetMonitoringService()))

	// Setup Prometheus and health endpoints
	router.GET("/metrics", gin.WrapH(presentationContainer.GetMonitoringHandlers().GetPrometheusHandler()))
	router.GET("/health", gin.WrapH(presentationContainer.GetMonitoringHandlers().GetHealthHandler()))

	// Root welcome route
	router.GET("/", func(c *gin.Context) {
		start := time.Now()

		welcomeData := map[string]interface{}{
			"application": "Safa Life API",
			"version":     "1.0.0",
			"description": "Welcome to Safa Life API",
			"endpoints": map[string]string{
				"health":  "/health",
				"metrics": "/metrics",
				"api":     "/v1",
				"docs":    "/swagger/index.html",
			},
		}

		core.Success(c, http.StatusOK, "Welcome to Safa Life API", welcomeData, start)
	})

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API version 1
	v1 := router.Group("/v1")
	{
		// Quran routes
		QuranRoutes(v1, presentationContainer.QuranHandler, presentationContainer.GetAuthMiddleware())
	}
}
