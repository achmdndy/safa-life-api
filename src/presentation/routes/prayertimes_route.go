package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/presentation/handlers/prayertimes"
)

// PrayerTimesRoutes sets up prayer times-related routes
func PrayerTimesRoutes(router *gin.RouterGroup, handler *prayertimes.Handler) {
	ptGroup := router.Group("/prayertimes")
	{
		ptGroup.GET("/daily", handler.GetDailyTimingsHandler().Handle)
		ptGroup.GET("/monthly", handler.GetMonthlyTimingsHandler().Handle)
		ptGroup.GET("/yearly", handler.GetYearlyTimingsHandler().Handle)
	}
}
