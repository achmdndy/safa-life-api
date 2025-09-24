package quran

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/achmdndy/safa-life-api/src/application/quran/command"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

// DeleteSurah handles DELETE /surahs/:id
// @Summary Delete a surah
// @Description Delete a specific surah (chapter) from the Quran
// @Tags Surah
// @Accept json
// @Produce json
// @Param id path int true "Surah ID"
// @Success 200 {object} core.SuccessResponse "Surah deleted successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid surah ID"
// @Failure 500 {object} core.ErrorResponse "Failed to delete surah"
// @Router /surahs/{id} [delete]
func (h *Handler) DeleteSurah(c *gin.Context) {
	start := time.Now()
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid surah ID", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.DeleteSurahCommand{ID: id}
	err = h.commandHandler.DeleteSurah(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to delete surah", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Surah deleted successfully", core.EmptyData{}, start)
}