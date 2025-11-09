package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetProgressHatamByIdHandler handles fetching progress hatam by ID
type GetProgressHatamByIdHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetProgressHatamByIdHandler creates a new handler instance
func NewGetProgressHatamByIdHandler(queryHandler *query.QueryHandler) *GetProgressHatamByIdHandler {
	return &GetProgressHatamByIdHandler{queryHandler: queryHandler}
}

// Handle processes the request
// @Summary Get progress hatam by ID
// @Description Get a specific progress hatam entry by its ID
// @Tags ProgressHatam
// @Accept json
// @Produce json
// @Param id path string true "ProgressHatam ID"
// @Param include query string false "Include related data" Enums(relations) example(relations)
// @Success 200 {object} ProgressHatamSuccessResponse "Progress hatam retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Progress hatam not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/progress-hatam/{id} [get]
func (h *GetProgressHatamByIdHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetProgressHatamByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid progress hatam ID", errorDetail, start)
		return
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid query parameters", errorDetail, start)
		return
	}

	qry := query.GetProgressHatamByIdQuery{ID: req.ID, Include: req.Include}

	result, err := h.queryHandler.GetProgressHatamById(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get progress hatam", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Progress hatam retrieved successfully", result, start)
}
