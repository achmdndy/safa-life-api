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

// GetJuzWithContent handles GET /juz/:id/content
// @Summary Get juz with its content
// @Description Get a specific juz (part) along with all its content (ayahs)
// @Tags Juz
// @Accept json
// @Produce json
// @Param id path int true "Juz ID"
// @Success 200 {object} core.SuccessResponse "Juz with content retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid juz ID"
// @Failure 404 {object} core.ErrorResponse "Juz not found"
// @Failure 500 {object} core.ErrorResponse "Failed to get juz with content"
// @Router /juz/{id}/content [get]
func (h *Handler) GetJuzWithContent(c *gin.Context) {
	start := time.Now()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid juz ID", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.GetJuzWithContentQuery{ID: id}
	result, err := h.queryHandler.GetJuzWithContent(c.Request.Context(), queryReq)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Juz not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get juz with content", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result == nil {
		core.Error(c, http.StatusNotFound, "Juz not found", &core.ErrorDetail{}, start)
		return
	}

	core.Success(c, http.StatusOK, "Juz with content retrieved successfully", result, start)
}
