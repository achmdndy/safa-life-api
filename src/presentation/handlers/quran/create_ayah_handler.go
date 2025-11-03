package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
)

// CreateAyahHandler handles the create ayah request
type CreateAyahHandler struct {
	commandHandler *command.CommandHandler
}

// NewCreateAyahHandler creates a new create ayah handler
func NewCreateAyahHandler(commandHandler *command.CommandHandler) *CreateAyahHandler {
	return &CreateAyahHandler{
		commandHandler: commandHandler,
	}
}

// Handle processes the create ayah request
// @Summary Create a new ayah
// @Description Create a new ayah in the Quran
// @Tags Ayahs
// @Accept json
// @Produce json
// @Param request body dto.CreateAyahRequest true "Create ayah request"
// @Success 201 {object} CreateAyahSuccessResponse "Ayah created successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /api/v1/quran/ayahs [post]
func (h *CreateAyahHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.CreateAyahRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid request body", errorDetail, start)
		return
	}

	cmd := command.CreateAyahCommand{
		SurahID:      req.SurahID,
		Text:         req.Text,
		PageNumber:   req.PageNumber,
		JuzNumber:    req.JuzNumber,
		HizbNumber:   req.HizbNumber,
		ManzilNumber: req.ManzilNumber,
		CreatedBy:    req.CreatedBy,
	}

	result, err := h.commandHandler.CreateAyah(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to create ayah", errorDetail, start)
		return
	}

	core.Success(c, http.StatusCreated, "Ayah created successfully", result, start)
}
