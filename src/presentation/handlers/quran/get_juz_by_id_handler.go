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

// GetJuzByID handles GET /juz/:id
// @Summary Get juz by ID
// @Description Get a specific juz (part) by its ID
// @Tags Juz
// @Accept json
// @Produce json
// @Param id path int true "Juz ID"
// @Success 200 {object} core.SuccessResponse{data=dto.JuzResponse} "Juz retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid juz ID"
// @Failure 404 {object} core.ErrorResponse "Juz not found"
// @Failure 500 {object} core.ErrorResponse "Failed to get juz"
// @Router /juz/{id} [get]
func (h *Handler) GetJuzByID(c *gin.Context) {
	start := time.Now()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid juz ID", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.GetJuzByIDQuery{ID: id}
	result, err := h.queryHandler.GetJuzByID(c.Request.Context(), queryReq)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Juz not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get juz", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result == nil {
		core.Error(c, http.StatusNotFound, "Juz not found", &core.ErrorDetail{}, start)
		return
	}

	core.Success(c, http.StatusOK, "Juz retrieved successfully", result, start)
}
