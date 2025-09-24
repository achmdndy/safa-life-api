package routes

import (
	"github.com/achmdndy/safa-life-api/src/presentation/handlers/translation"
	"github.com/gin-gonic/gin"
)

// TranslationRoutes sets up all routes for the Translation module.
func TranslationRoutes(router *gin.RouterGroup, handler *translation.Handler) {
	group := router.Group("/translations")
	{
		editions := group.Group("/editions")
		{
			editions.GET("", handler.GetAllTranslations)
			editions.POST("", handler.CreateTranslation)
			editions.GET("/:id", handler.GetTranslationByID)
			editions.PUT("/:id", handler.UpdateTranslation)
			editions.DELETE("/:id", handler.DeleteTranslation)
		}

		ayahs := group.Group("/ayahs")
		{
			ayahs.POST("", handler.CreateAyahTranslation)
			ayahs.GET("/:translation_id/:surah_id/:ayah_id", handler.GetAyahTranslation)
			ayahs.PUT("/:translation_id/:surah_id/:ayah_id", handler.UpdateAyahTranslation)
			ayahs.DELETE("/:translation_id/:surah_id/:ayah_id", handler.DeleteAyahTranslation)
			ayahs.GET("/by-ayah/:surah_id/:ayah_id", handler.GetTranslationsForAyah)
			ayahs.GET("/by-surah/:translation_id/:surah_id", handler.GetTranslationsForSurah)
		}
	}
}