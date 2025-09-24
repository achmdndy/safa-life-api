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

// UpdateAyahTafsir handles PUT /tafsirs/ayahs/:tafsir_id/:surah_id/:ayah_id
// @Summary Update Ayah Tafsir
// @Description Update an existing ayah tafsir.
// @Tags Ayah Tafsirs
// @Accept json
// @Produce json
// @Param tafsir_id path string true "Tafsir ID"
// @Param surah_id path int true "Surah ID"
// @Param ayah_id path int true "Ayah ID"
// @Param ayah_tafsir body dto.UpdateAyahTafsirRequest true "Ayah Tafsir Data"
// @Success 200 {object} core.SuccessResponse{data=dto.AyahTafsirResponse}
// @Failure 400 {object} core.ErrorResponse
// @Failure 404 {object} core.ErrorResponse "Tafsir not found"
// @Failure 500 {object} core.ErrorResponse
// @Router /tafsirs/ayahs/{tafsir_id}/{surah_id}/{ayah_id} [put]
func (h *Handler) UpdateAyahTafsir(c *gin.Context) {
	start := time.Now()
	var req dto.UpdateAyahTafsirRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid URI parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromUpdateAyahTafsirRequest(req)
	result, err := h.commandHandler.UpdateAyahTafsir(c.Request.Context(), cmd)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Tafsir not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to update ayah tafsir", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result == nil {
		core.Error(c, http.StatusNotFound, "Tafsir not found", &core.ErrorDetail{}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah tafsir updated successfully", dto.ToAyahTafsirResponse(*result), start)
}
