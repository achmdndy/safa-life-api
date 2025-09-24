package translation

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/translation/dto"
	"github.com/achmdndy/safa-life-api/src/application/translation/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetTranslationByID handles GET /translations/editions/:id
// GetTranslationByID handles GET /translations/editions/:id
// @Summary Get Translation Edition By ID
// @Description Retrieve a specific translation edition by its ID.
// @Tags Translation Editions
// @Accept json
// @Produce json
// @Param id path string true "Edition ID"
// @Success 200 {object} core.SuccessResponse{data=dto.TranslationResponse} "Translation edition retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid edition ID"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /translations/editions/{id} [get]
func (h *Handler) GetTranslationByID(c *gin.Context) {
	start := time.Now()
	var req dto.GetTranslationByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.ToGetTranslationByIDQuery(req.ID)
	result, err := h.queryHandler.GetTranslationByID(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get translation edition", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Translation edition retrieved successfully", result, start)
}
