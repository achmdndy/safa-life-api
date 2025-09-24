package quran

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/achmdndy/safa-life-api/src/application/quran/command"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

// DeleteJuz handles DELETE /juz/:id
// @Summary Delete a juz
// @Description Delete a specific juz (part) from the Quran
// @Tags Juz
// @Accept json
// @Produce json
// @Param id path int true "Juz ID"
// @Success 200 {object} core.SuccessResponse "Juz deleted successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid juz ID"
// @Failure 500 {object} core.ErrorResponse "Failed to delete juz"
// @Router /juz/{id} [delete]
func (h *Handler) DeleteJuz(c *gin.Context) {
	start := time.Now()
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid juz ID", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.DeleteJuzCommand{ID: id}
	err = h.commandHandler.DeleteJuz(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to delete juz", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Juz deleted successfully", core.EmptyData{}, start)
}