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

// UpdateSurahHandler handles the update surah request
type UpdateSurahHandler struct {
	commandHandler *command.CommandHandler
}

// NewUpdateSurahHandler creates a new update surah handler
func NewUpdateSurahHandler(commandHandler *command.CommandHandler) *UpdateSurahHandler {
	return &UpdateSurahHandler{
		commandHandler: commandHandler,
	}
}

// Handle processes the update surah request
// @Summary Update a surah
// @Description Update an existing surah in the Quran
// @Tags Surahs
// @Accept json
// @Produce json
// @Param id path string true "Surah ID"
// @Param request body dto.UpdateSurahRequest true "Update surah request"
// @Success 200 {object} UpdateSurahSuccessResponse "Surah updated successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Surah not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/surahs/{id} [put]
func (h *UpdateSurahHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.UpdateSurahRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid surah ID", errorDetail, start)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid request body", errorDetail, start)
		return
	}

	cmd := command.UpdateSurahCommand{
		ID:              req.ID,
		NameArabic:      req.NameArabic,
		NameEnglish:     req.NameEnglish,
		RevelationPlace: req.RevelationPlace,
		RevelationOrder: req.RevelationOrder,
		AyahCount:       req.AyahCount,
		UpdatedBy:       middlewares.GetUserID(c),
	}

	result, err := h.commandHandler.UpdateSurah(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to update surah", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Surah updated successfully", result, start)
}
