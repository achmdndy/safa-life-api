package reciter

import (
	"net/http"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/reciter/command"
	"github.com/achmdndy/safa-life-api/src/application/reciter/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// UpdateReciter handles PUT /reciters/:id
// UpdateReciter handles PUT /reciters/:id
// @Summary Update Reciter
// @Description Update an existing reciter's information.
// @Tags Reciters
// @Accept json
// @Produce json
// @Param id path string true "Reciter ID"
// @Param reciter body dto.UpdateReciterRequest true "Reciter Data"
// @Success 200 {object} core.SuccessResponse{data=dto.ReciterResponse} "Reciter updated successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters or request body"
// @Failure 404 {object} core.ErrorResponse "Reciter not found"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /reciters/{id} [put]
func (h *Handler) UpdateReciter(c *gin.Context) {
	start := time.Now()
	var req dto.UpdateReciterRequest

	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid URI parameter", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromUpdateReciterRequest(req)
	result, err := h.commandHandler.UpdateReciter(c.Request.Context(), cmd)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Reciter not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to update reciter", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result == nil {
		core.Error(c, http.StatusNotFound, "Reciter not found", &core.ErrorDetail{}, start)
		return
	}

	core.Success(c, http.StatusOK, "Reciter updated successfully", dto.ToReciterResponse(*result), start)
}
