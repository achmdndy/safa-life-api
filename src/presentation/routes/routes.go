package routes

import (
	"time"

	"github.com/gin-gonic/gin"

	healthHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/health"
	"github.com/achmdndy/safa-life-api/src/presentation/middlewares"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

type RouteConfig struct {
	HealthHandler *healthHandler.Handler
}

func SetupRoutes(config RouteConfig) *gin.Engine {
	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// Middleware
	router.Use(middlewares.LoggingMiddleware())
	router.Use(middlewares.RecoveryMiddleware())
	router.Use(middlewares.CORSMiddleware())

	// API versioning
	v1 := router.Group("/api/v1")

	// Setup route groups
	SetupHealthRoutes(v1, config.HealthHandler)

	// Root endpoint
	router.GET("/", func(c *gin.Context) {
		start := time.Now()
		core.Success(c, 200, "🌙 Safa Life API - Islamic Lifestyle Backend", gin.H{
			"version": "v1.0.0",
			"status":  "running",
		}, start)
	})

	return router
}

// SetupRoutesWithRouter sets up routes on an existing router
func SetupRoutesWithRouter(router *gin.Engine, config RouteConfig) {
	// Middleware
	router.Use(middlewares.LoggingMiddleware())
	router.Use(middlewares.RecoveryMiddleware())
	router.Use(middlewares.CORSMiddleware())

	// API versioning
	v1 := router.Group("/api/v1")

	// Setup route groups
	SetupHealthRoutes(v1, config.HealthHandler)

	// Root endpoint
	router.GET("/", func(c *gin.Context) {
		start := time.Now()
		core.Success(c, 200, "🌙 Safa Life API - Islamic Lifestyle Backend", gin.H{
			"version": "v1.0.0",
			"status":  "running",
		}, start)
	})
}