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

// UpdateAyahHandler handles the update ayah request
type UpdateAyahHandler struct {
	commandHandler *command.CommandHandler
}

// NewUpdateAyahHandler creates a new update ayah handler
func NewUpdateAyahHandler(commandHandler *command.CommandHandler) *UpdateAyahHandler {
	return &UpdateAyahHandler{
		commandHandler: commandHandler,
	}
}

// Handle processes the update ayah request
// @Summary Update an ayah
// @Description Update an existing ayah in the Quran
// @Tags Ayahs
// @Accept json
// @Produce json
// @Param id path string true "Ayah ID"
// @Param request body dto.UpdateAyahRequest true "Update ayah request"
// @Success 200 {object} UpdateAyahSuccessResponse "Ayah updated successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Ayah not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/ayahs/{id} [put]
func (h *UpdateAyahHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.UpdateAyahRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid ayah ID", errorDetail, start)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid request body", errorDetail, start)
		return
	}

	cmd := command.UpdateAyahCommand{
		ID:           req.ID,
		Text:         req.Text,
		PageNumber:   req.PageNumber,
		JuzNumber:    req.JuzNumber,
		HizbNumber:   req.HizbNumber,
		ManzilNumber: req.ManzilNumber,
		UpdatedBy:    middlewares.GetUserID(c),
	}

	result, err := h.commandHandler.UpdateAyah(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to update ayah", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah updated successfully", result, start)
}
