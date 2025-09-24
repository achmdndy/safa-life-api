package routes

import (
	"github.com/achmdndy/safa-life-api/src/presentation/handlers/story"
	"github.com/gin-gonic/gin"
)

// StoryRoutes sets up all routes for the Story module.
func StoryRoutes(router *gin.RouterGroup, handler *story.Handler) {
	group := router.Group("/stories")
	{
		group.GET("", handler.GetAllStories)
		group.POST("", handler.CreateStory)
		group.GET("/:id", handler.GetStoryByID)
		group.PUT("/:id", handler.UpdateStory)
		group.DELETE("/:id", handler.DeleteStory)

		ayahs := group.Group("/:id/ayahs")
		{
			ayahs.GET("", handler.GetAyahsForStory)
			ayahs.POST("", handler.AddAyahToStory)
			ayahs.DELETE("/:surah_id/:ayah_id", handler.RemoveAyahFromStory)
		}
	}
}
