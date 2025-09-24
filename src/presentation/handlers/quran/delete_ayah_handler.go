package quran

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/achmdndy/safa-life-api/src/application/quran/command"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

// DeleteAyah handles DELETE /ayahs/:surahId/:ayahId
// @Summary Delete an ayah
// @Description Delete an existing ayah (verse) from the Quran
// @Tags Ayah
// @Accept json
// @Produce json
// @Param surahId path int true "Surah ID"
// @Param ayahId path int true "Ayah ID"
// @Success 200 {object} core.SuccessResponse "Ayah deleted successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid surah ID or ayah ID"
// @Failure 500 {object} core.ErrorResponse "Failed to delete ayah"
// @Router /ayahs/{surahId}/{ayahId} [delete]
func (h *Handler) DeleteAyah(c *gin.Context) {
	start := time.Now()
	
	surahID, err := strconv.Atoi(c.Param("surahId"))
	if err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid surah ID", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	ayahID, err := strconv.Atoi(c.Param("ayahId"))
	if err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid ayah ID", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.DeleteAyahCommand{SurahID: surahID, AyahID: ayahID}
	err = h.commandHandler.DeleteAyah(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to delete ayah", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah deleted successfully", core.EmptyData{}, start)
}