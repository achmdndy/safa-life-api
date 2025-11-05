package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
)

// CreateReciterHandler handles the create reciter request
type CreateReciterHandler struct {
	commandHandler *command.CommandHandler
}

// NewCreateReciterHandler creates a new create reciter handler
func NewCreateReciterHandler(commandHandler *command.CommandHandler) *CreateReciterHandler {
	return &CreateReciterHandler{
		commandHandler: commandHandler,
	}
}

// Handle processes the create reciter request
// @Summary Create a new reciter
// @Description Create a new reciter
// @Tags Reciters
// @Accept json
// @Produce json
// @Param request body dto.CreateReciterRequest true "Create reciter request"
// @Success 201 {object} CreateReciterSuccessResponse "Reciter created successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/reciters [post]
func (h *CreateReciterHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.CreateReciterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid request body", errorDetail, start)
		return
	}

	cmd := command.CreateReciterCommand{
		Name:      req.Name,
		Style:     req.Style,
		CreatedBy: req.CreatedBy,
	}

	result, err := h.commandHandler.CreateReciter(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to create reciter", errorDetail, start)
		return
	}

	core.Success(c, http.StatusCreated, "Reciter created successfully", result, start)
}
