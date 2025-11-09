package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetLastReadByIdHandler handles fetching a last read by ID
type GetLastReadByIdHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetLastReadByIdHandler creates a new handler instance
func NewGetLastReadByIdHandler(queryHandler *query.QueryHandler) *GetLastReadByIdHandler {
	return &GetLastReadByIdHandler{queryHandler: queryHandler}
}

// Handle processes the request
// @Summary Get last read by ID
// @Description Get a specific last read entry by its ID
// @Tags LastRead
// @Accept json
// @Produce json
// @Param id path string true "LastRead ID"
// @Param include query string false "Include related data" Enums(relations) example(relations)
// @Success 200 {object} LastReadSuccessResponse "Last read retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Last read not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/last-reads/{id} [get]
func (h *GetLastReadByIdHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetLastReadByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid last read ID", errorDetail, start)
		return
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid query parameters", errorDetail, start)
		return
	}

	qry := query.GetLastReadByIdQuery{ID: req.ID, Include: req.Include}

	result, err := h.queryHandler.GetLastReadById(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get last read", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Last read retrieved successfully", result, start)
}
