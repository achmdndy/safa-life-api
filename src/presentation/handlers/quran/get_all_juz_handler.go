package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/achmdndy/safa-life-api/src/application/quran/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

// GetAllJuz handles GET /juz
// @Summary Get all juz
// @Description Get a list of all juz (parts) in the Quran
// @Tags Juz
// @Accept json
// @Produce json
// @Success 200 {object} core.SuccessResponse{data=[]dto.JuzResponse} "Juz retrieved successfully"
// @Failure 500 {object} core.ErrorResponse "Failed to get juz"
// @Router /juz [get]
func (h *Handler) GetAllJuz(c *gin.Context) {
	start := time.Now()
	
	queryReq := query.GetAllJuzQuery{}
	result, err := h.queryHandler.GetAllJuz(c.Request.Context(), queryReq)
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