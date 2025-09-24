package quran

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/quran/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetSurahByID handles GET /surahs/:id
// @Summary Get surah by ID
// @Description Get a specific surah (chapter) by its ID
// @Tags Surah
// @Accept json
// @Produce json
// @Param id path int true "Surah ID"
// @Success 200 {object} core.SuccessResponse{data=dto.SurahResponse} "Surah retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid surah ID"
// @Failure 404 {object} core.ErrorResponse "Surah not found"
// @Failure 500 {object} core.ErrorResponse "Failed to get surah"
// @Router /surahs/{id} [get]
func (h *Handler) GetSurahByID(c *gin.Context) {
	start := time.Now()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid surah ID", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.GetSurahByIDQuery{ID: id}
	result, err := h.queryHandler.GetSurahByID(c.Request.Context(), queryReq)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Surah not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get surah", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result == nil {
		core.Error(c, http.StatusNotFound, "Surah not found", &core.ErrorDetail{}, start)
		return
	}

	core.Success(c, http.StatusOK, "Surah retrieved successfully", result, start)
}
