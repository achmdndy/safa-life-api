package routes

import (
	"github.com/achmdndy/safa-life-api/src/presentation/handlers/topic"
	"github.com/gin-gonic/gin"
)

// TopicRoutes sets up all routes for the Topic module.
func TopicRoutes(router *gin.RouterGroup, handler *topic.Handler) {
	group := router.Group("/topics")
	{
		group.GET("", handler.GetAllTopics)
		group.POST("", handler.CreateTopic)
		group.GET("/:id", handler.GetTopicByID)
		group.PUT("/:id", handler.UpdateTopic)
		group.DELETE("/:id", handler.DeleteTopic)

		ayahs := group.Group("/:id/ayahs")
		{
			ayahs.GET("", handler.GetAyahsForTopic)
			ayahs.POST("", handler.AddAyahToTopic)
			ayahs.DELETE("/:surah_id/:ayah_id", handler.RemoveAyahFromTopic)
		}
	}
}
