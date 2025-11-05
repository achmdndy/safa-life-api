package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
)

// DeleteSurahHandler handles the delete surah request
type DeleteSurahHandler struct {
	commandHandler *command.CommandHandler
}

// NewDeleteSurahHandler creates a new delete surah handler
func NewDeleteSurahHandler(commandHandler *command.CommandHandler) *DeleteSurahHandler {
	return &DeleteSurahHandler{
		commandHandler: commandHandler,
	}
}

// Handle processes the delete surah request
// @Summary Delete a surah
// @Description Delete an existing surah from the Quran
// @Tags Surahs
// @Accept json
// @Produce json
// @Param id path string true "Surah ID"
// @Success 200 {object} DeleteSurahSuccessResponse "Surah deleted successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Surah not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/surahs/{id} [delete]
func (h *DeleteSurahHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.DeleteSurahRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid surah ID", errorDetail, start)
		return
	}

	cmd := command.DeleteSurahCommand{
		ID: req.ID,
	}

	err := h.commandHandler.DeleteSurah(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to delete surah", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Surah deleted successfully", map[string]interface{}{}, start)
}
