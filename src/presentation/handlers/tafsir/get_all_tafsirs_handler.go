package tafsir

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tafsir/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetAllTafsirs handles GET /tafsirs/editions
// @Summary Get All Tafsir Editions
// @Description Get a list of all tafsir editions.
// @Tags Tafsir Editions
// @Accept json
// @Produce json
// @Success 200 {object} core.SuccessResponse{data=[]dto.TafsirResponse}
// @Failure 500 {object} core.ErrorResponse
// @Router /tafsirs/editions [get]
func (h *Handler) GetAllTafsirs(c *gin.Context) {
	start := time.Now()
	queryReq := query.ToGetAllTafsirsQuery()
	result, err := h.queryHandler.GetAllTafsirs(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get tafsir editions", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Tafsir editions retrieved successfully", result, start)
}
