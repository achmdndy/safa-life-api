package routes

import (
	"github.com/achmdndy/safa-life-api/src/presentation/handlers/tafsir"
	"github.com/gin-gonic/gin"
)

// TafsirRoutes sets up all routes for the Tafsir module.
func TafsirRoutes(router *gin.RouterGroup, handler *tafsir.Handler) {
	group := router.Group("/tafsirs")
	{
		editions := group.Group("/editions")
		{
			editions.GET("", handler.GetAllTafsirs)
			editions.POST("", handler.CreateTafsir)
			editions.GET("/:id", handler.GetTafsirByID)
			editions.PUT("/:id", handler.UpdateTafsir)
			editions.DELETE("/:id", handler.DeleteTafsir)
		}

		ayahs := group.Group("/ayahs")
		{
			ayahs.POST("", handler.CreateAyahTafsir)
			ayahs.GET("/:tafsir_id/:surah_id/:ayah_id", handler.GetAyahTafsir)
			ayahs.PUT("/:tafsir_id/:surah_id/:ayah_id", handler.UpdateAyahTafsir)
			ayahs.DELETE("/:tafsir_id/:surah_id/:ayah_id", handler.DeleteAyahTafsir)
			ayahs.GET("/by-ayah/:surah_id/:ayah_id", handler.GetTafsirsForAyah)
			ayahs.GET("/by-surah/:tafsir_id/:surah_id", handler.GetTafsirsForSurah)
		}
	}
}
