package reciter

import (
	"net/http"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/reciter/dto"
	"github.com/achmdndy/safa-life-api/src/application/reciter/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetReciterByID handles GET /reciters/:id
// GetReciterByID handles GET /reciters/:id
// @Summary Get Reciter By ID
// @Description Retrieve a specific reciter by their ID.
// @Tags Reciters
// @Accept json
// @Produce json
// @Param id path string true "Reciter ID"
// @Success 200 {object} core.SuccessResponse{data=dto.ReciterResponse} "Reciter retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid reciter ID"
// @Failure 404 {object} core.ErrorResponse "Reciter not found"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /reciters/{id} [get]
func (h *Handler) GetReciterByID(c *gin.Context) {
	start := time.Now()
	var req dto.GetReciterByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.ToGetReciterByIDQuery(req.ID)
	result, err := h.queryHandler.GetReciterByID(c.Request.Context(), queryReq)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Reciter not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get reciter", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result == nil {
		core.Error(c, http.StatusNotFound, "Reciter not found", &core.ErrorDetail{}, start)
		return
	}

	core.Success(c, http.StatusOK, "Reciter retrieved successfully", result, start)
}
