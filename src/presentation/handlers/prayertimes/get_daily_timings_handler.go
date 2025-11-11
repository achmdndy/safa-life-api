package prayertimes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	appdto "github.com/safalife/core-api/src/application/prayertimes/dto"
	appquery "github.com/safalife/core-api/src/application/prayertimes/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetDailyTimingsHandler handles the get daily prayer timings request
type GetDailyTimingsHandler struct {
	queryHandler *appquery.QueryHandler
}

// NewGetDailyTimingsHandler creates a new handler
func NewGetDailyTimingsHandler(queryHandler *appquery.QueryHandler) *GetDailyTimingsHandler {
	return &GetDailyTimingsHandler{queryHandler: queryHandler}
}

// Handle processes the get daily prayer timings request
// @Summary Get daily prayer timings
// @Description Retrieve daily prayer timings for a given location and date
// @Tags PrayerTimes
// @Accept json
// @Produce json
// @Param latitude query number true "Latitude" example(-6.2000)
// @Param longitude query number true "Longitude" example(106.8166)
// @Param timezone query string true "Timezone (IANA)" example("Asia/Jakarta")
// @Param date query string false "Date (YYYY-MM-DD)" example("2024-01-01")
// @Param method query int false "Calculation method (provider-specific)" example(5)
// @Success 200 {object} PrayerTimesSuccessResponse "Prayer times retrieved successfully"
// @Failure 400 {object} PrayerTimesErrorResponse "Bad request"
// @Failure 500 {object} PrayerTimesErrorResponse "Internal server error"
// @Router /prayertimes/daily [get]
func (h *GetDailyTimingsHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req appdto.GetDailyTimingsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid query parameters", errorDetail, start)
		return
	}

	// If date not provided, default to today's date in the specified timezone
	if req.Date.IsZero() {
		loc := time.UTC
		if req.Timezone != "" {
			if l, err := time.LoadLocation(req.Timezone); err == nil {
				loc = l
			}
		}
		now := time.Now().In(loc)
		req.Date = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	}

	result, err := h.queryHandler.GetDailyTimings(ctx, &req)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get prayer times", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Prayer times retrieved successfully", result, start)
}
