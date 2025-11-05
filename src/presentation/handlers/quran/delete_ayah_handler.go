package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
)

// DeleteAyahHandler handles the delete ayah request
type DeleteAyahHandler struct {
	commandHandler *command.CommandHandler
}

// NewDeleteAyahHandler creates a new delete ayah handler
func NewDeleteAyahHandler(commandHandler *command.CommandHandler) *DeleteAyahHandler {
	return &DeleteAyahHandler{
		commandHandler: commandHandler,
	}
}

// Handle processes the delete ayah request
// @Summary Delete an ayah
// @Description Delete an existing ayah from the Quran
// @Tags Ayahs
// @Accept json
// @Produce json
// @Param id path string true "Ayah ID"
// @Success 200 {object} DeleteAyahSuccessResponse "Ayah deleted successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Ayah not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/ayahs/{id} [delete]
func (h *DeleteAyahHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.DeleteAyahRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid ayah ID", errorDetail, start)
		return
	}

	cmd := command.DeleteAyahCommand{
		ID: req.ID,
	}

	err := h.commandHandler.DeleteAyah(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to delete ayah", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah deleted successfully", map[string]interface{}{}, start)
}
