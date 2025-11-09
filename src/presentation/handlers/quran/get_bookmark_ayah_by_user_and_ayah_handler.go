package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
	"github.com/safalife/core-api/src/presentation/middlewares"
)

// GetBookmarkAyahByUserAndAyahHandler handles fetching a bookmark by user and ayah
type GetBookmarkAyahByUserAndAyahHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetBookmarkAyahByUserAndAyahHandler creates a new handler instance
func NewGetBookmarkAyahByUserAndAyahHandler(queryHandler *query.QueryHandler) *GetBookmarkAyahByUserAndAyahHandler {
	return &GetBookmarkAyahByUserAndAyahHandler{queryHandler: queryHandler}
}

// Handle processes the request
// @Summary Get bookmark ayah by user and ayah
// @Description Get a bookmark for a specific user and ayah
// @Tags Bookmarks
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param ayahId path string true "Ayah ID"
// @Param include query string false "Include related data" Enums(ayah) example(ayah)
// @Success 200 {object} BookmarkAyahSuccessResponse "Bookmark retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Bookmark not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/bookmarks/ayahs/user/{userId}/ayah/{ayahId} [get]
func (h *GetBookmarkAyahByUserAndAyahHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetBookmarkAyahByUserAndAyahRequest
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

	qry := query.GetBookmarkAyahByUserAndAyahQuery{UserID: middlewares.GetUserID(c), AyahID: req.AyahID, Include: req.Include}

	result, err := h.queryHandler.GetBookmarkAyahByUserAndAyah(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get bookmark", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Bookmark retrieved successfully", result, start)
}
