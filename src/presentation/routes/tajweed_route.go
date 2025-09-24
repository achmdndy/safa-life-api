package routes

import (
	"github.com/achmdndy/safa-life-api/src/presentation/handlers/tajweed"
	"github.com/gin-gonic/gin"
)

// TajweedRoutes sets up all routes for the Tajweed module
func TajweedRoutes(router *gin.RouterGroup, handler *tajweed.Handler) {
	tajweedGroup := router.Group("/tajweed")
	{
		// AyahTajweed routes
		ayahGroup := tajweedGroup.Group("/ayah")
		{
			ayahGroup.GET("/:tajweed_id/:surah_id/:ayah_id", handler.GetAyahTajweed)
			ayahGroup.POST("", handler.CreateAyahTajweed)
			ayahGroup.PUT("/:tajweed_id/:surah_id/:ayah_id", handler.UpdateAyahTajweed)
			ayahGroup.DELETE("/:tajweed_id/:surah_id/:ayah_id", handler.DeleteAyahTajweed)
		}

		// TajweedRule routes
		ruleGroup := tajweedGroup.Group("/rules")
		{
			ruleGroup.GET("", handler.GetAllTajweedRules)
			ruleGroup.GET("/id/:id", handler.GetTajweedRuleByID)
			ruleGroup.GET("/name/:rule_name", handler.GetTajweedRuleByName)
			ruleGroup.POST("", handler.CreateTajweedRule)
			ruleGroup.PUT("/:id", handler.UpdateTajweedRule)
			ruleGroup.DELETE("/:id", handler.DeleteTajweedRule)
		}
	}
}