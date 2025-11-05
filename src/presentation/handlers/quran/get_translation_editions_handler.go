package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetTranslationEditionsHandler handles retrieval of translation editions
type GetTranslationEditionsHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetTranslationEditionsHandler creates a new handler instance
func NewGetTranslationEditionsHandler(queryHandler *query.QueryHandler) *GetTranslationEditionsHandler {
	return &GetTranslationEditionsHandler{queryHandler: queryHandler}
}

// Handle processes the get translation editions request
// @Summary Get translation editions
// @Description Get translation editions with optional language filter and pagination
// @Tags Translations
// @Accept json
// @Produce json
// @Param language query string false "Language code (e.g., en, id)"
// @Param limit query int false "Limit (0 for all)" default(0)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} GetTranslationEditionListSuccessResponse "Translation editions retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/translations/editions [get]
func (h *GetTranslationEditionsHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetTranslationEditionsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid query parameters", errorDetail, start)
		return
	}

	result, err := h.queryHandler.GetTranslationEditions(ctx, query.GetTranslationEditionsQuery{
		Language: req.Language,
		Limit:    req.Limit,
		Offset:   req.Offset,
	})
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get translation editions", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Translation editions retrieved successfully", result, start)
}
