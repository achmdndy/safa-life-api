package tafsir

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tafsir/command"
	"github.com/achmdndy/safa-life-api/src/application/tafsir/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// DeleteAyahTafsir handles DELETE /tafsirs/ayahs/:tafsir_id/:surah_id/:ayah_id
// @Summary Delete Ayah Tafsir
// @Description Delete a specific ayah tafsir.
// @Tags Ayah Tafsirs
// @Accept json
// @Produce json
// @Param tafsir_id path string true "Tafsir ID"
// @Param surah_id path int true "Surah ID"
// @Param ayah_id path int true "Ayah ID"
// @Success 200 {object} core.SuccessResponse
// @Failure 400 {object} core.ErrorResponse
// @Failure 500 {object} core.ErrorResponse
// @Router /tafsirs/ayahs/{tafsir_id}/{surah_id}/{ayah_id} [delete]
func (h *Handler) DeleteAyahTafsir(c *gin.Context) {
	start := time.Now()
	var req dto.GetAyahTafsirRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.DeleteAyahTafsirCommand(req)
	if err := h.commandHandler.DeleteAyahTafsir(c.Request.Context(), cmd); err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to delete ayah tafsir", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah tafsir deleted successfully", core.EmptyData{}, start)
}
