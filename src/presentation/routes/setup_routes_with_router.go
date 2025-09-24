package routes

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

// SetupRoutesWithRouter sets up routes on an existing router instance
func SetupRoutesWithRouter(router *gin.Engine, config RouteConfig) {
	v1 := router.Group("/api/v1")

	SetupHealthRoutes(v1, config.HealthHandler)
	QuranRoutes(v1, config.QuranHandler)
	ReciterRoutes(v1, config.ReciterHandler)
	TajweedRoutes(v1, config.TajweedHandler)
	TranslationRoutes(v1, config.TranslationHandler)
	AudioRoutes(v1, config.AudioHandler)
	ResourceRoutes(v1, config.ResourceHandler)
	StoryRoutes(v1, config.StoryHandler)
	TafsirRoutes(v1, config.TafsirHandler)
	TopicRoutes(v1, config.TopicHandler)

	router.GET("/", func(c *gin.Context) {
		start := time.Now()
		core.Success(c, 200, "🌙 Safa Life API - Islamic Lifestyle Backend", gin.H{
			"version": "v1.0.0",
			"status":  "running",
		}, start)
	})
}