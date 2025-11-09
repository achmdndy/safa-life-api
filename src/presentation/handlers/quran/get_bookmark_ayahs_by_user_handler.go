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

// GetBookmarkAyahsByUserHandler handles listing bookmark ayahs for a user
type GetBookmarkAyahsByUserHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetBookmarkAyahsByUserHandler creates a new handler instance
func NewGetBookmarkAyahsByUserHandler(queryHandler *query.QueryHandler) *GetBookmarkAyahsByUserHandler {
	return &GetBookmarkAyahsByUserHandler{queryHandler: queryHandler}
}

// Handle processes the request
// @Summary List bookmark ayahs by user
// @Description Get a paginated list of ayah bookmarks for a user
// @Tags Bookmarks
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} BookmarkAyahListSuccessResponse "Bookmarks retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/bookmarks/ayahs/user/{userId} [get]
func (h *GetBookmarkAyahsByUserHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetBookmarkAyahsByUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid user ID", errorDetail, start)
		return
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
