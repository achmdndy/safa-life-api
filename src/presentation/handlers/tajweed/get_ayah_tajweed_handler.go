package tajweed

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tajweed/dto"
	"github.com/achmdndy/safa-life-api/src/application/tajweed/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetAyahTajweed handles GET /tajweed/ayah/:tajweed_id/:surah_id/:ayah_id
// GetAyahTajweed handles GET /tajweed/ayah/:tajweed_id/:surah_id/:ayah_id
// @Summary Get Ayah Tajweed
// @Description Retrieve the set of tajweed rules for a specific ayah.
// @Tags Tajweed
// @Accept json
// @Produce json
// @Param tajweed_id path string true "Tajweed ID"
// @Param surah_id path int true "Surah ID"
// @Param ayah_id path int true "Ayah ID"
// @Success 200 {object} core.SuccessResponse{data=dto.AyahTajweedResponse} "Ayah tajweed retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /tajweed/ayah/{tajweed_id}/{surah_id}/{ayah_id} [get]
func (h *Handler) GetAyahTajweed(c *gin.Context) {
	start := time.Now()
	var req dto.GetAyahTajweedRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.ToGetAyahTajweedQuery(req.TajweedID, req.SurahID, req.AyahID)
	result, err := h.queryHandler.GetAyahTajweed(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get ayah tajweed", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah tajweed retrieved successfully", result, start)
}
