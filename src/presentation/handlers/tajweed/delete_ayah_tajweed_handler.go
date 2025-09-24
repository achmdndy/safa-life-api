package tajweed

import (
	"net/http"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tajweed/command"
	"github.com/achmdndy/safa-life-api/src/application/tajweed/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// DeleteAyahTajweed handles DELETE /tajweed/ayah/:tajweed_id/:surah_id/:ayah_id
// DeleteAyahTajweed handles DELETE /tajweed/ayah/:tajweed_id/:surah_id/:ayah_id
// @Summary Delete Ayah Tajweed
// @Description Delete the set of tajweed rules for a specific ayah.
// @Tags Tajweed
// @Accept json
// @Produce json
// @Param tajweed_id path string true "Tajweed ID"
// @Param surah_id path int true "Surah ID"
// @Param ayah_id path int true "Ayah ID"
// @Success 200 {object} core.SuccessResponse "Ayah tajweed deleted successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters"
// @Failure 404 {object} core.ErrorResponse "Ayah tajweed not found"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /tajweed/ayah/{tajweed_id}/{surah_id}/{ayah_id} [delete]
func (h *Handler) DeleteAyahTajweed(c *gin.Context) {
	start := time.Now()
	var req dto.GetAyahTajweedRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.DeleteAyahTajweedCommand{TajweedID: req.TajweedID, SurahID: req.SurahID, AyahID: req.AyahID}
	err := h.commandHandler.DeleteAyahTajweed(c.Request.Context(), cmd)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Ayah tajweed not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to delete ayah tajweed", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah tajweed deleted successfully", core.EmptyData{}, start)
}
