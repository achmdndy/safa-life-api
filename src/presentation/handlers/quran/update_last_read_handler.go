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

// UpdateLastReadHandler handles updating a last read entry
type UpdateLastReadHandler struct {
	commandHandler *command.CommandHandler
}

// NewUpdateLastReadHandler constructs the handler
func NewUpdateLastReadHandler(commandHandler *command.CommandHandler) *UpdateLastReadHandler {
	return &UpdateLastReadHandler{commandHandler: commandHandler}
}

// Handle processes the request
// @Summary Update last read
// @Description Update fields of an existing last read entry
// @Tags LastRead
// @Accept json
// @Produce json
// @Param id path string true "LastRead ID"
// @Param lastRead body dto.UpdateLastReadRequest true "LastRead update information"
// @Success 200 {object} LastReadSuccessResponse "Last read updated successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Last read not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/last-reads/{id} [put]
func (h *UpdateLastReadHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var uri dto.UpdateLastReadRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid last read ID", errorDetail, start)
		return
	}

	var req dto.UpdateLastReadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid request body", errorDetail, start)
		return
	}

	cmd := command.UpdateLastReadCommand{
		ID:          uri.ID,
		AyahID:      req.AyahID,
		AyahNumber:  req.AyahNumber,
		ProgressPct: req.ProgressPct,
		UpdatedBy:   middlewares.GetUserID(c),
	}

	result, err := h.commandHandler.UpdateLastRead(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to update last read", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Last read updated successfully", result, start)
}
