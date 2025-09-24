package quran

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/achmdndy/safa-life-api/src/application/quran/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

// GetAyahsBySurah handles GET /surahs/:id/ayahs
// @Summary Get ayahs by surah
// @Description Get all ayahs (verses) from a specific surah (chapter) with pagination
// @Tags Ayah
// @Accept json
// @Produce json
// @Param id path int true "Surah ID"
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Number of items per page (default: 10)"
// @Success 200 {object} core.SuccessResponse "Ayahs retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid surah ID or pagination parameters"
// @Failure 500 {object} core.ErrorResponse "Failed to get ayahs"
// @Router /surahs/{id}/ayahs [get]
func (h *Handler) GetAyahsBySurah(c *gin.Context) {
	start := time.Now()
	
	surahID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid surah ID", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	queryReq := query.GetAyahsBySurahQuery{
		SurahID: surahID,
		Page:    page,
		Limit:   limit,
	}

	result, err := h.queryHandler.GetAyahsBySurah(c.Request.Context(), queryReq)
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