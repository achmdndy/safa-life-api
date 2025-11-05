package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetAyahTranslationsBySurahAndEditionHandler handles retrieval of ayah translations by surah and edition
type GetAyahTranslationsBySurahAndEditionHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetAyahTranslationsBySurahAndEditionHandler creates a new handler instance
func NewGetAyahTranslationsBySurahAndEditionHandler(queryHandler *query.QueryHandler) *GetAyahTranslationsBySurahAndEditionHandler {
	return &GetAyahTranslationsBySurahAndEditionHandler{queryHandler: queryHandler}
}

// Handle processes the get ayah translations by surah and edition request
// @Summary Get ayah translations by surah and edition
// @Description Get translated ayahs for a surah by translation edition with pagination
// @Tags Translations
// @Accept json
// @Produce json
// @Param surahId path string true "Surah ID"
// @Param editionId path string true "Translation Edition ID"
// @Param limit query int false "Limit (0 for all)" default(0)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} GetAyahTranslationListSuccessResponse "Ayah translations retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/translations/ayahs/surah/{surahId}/edition/{editionId} [get]
func (h *GetAyahTranslationsBySurahAndEditionHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	// Bind URI params first
	var req dto.GetAyahTranslationsBySurahAndEditionRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid path parameters", errorDetail, start)
		return
	}

	// Bind query params for pagination
	if err := c.ShouldBindQuery(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid query parameters", errorDetail, start)
		return
	}

	result, err := h.queryHandler.GetAyahTranslationsBySurahAndEdition(ctx, query.GetAyahTranslationsBySurahAndEditionQuery{
		SurahID:   req.SurahID,
		EditionID: req.EditionID,
		Limit:     req.Limit,
		Offset:    req.Offset,
	})
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get ayah translations", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah translations retrieved successfully", result, start)
}
