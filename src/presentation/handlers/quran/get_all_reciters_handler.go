package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetAllRecitersHandler handles the get all reciters request
type GetAllRecitersHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetAllRecitersHandler creates a new get all reciters handler
func NewGetAllRecitersHandler(queryHandler *query.QueryHandler) *GetAllRecitersHandler {
	return &GetAllRecitersHandler{queryHandler: queryHandler}
}

// Handle processes the get all reciters request
// @Summary Get all reciters
// @Description Get all reciters with pagination
// @Tags Reciters
// @Accept json
// @Produce json
// @Param limit query int false "Limit (0 for all)" default(0)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} ReciterListSuccessResponse "Reciters retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/reciters [get]
func (h *GetAllRecitersHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetAllRecitersRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid query parameters", errorDetail, start)
		return
	}

	qry := query.GetAllRecitersQuery{Limit: req.Limit, Offset: req.Offset}
	result, err := h.queryHandler.GetAllReciters(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get reciters", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Reciters retrieved successfully", result, start)
}
