package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
)

// UpdateJuzHandler handles the update juz request
type UpdateJuzHandler struct {
	commandHandler *command.CommandHandler
}

// NewUpdateJuzHandler creates a new update juz handler
func NewUpdateJuzHandler(commandHandler *command.CommandHandler) *UpdateJuzHandler {
	return &UpdateJuzHandler{
		commandHandler: commandHandler,
	}
}

// Handle processes the update juz request
// @Summary Update a juz
// @Description Update an existing juz with the provided information
// @Tags Juz
// @Accept json
// @Produce json
// @Param id path string true "Juz ID"
// @Param juz body dto.UpdateJuzRequest true "Juz information"
// @Success 200 {object} JuzSuccessResponse "Juz updated successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Juz not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/juz/{id} [put]
func (h *UpdateJuzHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.UpdateJuzRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid juz ID", errorDetail, start)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid request body", errorDetail, start)
		return
	}

	cmd := command.UpdateJuzCommand{
		ID:           req.ID,
		StartSurahID: req.StartSurahID,
		EndSurahID:   req.EndSurahID,
		StartAyahID:  req.StartAyahID,
		EndAyahID:    req.EndAyahID,
		UpdatedBy:    req.UpdatedBy,
	}

	result, err := h.commandHandler.UpdateJuz(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to update juz", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Juz updated successfully", result, start)
}
