package translation

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/translation/command"
	"github.com/achmdndy/safa-life-api/src/application/translation/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// UpdateTranslation handles PUT /translations/editions/:id
// UpdateTranslation handles PUT /translations/editions/:id
// @Summary Update Translation Edition
// @Description Update an existing translation edition.
// @Tags Translation Editions
// @Accept json
// @Produce json
// @Param id path string true "Edition ID"
// @Param edition body dto.UpdateTranslationRequest true "Translation Edition Data"
// @Success 200 {object} core.SuccessResponse{data=dto.TranslationResponse} "Translation edition updated successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters or request body"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /translations/editions/{id} [put]
func (h *Handler) UpdateTranslation(c *gin.Context) {
	start := time.Now()
	var req dto.UpdateTranslationRequest

	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid URI parameter", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromUpdateTranslationRequest(req)
	result, err := h.commandHandler.UpdateTranslation(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to update translation edition", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Translation edition updated successfully", dto.ToTranslationResponse(*result), start)
}
