package translation

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/translation/command"
	"github.com/achmdndy/safa-life-api/src/application/translation/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// CreateAyahTranslation handles POST /translations/ayahs
// CreateAyahTranslation handles POST /translations/ayahs
// @Summary Create Ayah Translation
// @Description Add a new translation for a specific ayah.
// @Tags Ayah Translations
// @Accept json
// @Produce json
// @Param ayah_translation body dto.CreateAyahTranslationRequest true "Ayah Translation Data"
// @Success 201 {object} core.SuccessResponse{data=dto.AyahTranslationResponse} "Ayah translation created successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid request body"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /translations/ayahs [post]
func (h *Handler) CreateAyahTranslation(c *gin.Context) {
	start := time.Now()
	var req dto.CreateAyahTranslationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromCreateAyahTranslationRequest(req)
	result, err := h.commandHandler.CreateAyahTranslation(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to create ayah translation", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusCreated, "Ayah translation created successfully", dto.ToAyahTranslationResponse(*result), start)
}
