package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
	"github.com/safalife/core-api/src/presentation/middlewares"
)

// GetLastReadByUserAndSurahHandler handles fetching a last read by user and surah
type GetLastReadByUserAndSurahHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetLastReadByUserAndSurahHandler creates a new handler instance
func NewGetLastReadByUserAndSurahHandler(queryHandler *query.QueryHandler) *GetLastReadByUserAndSurahHandler {
	return &GetLastReadByUserAndSurahHandler{queryHandler: queryHandler}
}

// Handle processes the request
// @Summary Get last read by current user and surah
// @Description Get the last read entry for the authenticated user and surah
// @Tags LastRead
// @Accept json
// @Produce json
// @Param surahId path string true "Surah ID"
// @Param include query string false "Include related data" Enums(relations) example(relations)
// @Success 200 {object} LastReadSuccessResponse "Last read retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Last read not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/last-reads/surah/{surahId} [get]
func (h *GetLastReadByUserAndSurahHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	// Bind hanya surahId dari URI, userId diambil dari middleware
	var req struct {
		SurahID string `uri:"surahId" binding:"required"`
		Include string `form:"include" binding:"omitempty"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid path parameters", errorDetail, start)
		return
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid query parameters", errorDetail, start)
		return
	}

	qry := query.GetLastReadByUserAndSurahQuery{UserID: middlewares.GetUserID(c), SurahID: req.SurahID, Include: req.Include}

	result, err := h.queryHandler.GetLastReadByUserAndSurah(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get last read", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Last read retrieved successfully", result, start)
}
