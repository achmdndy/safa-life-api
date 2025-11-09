package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetBookmarkAyahByIdHandler handles fetching a bookmark ayah by ID
type GetBookmarkAyahByIdHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetBookmarkAyahByIdHandler creates a new handler instance
func NewGetBookmarkAyahByIdHandler(queryHandler *query.QueryHandler) *GetBookmarkAyahByIdHandler {
	return &GetBookmarkAyahByIdHandler{queryHandler: queryHandler}
}

// Handle processes the request
// @Summary Get bookmark ayah by ID
// @Description Get a specific bookmark by its ID
// @Tags Bookmarks
// @Accept json
// @Produce json
// @Param id path string true "Bookmark ID"
// @Param include query string false "Include related data" Enums(ayah) example(ayah)
// @Success 200 {object} BookmarkAyahSuccessResponse "Bookmark retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Bookmark not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/bookmarks/ayahs/{id} [get]
func (h *GetBookmarkAyahByIdHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetBookmarkAyahByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid bookmark ID", errorDetail, start)
		return
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid query parameters", errorDetail, start)
		return
	}

	qry := query.GetBookmarkAyahByIdQuery{ID: req.ID, Include: req.Include}

	result, err := h.queryHandler.GetBookmarkAyahById(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get bookmark", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Bookmark retrieved successfully", result, start)
}
