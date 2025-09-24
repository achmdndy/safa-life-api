package quran

import (
	"net/http"
	"strconv"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/quran/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetAyahsByJuz handles GET /juz/:id/ayahs
// @Summary Get ayahs by juz
// @Description Get all ayahs (verses) from a specific juz (part)
// @Tags Ayah
// @Accept json
// @Produce json
// @Param id path int true "Juz ID"
// @Success 200 {object} core.SuccessResponse "Ayahs retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid juz ID"
// @Failure 500 {object} core.ErrorResponse "Failed to get ayahs"
// @Router /juz/{id}/ayahs [get]
func (h *Handler) GetAyahsByJuz(c *gin.Context) {
	start := time.Now()

	juzID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid juz ID", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.GetAyahsByJuzQuery{JuzID: juzID}
	result, err := h.queryHandler.GetAyahsByJuz(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get ayahs", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result != nil {
		core.Success(c, http.StatusOK, "Ayahs retrieved successfully", result, start)
	} else {
		core.Success(c, http.StatusOK, "Ayahs retrieved successfully", core.EmptyData{}, start)
	}
}
