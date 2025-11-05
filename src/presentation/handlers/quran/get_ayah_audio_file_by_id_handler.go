package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetAyahAudioFileByIdHandler handles the get ayah audio file by ID request
type GetAyahAudioFileByIdHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetAyahAudioFileByIdHandler creates a new handler
func NewGetAyahAudioFileByIdHandler(queryHandler *query.QueryHandler) *GetAyahAudioFileByIdHandler {
	return &GetAyahAudioFileByIdHandler{queryHandler: queryHandler}
}

// Handle processes the get ayah audio file by ID request
// @Summary Get ayah audio file by ID
// @Description Get an ayah audio file by ID
// @Tags Audio
// @Accept json
// @Produce json
// @Param id path string true "Ayah audio file ID"
// @Success 200 {object} AyahAudioFileSuccessResponse "Ayah audio file retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/audio/ayahs/{id} [get]
func (h *GetAyahAudioFileByIdHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetAyahAudioFileByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid URI parameters", errorDetail, start)
		return
	}

	qry := query.GetAyahAudioFileByIdQuery{ID: req.ID}
	result, err := h.queryHandler.GetAyahAudioFileById(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get ayah audio file", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah audio file retrieved successfully", result, start)
}
