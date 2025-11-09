package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
	"github.com/safalife/core-api/src/presentation/middlewares"
)

// UpdateProgressHatamHandler handles updating a progress hatam entry
type UpdateProgressHatamHandler struct {
	commandHandler *command.CommandHandler
}

// NewUpdateProgressHatamHandler constructs the handler
func NewUpdateProgressHatamHandler(commandHandler *command.CommandHandler) *UpdateProgressHatamHandler {
	return &UpdateProgressHatamHandler{commandHandler: commandHandler}
}

// Handle processes the request
// @Summary Update progress hatam
// @Description Update fields of an existing progress hatam entry
// @Tags ProgressHatam
// @Accept json
// @Produce json
// @Param id path string true "ProgressHatam ID"
// @Param progressHatam body dto.UpdateProgressHatamRequest true "ProgressHatam update information"
// @Success 200 {object} ProgressHatamSuccessResponse "Progress hatam updated successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Progress hatam not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/progress-hatam/{id} [put]
func (h *UpdateProgressHatamHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var uri dto.UpdateProgressHatamRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid progress hatam ID", errorDetail, start)
		return
	}

	var req dto.UpdateProgressHatamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid request body", errorDetail, start)
		return
	}

	cmd := command.UpdateProgressHatamCommand{
		ID:          uri.ID,
		LastAyahID:  req.LastAyahID,
		ProgressPct: req.ProgressPct,
		IsCompleted: req.IsCompleted,
		CompletedAt: req.CompletedAt,
		UpdatedBy:   middlewares.GetUserID(c),
	}

	result, err := h.commandHandler.UpdateProgressHatam(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to update progress hatam", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Progress hatam updated successfully", result, start)
}
