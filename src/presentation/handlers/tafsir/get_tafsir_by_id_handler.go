package tafsir

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tafsir/dto"
	"github.com/achmdndy/safa-life-api/src/application/tafsir/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetTafsirByID handles GET /tafsirs/editions/:id
// @Summary Get Tafsir Edition By ID
// @Description Get a specific tafsir edition by ID.
// @Tags Tafsir Editions
// @Accept json
// @Produce json
// @Param id path string true "Tafsir ID"
// @Success 200 {object} core.SuccessResponse{data=dto.TafsirResponse}
// @Failure 400 {object} core.ErrorResponse
// @Failure 500 {object} core.ErrorResponse
// @Router /tafsirs/editions/{id} [get]
func (h *Handler) GetTafsirByID(c *gin.Context) {
	start := time.Now()
	var req dto.GetTafsirByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.ToGetTafsirByIDQuery(req.ID)
	result, err := h.queryHandler.GetTafsirByID(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get tafsir edition", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Tafsir edition retrieved successfully", result, start)
}
