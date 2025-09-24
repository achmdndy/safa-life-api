package quran

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/achmdndy/safa-life-api/src/application/quran/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
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
		core.Error(c, http.StatusInternalServerError, "Failed to get juz", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result != nil {
		core.Success(c, http.StatusOK, "Juz retrieved successfully", result, start)
	} else {
		core.Success(c, http.StatusOK, "Juz retrieved successfully", core.EmptyData{}, start)
	}
}