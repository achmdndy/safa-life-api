package quran

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/achmdndy/safa-life-api/src/application/quran/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

// GetAyahsByPage handles GET /pages/:page/ayahs
// @Summary Get ayahs by page
// @Description Get all ayahs (verses) from a specific page in the Quran
// @Tags Ayah
// @Accept json
// @Produce json
// @Param page path int true "Page number"
// @Success 200 {object} core.SuccessResponse "Ayahs retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid page number"
// @Failure 500 {object} core.ErrorResponse "Failed to get ayahs"
// @Router /pages/{page}/ayahs [get]
func (h *Handler) GetAyahsByPage(c *gin.Context) {
	start := time.Now()
	
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid page number", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.GetAyahsByPageQuery{Page: page}
	result, err := h.queryHandler.GetAyahsByPage(c.Request.Context(), queryReq)
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