package tafsir

import (
	"net/http"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tafsir/command"
	"github.com/achmdndy/safa-life-api/src/application/tafsir/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// UpdateTafsir handles PUT /tafsirs/editions/:id
// @Summary Update Tafsir Edition
// @Description Update an existing tafsir edition.
// @Tags Tafsir Editions
// @Accept json
// @Produce json
// @Param id path string true "Tafsir ID"
// @Param edition body dto.UpdateTafsirRequest true "Tafsir Edition Data"
// @Success 200 {object} core.SuccessResponse{data=dto.TafsirResponse}
// @Failure 400 {object} core.ErrorResponse
// @Failure 404 {object} core.ErrorResponse "Tafsir edition not found"
// @Failure 500 {object} core.ErrorResponse
// @Router /tafsirs/editions/{id} [put]
func (h *Handler) UpdateTafsir(c *gin.Context) {
	start := time.Now()
	var req dto.UpdateTafsirRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid URI parameter", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromUpdateTafsirRequest(req)
	result, err := h.commandHandler.UpdateTafsir(c.Request.Context(), cmd)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Tafsir edition not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to update tafsir edition", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result == nil {
		core.Error(c, http.StatusNotFound, "Tafsir edition not found", &core.ErrorDetail{}, start)
		return
	}

	core.Success(c, http.StatusOK, "Tafsir edition updated successfully", dto.ToTafsirResponse(*result), start)
}
