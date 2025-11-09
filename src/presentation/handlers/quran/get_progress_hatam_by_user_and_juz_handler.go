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

// GetProgressHatamByUserAndJuzHandler handles fetching progress hatam by user and juz
type GetProgressHatamByUserAndJuzHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetProgressHatamByUserAndJuzHandler creates a new handler instance
func NewGetProgressHatamByUserAndJuzHandler(queryHandler *query.QueryHandler) *GetProgressHatamByUserAndJuzHandler {
	return &GetProgressHatamByUserAndJuzHandler{queryHandler: queryHandler}
}

// Handle processes the request
// @Summary Get progress hatam by user and juz
// @Description Get a progress hatam entry for a specific user and juz
// @Tags ProgressHatam
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param juzId path string true "Juz ID"
// @Param include query string false "Include related data" Enums(relations) example(relations)
// @Success 200 {object} ProgressHatamSuccessResponse "Progress hatam retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Progress hatam not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/progress-hatam/user/{userId}/juz/{juzId} [get]
func (h *GetProgressHatamByUserAndJuzHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetProgressHatamByUserAndJuzRequest
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

	qry := query.GetProgressHatamByUserAndJuzQuery{UserID: middlewares.GetUserID(c), JuzID: req.JuzID, Include: req.Include}

	result, err := h.queryHandler.GetProgressHatamByUserAndJuz(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get progress hatam", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Progress hatam retrieved successfully", result, start)
}
