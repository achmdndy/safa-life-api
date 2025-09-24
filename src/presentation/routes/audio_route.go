package routes

import (
	"github.com/achmdndy/safa-life-api/src/presentation/handlers/audio"
	"github.com/gin-gonic/gin"
)

// AudioRoutes sets up all routes for the Audio module.
func AudioRoutes(router *gin.RouterGroup, handler *audio.Handler) {
	group := router.Group("/audio")
	{
		ayahs := group.Group("/ayahs")
		{
			ayahs.POST("", handler.CreateAyahAudio)
			ayahs.GET("/:reciter_id/:surah_id/:ayah_id", handler.GetAyahAudio)
			ayahs.PUT("/:reciter_id/:surah_id/:ayah_id", handler.UpdateAyahAudio)
			ayahs.DELETE("/:reciter_id/:surah_id/:ayah_id", handler.DeleteAyahAudio)
		}

		surahs := group.Group("/surahs")
		{
			surahs.GET("/:reciter_id/:surah_id", handler.GetAudioForSurah)
		}
	}
}
