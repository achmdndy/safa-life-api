package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetAyahAudioFilesBySurahAndReciterHandler handles listing audio files by surah and reciter
type GetAyahAudioFilesBySurahAndReciterHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetAyahAudioFilesBySurahAndReciterHandler creates a new handler
func NewGetAyahAudioFilesBySurahAndReciterHandler(queryHandler *query.QueryHandler) *GetAyahAudioFilesBySurahAndReciterHandler {
	return &GetAyahAudioFilesBySurahAndReciterHandler{queryHandler: queryHandler}
}

// Handle processes the list by surah and reciter request
// @Summary List ayah audio files by surah and reciter
// @Description List ayah audio files by surah and reciter IDs with pagination
// @Tags Audio
// @Accept json
// @Produce json
// @Param surahId path string true "Surah ID"
// @Param reciterId path string true "Reciter ID"
// @Param limit query int false "Limit (0 for all)" default(0)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} AyahAudioFileListSuccessResponse "Ayah audio files retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/audio/ayahs/surah/{surahId}/reciter/{reciterId} [get]
func (h *GetAyahAudioFilesBySurahAndReciterHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var uri dto.GetAyahAudioFilesBySurahAndReciterRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid URI parameters", errorDetail, start)
		return
	}

	var queryParams dto.GetAyahAudioFilesBySurahAndReciterRequest
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid query parameters", errorDetail, start)
		return
	}

	qry := query.GetAyahAudioFilesBySurahAndReciterQuery{
		SurahID:   uri.SurahID,
		ReciterID: uri.ReciterID,
		Limit:     queryParams.Limit,
		Offset:    queryParams.Offset,
	}

	result, err := h.queryHandler.GetAyahAudioFilesBySurahAndReciter(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get ayah audio files", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah audio files retrieved successfully", result, start)
}
