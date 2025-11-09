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

// CreateLastReadHandler handles creating last read entries
type CreateLastReadHandler struct {
	commandHandler *command.CommandHandler
}

// NewCreateLastReadHandler constructs the handler
func NewCreateLastReadHandler(commandHandler *command.CommandHandler) *CreateLastReadHandler {
	return &CreateLastReadHandler{commandHandler: commandHandler}
}

// Handle processes the request
// @Summary Create last read
// @Description Create a new last read entry for a user
// @Tags LastRead
// @Accept json
// @Produce json
// @Param lastRead body dto.CreateLastReadRequest true "LastRead information"
// @Success 201 {object} LastReadSuccessResponse "Last read created successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/last-reads [post]
func (h *CreateLastReadHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.CreateLastReadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid request body", errorDetail, start)
		return
	}

	cmd := command.CreateLastReadCommand{
		UserID:      middlewares.GetUserID(c),
		SurahID:     req.SurahID,
		AyahID:      req.AyahID,
		AyahNumber:  req.AyahNumber,
		ProgressPct: req.ProgressPct,
		CreatedBy:   middlewares.GetUserID(c),
	}

	result, err := h.commandHandler.CreateLastRead(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to create last read", errorDetail, start)
		return
	}

	core.Success(c, http.StatusCreated, "Last read created successfully", result, start)
}
