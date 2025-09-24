package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/achmdndy/safa-life-api/src/presentation/handlers/quran"
)

// QuranRoutes sets up all routes for the Quran module
func QuranRoutes(router *gin.RouterGroup, handler *quran.Handler) {
	// Surah routes
	surahGroup := router.Group("/surahs")
	{
		// Read operations
		surahGroup.GET("", handler.GetAllSurahs)           // GET /surahs
		surahGroup.GET("/:id", handler.GetSurahByID)       // GET /surahs/:id
		surahGroup.GET("/:id/ayahs", handler.GetSurahWithAyahs) // GET /surahs/:id/ayahs
		surahGroup.GET("/:id/ayahs/list", handler.GetAyahsBySurah) // GET /surahs/:id/ayahs/list

		// Write operations
		surahGroup.POST("", handler.CreateSurah)           // POST /surahs
		surahGroup.PUT("/:id", handler.UpdateSurah)        // PUT /surahs/:id
		surahGroup.DELETE("/:id", handler.DeleteSurah)     // DELETE /surahs/:id
	}

	// Ayah routes
	ayahGroup := router.Group("/ayahs")
	{
		// Read operations
		ayahGroup.GET("/:surahId/:ayahId", handler.GetAyahByID)    // GET /ayahs/:surahId/:ayahId
		ayahGroup.GET("/page/:page", handler.GetAyahsByPage)       // GET /ayahs/page/:page

		// Write operations
		ayahGroup.POST("", handler.CreateAyah)                     // POST /ayahs
		ayahGroup.PUT("/:surahId/:ayahId", handler.UpdateAyah)     // PUT /ayahs/:surahId/:ayahId
		ayahGroup.DELETE("/:surahId/:ayahId", handler.DeleteAyah)  // DELETE /ayahs/:surahId/:ayahId
	}

	// Juz routes
	juzGroup := router.Group("/juz")
	{
		// Read operations
		juzGroup.GET("", handler.GetAllJuz)                // GET /juz
		juzGroup.GET("/:id", handler.GetJuzByID)           // GET /juz/:id
		juzGroup.GET("/:id/content", handler.GetJuzWithContent) // GET /juz/:id/content
		juzGroup.GET("/:id/ayahs", handler.GetAyahsByJuz)  // GET /juz/:id/ayahs

		// Write operations
		juzGroup.POST("", handler.CreateJuz)               // POST /juz
		juzGroup.PUT("/:id", handler.UpdateJuz)            // PUT /juz/:id
		juzGroup.DELETE("/:id", handler.DeleteJuz)         // DELETE /juz/:id
	}

	// Search routes
	router.GET("/search", handler.SearchHandler) // GET /search?q=query&page=1&limit=10
}

// SetupQuranRoutes is a convenience function to set up Quran routes with a specific prefix
func SetupQuranRoutes(router *gin.Engine, handler *quran.Handler) {
	api := router.Group("/api/v1/quran")
	QuranRoutes(api, handler)
}