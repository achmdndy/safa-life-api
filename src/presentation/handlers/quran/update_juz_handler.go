package quran

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/quran/command"
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// UpdateJuz handles PUT /juz/:id
// @Summary Update an existing juz
// @Description Update an existing juz (part) in the Quran
// @Tags Juz
// @Accept json
// @Produce json
// @Param id path int true "Juz ID"
// @Param juz body dto.UpdateJuzRequest true "Juz data"
// @Success 200 {object} core.SuccessResponse{data=dto.JuzResponse} "Juz updated successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid juz ID or request body"
// @Failure 404 {object} core.ErrorResponse "Juz not found"
// @Failure 500 {object} core.ErrorResponse "Failed to update juz"
// @Router /juz/{id} [put]
func (h *Handler) UpdateJuz(c *gin.Context) {
	start := time.Now()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid juz ID", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	var req dto.UpdateJuzRequest
	if bindErr := c.ShouldBindJSON(&req); bindErr != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: bindErr.Error()}, start)
		return
	}

	cmd := command.FromUpdateJuzRequest(req)
	cmd.ID = id
	domainResult, err := h.commandHandler.UpdateJuz(c.Request.Context(), cmd)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Juz not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to update juz", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if domainResult == nil {
		core.Error(c, http.StatusNotFound, "Juz not found", &core.ErrorDetail{}, start)
		return
	}

	result := dto.ToJuzResponse(*domainResult)
	core.Success(c, http.StatusOK, "Juz updated successfully", result, start)
}
