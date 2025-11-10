package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
	"github.com/safalife/core-api/src/presentation/middlewares"
)

// GetBookmarkAyahsByUserHandler handles listing bookmark ayahs for a user
type GetBookmarkAyahsByUserHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetBookmarkAyahsByUserHandler creates a new handler instance
func NewGetBookmarkAyahsByUserHandler(queryHandler *query.QueryHandler) *GetBookmarkAyahsByUserHandler {
	return &GetBookmarkAyahsByUserHandler{queryHandler: queryHandler}
}

// Handle processes the request
// @Summary List bookmark ayahs by current user
// @Description Get a paginated list of ayah bookmarks for the authenticated user
// @Tags Bookmarks
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} BookmarkAyahListSuccessResponse "Bookmarks retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/bookmarks/ayahs [get]
func (h *GetBookmarkAyahsByUserHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	// userId diambil dari middleware; bind hanya parameter query
	var req struct {
		Limit   int    `form:"limit" binding:"omitempty,min=0,max=100"`
		Offset  int    `form:"offset" binding:"omitempty,min=0"`
		Include string `form:"include" binding:"omitempty"`
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid query parameters", errorDetail, start)
		return
	}

	qry := query.GetBookmarkAyahsByUserQuery{UserID: middlewares.GetUserID(c), Limit: req.Limit, Offset: req.Offset}

	result, err := h.queryHandler.GetBookmarkAyahsByUser(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get bookmarks", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Bookmarks retrieved successfully", result, start)
}
