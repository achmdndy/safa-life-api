package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
)

// CreateJuzHandler handles the create juz request
type CreateJuzHandler struct {
	commandHandler *command.CommandHandler
}

// NewCreateJuzHandler creates a new create juz handler
func NewCreateJuzHandler(commandHandler *command.CommandHandler) *CreateJuzHandler {
	return &CreateJuzHandler{
		commandHandler: commandHandler,
	}
}

// Handle processes the create juz request
// @Summary Create a new juz
// @Description Create a new juz with the provided information
// @Tags Juz
// @Accept json
// @Produce json
// @Param juz body dto.CreateJuzRequest true "Juz information"
// @Success 201 {object} JuzSuccessResponse "Juz created successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/juz [post]
func (h *CreateJuzHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.CreateJuzRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid request body", errorDetail, start)
		return
	}

	cmd := command.CreateJuzCommand{
		StartSurahID: req.StartSurahID,
		EndSurahID:   req.EndSurahID,
		StartAyahID:  req.StartAyahID,
		EndAyahID:    req.EndAyahID,
		CreatedBy:    req.CreatedBy,
	}

	result, err := h.commandHandler.CreateJuz(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to create juz", errorDetail, start)
		return
	}

	core.Success(c, http.StatusCreated, "Juz created successfully", result, start)
}
