package quran

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/achmdndy/safa-life-api/src/application/quran/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

// GetSurahWithAyahs handles GET /surahs/:id/ayahs
// @Summary Get surah with its ayahs
// @Description Get a specific surah (chapter) along with all its ayahs (verses)
// @Tags Surah
// @Accept json
// @Produce json
// @Param id path int true "Surah ID"
// @Success 200 {object} core.SuccessResponse "Surah with ayahs retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid surah ID"
// @Failure 500 {object} core.ErrorResponse "Failed to get surah with ayahs"
// @Router /surahs/{id}/ayahs [get]
func (h *Handler) GetSurahWithAyahs(c *gin.Context) {
	start := time.Now()
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid surah ID", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.GetSurahWithAyahsQuery{ID: id}
	result, err := h.queryHandler.GetSurahWithAyahs(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get surah with ayahs", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result != nil {
		core.Success(c, http.StatusOK, "Surah with ayahs retrieved successfully", result, start)
	} else {
		core.Success(c, http.StatusOK, "Surah with ayahs retrieved successfully", core.EmptyData{}, start)
	}
}