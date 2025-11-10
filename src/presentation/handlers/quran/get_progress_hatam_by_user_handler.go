package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
	"github.com/safalife/core-api/src/presentation/middlewares"
)

// GetProgressHatamByUserHandler handles listing progress hatam for a user
type GetProgressHatamByUserHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetProgressHatamByUserHandler creates a new handler instance
func NewGetProgressHatamByUserHandler(queryHandler *query.QueryHandler) *GetProgressHatamByUserHandler {
	return &GetProgressHatamByUserHandler{queryHandler: queryHandler}
}

// Handle processes the request
// @Summary List progress hatam by current user
// @Description Get a paginated list of progress hatam entries for the authenticated user
// @Tags ProgressHatam
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Param include query string false "Include related data" Enums(relations) example(relations)
// @Success 200 {object} ProgressHatamListSuccessResponse "Progress hatam list retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/progress-hatam [get]
func (h *GetProgressHatamByUserHandler) Handle(c *gin.Context) {
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

	qry := query.GetProgressHatamByUserQuery{UserID: middlewares.GetUserID(c), Limit: req.Limit, Offset: req.Offset, Include: req.Include}

	result, err := h.queryHandler.GetProgressHatamByUser(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get progress hatam list", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Progress hatam list retrieved successfully", result, start)
}
