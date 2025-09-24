package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/achmdndy/safa-life-api/src/application/quran/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

// GetAllSurahs handles GET /surahs
// @Summary Get all surahs
// @Description Get a list of all surahs (chapters) in the Quran
// @Tags Surah
// @Accept json
// @Produce json
// @Success 200 {object} core.SuccessResponse{data=[]dto.SurahResponse} "Surahs retrieved successfully"
// @Failure 500 {object} core.ErrorResponse "Failed to get surahs"
// @Router /surahs [get]
func (h *Handler) GetAllSurahs(c *gin.Context) {
	start := time.Now()
	
	queryReq := query.GetAllSurahsQuery{}
	result, err := h.queryHandler.GetAllSurahs(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get surahs", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result != nil {
		core.Success(c, http.StatusOK, "Surahs retrieved successfully", result, start)
	} else {
		core.Success(c, http.StatusOK, "Surahs retrieved successfully", core.EmptyData{}, start)
	}
}