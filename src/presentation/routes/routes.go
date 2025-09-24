package routes

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/achmdndy/safa-life-api/src/presentation/core"
	audioHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/audio"
	healthHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/health"
	quranHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/quran"
	reciterHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/reciter"
	resourceHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/resource"
	storyHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/story"
	tafsirHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/tafsir"
	tajweedHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/tajweed"
	topicHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/topic"
	translationHandler "github.com/achmdndy/safa-life-api/src/presentation/handlers/translation"
	"github.com/achmdndy/safa-life-api/src/presentation/middlewares"
)

type RouteConfig struct {
	HealthHandler      *healthHandler.Handler
	QuranHandler       *quranHandler.Handler
	ReciterHandler     *reciterHandler.Handler
	TajweedHandler     *tajweedHandler.Handler
	TranslationHandler *translationHandler.Handler
	AudioHandler       *audioHandler.Handler
	ResourceHandler    *resourceHandler.Handler
	StoryHandler       *storyHandler.Handler
	TafsirHandler      *tafsirHandler.Handler
	TopicHandler       *topicHandler.Handler
}

func SetupRoutes(config RouteConfig) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(middlewares.LoggingMiddleware())
	router.Use(middlewares.RecoveryMiddleware())
	router.Use(middlewares.CORSMiddleware())

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

	return router
}