package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetAyahTranslationByAyahAndEditionHandler handles retrieval of a single ayah translation
type GetAyahTranslationByAyahAndEditionHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetAyahTranslationByAyahAndEditionHandler creates a new handler instance
func NewGetAyahTranslationByAyahAndEditionHandler(queryHandler *query.QueryHandler) *GetAyahTranslationByAyahAndEditionHandler {
	return &GetAyahTranslationByAyahAndEditionHandler{queryHandler: queryHandler}
}

// Handle processes the get ayah translation by ayah and edition request
// @Summary Get ayah translation by ayah and edition
// @Description Get a translated ayah text by ayah ID and translation edition ID
// @Tags Translations
// @Accept json
// @Produce json
// @Param ayahId path string true "Ayah ID"
// @Param editionId path string true "Translation Edition ID"
// @Success 200 {object} GetAyahTranslationSuccessResponse "Ayah translation retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/translations/ayahs/ayah/{ayahId}/edition/{editionId} [get]
func (h *GetAyahTranslationByAyahAndEditionHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetAyahTranslationByAyahAndEditionRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid path parameters", errorDetail, start)
		return
	}

	result, err := h.queryHandler.GetAyahTranslationByAyahAndEdition(ctx, query.GetAyahTranslationByAyahAndEditionQuery{
		AyahID:    req.AyahID,
		EditionID: req.EditionID,
	})
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get ayah translation", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah translation retrieved successfully", result, start)
}
