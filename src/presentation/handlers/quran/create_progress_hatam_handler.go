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

// CreateProgressHatamHandler handles creating progress hatam entries
type CreateProgressHatamHandler struct {
	commandHandler *command.CommandHandler
}

// NewCreateProgressHatamHandler constructs the handler
func NewCreateProgressHatamHandler(commandHandler *command.CommandHandler) *CreateProgressHatamHandler {
	return &CreateProgressHatamHandler{commandHandler: commandHandler}
}

// Handle processes the request
// @Summary Create progress hatam
// @Description Create a new progress hatam entry for a user
// @Tags ProgressHatam
// @Accept json
// @Produce json
// @Param progressHatam body dto.CreateProgressHatamRequest true "ProgressHatam information"
// @Success 201 {object} ProgressHatamSuccessResponse "Progress hatam created successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/progress-hatam [post]
func (h *CreateProgressHatamHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.CreateProgressHatamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid request body", errorDetail, start)
		return
	}

	cmd := command.CreateProgressHatamCommand{
		UserID:      middlewares.GetUserID(c),
		JuzID:       req.JuzID,
		StartAyahID: req.StartAyahID,
		ProgressPct: req.ProgressPct,
		CreatedBy:   middlewares.GetUserID(c),
	}

	result, err := h.commandHandler.CreateProgressHatam(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to create progress hatam", errorDetail, start)
		return
	}

	core.Success(c, http.StatusCreated, "Progress hatam created successfully", result, start)
}
