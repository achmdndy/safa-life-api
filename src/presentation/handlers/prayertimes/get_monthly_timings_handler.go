package prayertimes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	appdto "github.com/safalife/core-api/src/application/prayertimes/dto"
	appquery "github.com/safalife/core-api/src/application/prayertimes/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetMonthlyTimingsHandler handles the get monthly prayer timings request
type GetMonthlyTimingsHandler struct {
	queryHandler *appquery.QueryHandler
}

// NewGetMonthlyTimingsHandler creates a new handler
func NewGetMonthlyTimingsHandler(queryHandler *appquery.QueryHandler) *GetMonthlyTimingsHandler {
	return &GetMonthlyTimingsHandler{queryHandler: queryHandler}
}

// Handle processes the get monthly prayer timings request
// @Summary Get monthly prayer timings
// @Description Retrieve monthly prayer timings for a given location and month/year
// @Tags PrayerTimes
// @Accept json
// @Produce json
// @Param latitude query number true "Latitude" example(-6.2000)
// @Param longitude query number true "Longitude" example(106.8166)
// @Param timezone query string true "Timezone (IANA)" example("Asia/Jakarta")
// @Param month query int true "Month (1-12)" example(1)
// @Param year query int true "Year" example(2024)
// @Param method query int false "Calculation method (provider-specific)" example(5)
// @Success 200 {object} PrayerTimesListSuccessResponse "Prayer times list retrieved successfully"
// @Failure 400 {object} PrayerTimesErrorResponse "Bad request"
// @Failure 500 {object} PrayerTimesErrorResponse "Internal server error"
// @Router /prayertimes/monthly [get]
func (h *GetMonthlyTimingsHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req appdto.GetMonthlyTimingsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid query parameters", errorDetail, start)
		return
	}

	result, err := h.queryHandler.GetMonthlyTimings(ctx, &req)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get monthly prayer times", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Prayer times list retrieved successfully", result, start)
}
