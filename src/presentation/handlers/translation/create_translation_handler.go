package translation

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/translation/command"
	"github.com/achmdndy/safa-life-api/src/application/translation/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// CreateTranslation handles POST /translations/editions
// CreateTranslation handles POST /translations/editions
// @Summary Create Translation Edition
// @Description Add a new translation edition to the system.
// @Tags Translation Editions
// @Accept json
// @Produce json
// @Param edition body dto.CreateTranslationRequest true "Translation Edition Data"
// @Success 201 {object} core.SuccessResponse{data=dto.TranslationResponse} "Translation edition created successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid request body"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /translations/editions [post]
func (h *Handler) CreateTranslation(c *gin.Context) {
	start := time.Now()
	var req dto.CreateTranslationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromCreateTranslationRequest(req)
	result, err := h.commandHandler.CreateTranslation(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to create translation edition", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusCreated, "Translation edition created successfully", dto.ToTranslationResponse(*result), start)
}
