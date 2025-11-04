package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
)

// CreateSurahHandler handles the create surah request
type CreateSurahHandler struct {
	commandHandler *command.CommandHandler
}

// NewCreateSurahHandler creates a new create surah handler
func NewCreateSurahHandler(commandHandler *command.CommandHandler) *CreateSurahHandler {
	return &CreateSurahHandler{
		commandHandler: commandHandler,
	}
}

// Handle processes the create surah request
// @Summary Create a new surah
// @Description Create a new surah in the Quran
// @Tags Surahs
// @Accept json
// @Produce json
// @Param request body dto.CreateSurahRequest true "Create surah request"
// @Success 201 {object} CreateSurahSuccessResponse "Surah created successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /v1/quran/surahs [post]
func (h *CreateSurahHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.CreateSurahRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid request body", errorDetail, start)
		return
	}

	cmd := command.CreateSurahCommand{
		NameArabic:      req.NameArabic,
		NameEnglish:     req.NameEnglish,
		RevelationPlace: req.RevelationPlace,
		RevelationOrder: req.RevelationOrder,
		AyahCount:       req.AyahCount,
		CreatedBy:       req.CreatedBy,
	}

	result, err := h.commandHandler.CreateSurah(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to create surah", errorDetail, start)
		return
	}

	core.Success(c, http.StatusCreated, "Surah created successfully", result, start)
}
