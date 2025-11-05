package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
)

// UpdateReciterHandler handles the update reciter request
type UpdateReciterHandler struct {
	commandHandler *command.CommandHandler
}

// NewUpdateReciterHandler creates a new update reciter handler
func NewUpdateReciterHandler(commandHandler *command.CommandHandler) *UpdateReciterHandler {
	return &UpdateReciterHandler{commandHandler: commandHandler}
}

// Handle processes the update reciter request
// @Summary Update a reciter
// @Description Update an existing reciter
// @Tags Reciters
// @Accept json
// @Produce json
// @Param id path string true "Reciter ID"
// @Param request body dto.UpdateReciterRequest true "Update reciter request"
// @Success 200 {object} UpdateReciterSuccessResponse "Reciter updated successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/reciters/{id} [put]
func (h *UpdateReciterHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var uri dto.UpdateReciterRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid URI parameters", errorDetail, start)
		return
	}

	var body dto.UpdateReciterRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid request body", errorDetail, start)
		return
	}

	cmd := command.UpdateReciterCommand{
		ID:        uri.ID,
		Name:      body.Name,
		Style:     body.Style,
		UpdatedBy: body.UpdatedBy,
	}

	result, err := h.commandHandler.UpdateReciter(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to update reciter", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Reciter updated successfully", result, start)
}
