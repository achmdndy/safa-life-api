package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/presentation/handlers/quran"
)

// QuranRoutes sets up Quran-related routes
func QuranRoutes(router *gin.RouterGroup, quranHandler *quran.Handler) {
	quranGroup := router.Group("/quran")
	{
		// Surah routes
		surahGroup := quranGroup.Group("/surahs")
		{
			surahGroup.GET("", quranHandler.GetAllSurahsHandler().Handle)
			surahGroup.GET("/:id", quranHandler.GetSurahByIdHandler().Handle)
			surahGroup.GET("/number/:number", quranHandler.GetSurahByNumberHandler().Handle)
			surahGroup.POST("", quranHandler.CreateSurahHandler().Handle)
			surahGroup.PUT("/:id", quranHandler.UpdateSurahHandler().Handle)
			surahGroup.DELETE("/:id", quranHandler.DeleteSurahHandler().Handle)
		}

		// Ayah routes
		ayahGroup := quranGroup.Group("/ayahs")
		{
			ayahGroup.GET("/:id", quranHandler.GetAyahByIdHandler().Handle)
			ayahGroup.GET("/surah/:surahId", quranHandler.GetAyahsBySurahHandler().Handle)
			ayahGroup.GET("/juz/:juzNumber", quranHandler.GetAyahsByJuzHandler().Handle)
			ayahGroup.POST("", quranHandler.CreateAyahHandler().Handle)
			ayahGroup.PUT("/:id", quranHandler.UpdateAyahHandler().Handle)
			ayahGroup.DELETE("/:id", quranHandler.DeleteAyahHandler().Handle)
		}

		// Juz routes
		juzGroup := quranGroup.Group("/juz")
		{
			juzGroup.GET("", quranHandler.GetAllJuzHandler().Handle)
			juzGroup.GET("/:id", quranHandler.GetJuzByIdHandler().Handle)
			juzGroup.GET("/number/:number", quranHandler.GetJuzByNumberHandler().Handle)
			juzGroup.POST("", quranHandler.CreateJuzHandler().Handle)
			juzGroup.PUT("/:id", quranHandler.UpdateJuzHandler().Handle)
			juzGroup.DELETE("/:id", quranHandler.DeleteJuzHandler().Handle)
		}

		// Translation routes
		translationGroup := quranGroup.Group("/translations")
		{
			editionGroup := translationGroup.Group("/editions")
			{
				editionGroup.GET("", quranHandler.GetTranslationEditionsHandler().Handle)
			}

			ayahTranslationGroup := translationGroup.Group("/ayahs")
			{
				ayahTranslationGroup.GET("/surah/:surahId/edition/:editionId", quranHandler.GetAyahTranslationsBySurahAndEditionHandler().Handle)
				ayahTranslationGroup.GET("/ayah/:ayahId/edition/:editionId", quranHandler.GetAyahTranslationByAyahAndEditionHandler().Handle)
			}
		}

		// Reciter routes
		reciterGroup := quranGroup.Group("/reciters")
		{
			reciterGroup.GET("", quranHandler.GetAllRecitersHandler().Handle)
			reciterGroup.GET("/:id", quranHandler.GetReciterByIdHandler().Handle)
			reciterGroup.GET("/name/:name", quranHandler.GetReciterByNameHandler().Handle)
			reciterGroup.POST("", quranHandler.CreateReciterHandler().Handle)
			reciterGroup.PUT("/:id", quranHandler.UpdateReciterHandler().Handle)
			reciterGroup.DELETE("/:id", quranHandler.DeleteReciterHandler().Handle)
		}

		// Ayah Audio routes
		audioGroup := quranGroup.Group("/audio")
		{
			ayahAudioGroup := audioGroup.Group("/ayahs")
			{
				ayahAudioGroup.GET("/:id", quranHandler.GetAyahAudioFileByIdHandler().Handle)
				ayahAudioGroup.GET("/ayah/:ayahId/reciter/:reciterId", quranHandler.GetAyahAudioFileByAyahAndReciterHandler().Handle)
				ayahAudioGroup.GET("/surah/:surahId/reciter/:reciterId", quranHandler.GetAyahAudioFilesBySurahAndReciterHandler().Handle)
				ayahAudioGroup.POST("", quranHandler.CreateAyahAudioFileHandler().Handle)
				ayahAudioGroup.PUT("/:id", quranHandler.UpdateAyahAudioFileHandler().Handle)
				ayahAudioGroup.DELETE("/:id", quranHandler.DeleteAyahAudioFileHandler().Handle)
			}
		}
	}
}
