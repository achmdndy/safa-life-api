package translation

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/translation/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetAllTranslations handles GET /translations/editions
// GetAllTranslations handles GET /translations/editions
// @Summary Get All Translation Editions
// @Description Retrieve a list of all available translation editions.
// @Tags Translation Editions
// @Accept json
// @Produce json
// @Success 200 {object} core.SuccessResponse{data=[]dto.TranslationResponse} "Translation editions retrieved successfully"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /translations/editions [get]
func (h *Handler) GetAllTranslations(c *gin.Context) {
	start := time.Now()
	queryReq := query.ToGetAllTranslationsQuery()
	result, err := h.queryHandler.GetAllTranslations(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get translation editions", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Translation editions retrieved successfully", result, start)
}
