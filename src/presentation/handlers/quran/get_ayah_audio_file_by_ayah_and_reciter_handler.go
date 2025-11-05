package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetAyahAudioFileByAyahAndReciterHandler handles getting audio by ayah and reciter
type GetAyahAudioFileByAyahAndReciterHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetAyahAudioFileByAyahAndReciterHandler creates a new handler
func NewGetAyahAudioFileByAyahAndReciterHandler(queryHandler *query.QueryHandler) *GetAyahAudioFileByAyahAndReciterHandler {
	return &GetAyahAudioFileByAyahAndReciterHandler{queryHandler: queryHandler}
}

// Handle processes the get by ayah and reciter request
// @Summary Get ayah audio file by ayah and reciter
// @Description Get an ayah audio file by ayah and reciter IDs
// @Tags Audio
// @Accept json
// @Produce json
// @Param ayahId path string true "Ayah ID"
// @Param reciterId path string true "Reciter ID"
// @Success 200 {object} AyahAudioFileSuccessResponse "Ayah audio file retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/audio/ayahs/ayah/{ayahId}/reciter/{reciterId} [get]
func (h *GetAyahAudioFileByAyahAndReciterHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetAyahAudioFileByAyahAndReciterRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid URI parameters", errorDetail, start)
		return
	}

	qry := query.GetAyahAudioFileByAyahAndReciterQuery{AyahID: req.AyahID, ReciterID: req.ReciterID}
	result, err := h.queryHandler.GetAyahAudioFileByAyahAndReciter(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get ayah audio file", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah audio file retrieved successfully", result, start)
}
