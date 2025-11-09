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

// GetProgressHatamByUserHandler handles listing progress hatam for a user
type GetProgressHatamByUserHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetProgressHatamByUserHandler creates a new handler instance
func NewGetProgressHatamByUserHandler(queryHandler *query.QueryHandler) *GetProgressHatamByUserHandler {
	return &GetProgressHatamByUserHandler{queryHandler: queryHandler}
}

// Handle processes the request
// @Summary List progress hatam by user
// @Description Get a paginated list of progress hatam entries for a user
// @Tags ProgressHatam
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Param include query string false "Include related data" Enums(relations) example(relations)
// @Success 200 {object} ProgressHatamListSuccessResponse "Progress hatam list retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/progress-hatam/user/{userId} [get]
func (h *GetProgressHatamByUserHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetProgressHatamByUserRequest
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

	qry := query.GetProgressHatamByUserQuery{UserID: middlewares.GetUserID(c), Limit: req.Limit, Offset: req.Offset, Include: req.Include}

	result, err := h.queryHandler.GetProgressHatamByUser(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get progress hatam list", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Progress hatam list retrieved successfully", result, start)
}
