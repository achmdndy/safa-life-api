package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/presentation/handlers/quran"
)

// QuranRoutes sets up Quran-related routes
func QuranRoutes(router *gin.RouterGroup, quranHandler *quran.Handler, authMW gin.HandlerFunc) {
	quranGroup := router.Group("/quran")
	{
		// Surah routes
		surahGroup := quranGroup.Group("/surahs")
		{
			surahGroup.GET("", quranHandler.GetAllSurahsHandler().Handle)
			surahGroup.GET("/:id", quranHandler.GetSurahByIdHandler().Handle)
			surahGroup.GET("/number/:number", quranHandler.GetSurahByNumberHandler().Handle)

			surahProtected := surahGroup.Group("")
			surahProtected.Use(authMW)
			surahProtected.POST("", quranHandler.CreateSurahHandler().Handle)
			surahProtected.PUT("/:id", quranHandler.UpdateSurahHandler().Handle)
			surahProtected.DELETE("/:id", quranHandler.DeleteSurahHandler().Handle)
		}

		// Ayah routes
		ayahGroup := quranGroup.Group("/ayahs")
		{
			ayahGroup.GET("/:id", quranHandler.GetAyahByIdHandler().Handle)
			ayahGroup.GET("/surah/:surahId", quranHandler.GetAyahsBySurahHandler().Handle)
			ayahGroup.GET("/juz/:juzNumber", quranHandler.GetAyahsByJuzHandler().Handle)

			ayahProtected := ayahGroup.Group("")
			ayahProtected.Use(authMW)
			ayahProtected.POST("", quranHandler.CreateAyahHandler().Handle)
			ayahProtected.PUT("/:id", quranHandler.UpdateAyahHandler().Handle)
			ayahProtected.DELETE("/:id", quranHandler.DeleteAyahHandler().Handle)
		}

		// Juz routes
		juzGroup := quranGroup.Group("/juz")
		{
			juzGroup.GET("", quranHandler.GetAllJuzHandler().Handle)
			juzGroup.GET("/:id", quranHandler.GetJuzByIdHandler().Handle)
			juzGroup.GET("/number/:number", quranHandler.GetJuzByNumberHandler().Handle)

			juzProtected := juzGroup.Group("")
			juzProtected.Use(authMW)
			juzProtected.POST("", quranHandler.CreateJuzHandler().Handle)
			juzProtected.PUT("/:id", quranHandler.UpdateJuzHandler().Handle)
			juzProtected.DELETE("/:id", quranHandler.DeleteJuzHandler().Handle)
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

			reciterProtected := reciterGroup.Group("")
			reciterProtected.Use(authMW)
			reciterProtected.POST("", quranHandler.CreateReciterHandler().Handle)
			reciterProtected.PUT("/:id", quranHandler.UpdateReciterHandler().Handle)
			reciterProtected.DELETE("/:id", quranHandler.DeleteReciterHandler().Handle)
		}

		// Ayah Audio routes
		audioGroup := quranGroup.Group("/audio")
		{
			ayahAudioGroup := audioGroup.Group("/ayahs")
			{
				ayahAudioGroup.GET("/:id", quranHandler.GetAyahAudioFileByIdHandler().Handle)
				ayahAudioGroup.GET("/ayah/:ayahId/reciter/:reciterId", quranHandler.GetAyahAudioFileByAyahAndReciterHandler().Handle)
				ayahAudioGroup.GET("/surah/:surahId/reciter/:reciterId", quranHandler.GetAyahAudioFilesBySurahAndReciterHandler().Handle)

				ayahAudioProtected := ayahAudioGroup.Group("")
				ayahAudioProtected.Use(authMW)
				ayahAudioProtected.POST("", quranHandler.CreateAyahAudioFileHandler().Handle)
				ayahAudioProtected.PUT("/:id", quranHandler.UpdateAyahAudioFileHandler().Handle)
				ayahAudioProtected.DELETE("/:id", quranHandler.DeleteAyahAudioFileHandler().Handle)
			}

			// Bookmark Ayah routes
			bookmarkAyahGroup := quranGroup.Group("/bookmarks/ayahs")
			{
				bookmarkAyahGroup.GET("/:id", quranHandler.GetBookmarkAyahByIdHandler().Handle)
				bookmarkAyahGroup.GET("/user/:userId", quranHandler.GetBookmarkAyahsByUserHandler().Handle)
				bookmarkAyahGroup.GET("/user/:userId/ayah/:ayahId", quranHandler.GetBookmarkAyahByUserAndAyahHandler().Handle)

				bookmarkAyahProtected := bookmarkAyahGroup.Group("")
				bookmarkAyahProtected.Use(authMW)
				bookmarkAyahProtected.POST("", quranHandler.CreateBookmarkAyahHandler().Handle)
				bookmarkAyahProtected.DELETE("/:id", quranHandler.DeleteBookmarkAyahHandler().Handle)
			}

			// LastRead routes
			lastReadGroup := quranGroup.Group("/last-reads")
			{
				lastReadGroup.GET("/:id", quranHandler.GetLastReadByIdHandler().Handle)
				lastReadGroup.GET("/user/:userId", quranHandler.GetLastReadsByUserHandler().Handle)
				lastReadGroup.GET("/user/:userId/surah/:surahId", quranHandler.GetLastReadByUserAndSurahHandler().Handle)

				lastReadProtected := lastReadGroup.Group("")
				lastReadProtected.Use(authMW)
				lastReadProtected.POST("", quranHandler.CreateLastReadHandler().Handle)
				lastReadProtected.PUT("/:id", quranHandler.UpdateLastReadHandler().Handle)
				lastReadProtected.DELETE("/:id", quranHandler.DeleteLastReadHandler().Handle)
			}

			// Progress Hatam routes
			progressHatamGroup := quranGroup.Group("/progress-hatam")
			{
				progressHatamGroup.GET("/:id", quranHandler.GetProgressHatamByIdHandler().Handle)
				progressHatamGroup.GET("/user/:userId", quranHandler.GetProgressHatamByUserHandler().Handle)
				progressHatamGroup.GET("/user/:userId/juz/:juzId", quranHandler.GetProgressHatamByUserAndJuzHandler().Handle)

				progressHatamProtected := progressHatamGroup.Group("")
				progressHatamProtected.Use(authMW)
				progressHatamProtected.POST("", quranHandler.CreateProgressHatamHandler().Handle)
				progressHatamProtected.PUT("/:id", quranHandler.UpdateProgressHatamHandler().Handle)
				progressHatamProtected.DELETE("/:id", quranHandler.DeleteProgressHatamHandler().Handle)
			}
		}
	}
}
