package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
	"github.com/safalife/core-api/src/presentation/middlewares"
)

// GetLastReadsByUserHandler handles listing last reads for a user
type GetLastReadsByUserHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetLastReadsByUserHandler creates a new handler instance
func NewGetLastReadsByUserHandler(queryHandler *query.QueryHandler) *GetLastReadsByUserHandler {
	return &GetLastReadsByUserHandler{queryHandler: queryHandler}
}

// Handle processes the request
// @Summary List last reads by current user
// @Description Get a paginated list of last read entries for the authenticated user
// @Tags LastRead
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Param include query string false "Include related data" Enums(relations) example(relations)
// @Success 200 {object} LastReadListSuccessResponse "Last reads retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/last-reads [get]
func (h *GetLastReadsByUserHandler) Handle(c *gin.Context) {
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

	qry := query.GetLastReadsByUserQuery{UserID: middlewares.GetUserID(c), Limit: req.Limit, Offset: req.Offset, Include: req.Include}

	result, err := h.queryHandler.GetLastReadsByUser(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get last reads", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Last reads retrieved successfully", result, start)
}
