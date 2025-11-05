package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetReciterByIdHandler handles the get reciter by ID request
type GetReciterByIdHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetReciterByIdHandler creates a new get reciter by ID handler
func NewGetReciterByIdHandler(queryHandler *query.QueryHandler) *GetReciterByIdHandler {
	return &GetReciterByIdHandler{queryHandler: queryHandler}
}

// Handle processes the get reciter by ID request
// @Summary Get reciter by ID
// @Description Get a reciter by ID
// @Tags Reciters
// @Accept json
// @Produce json
// @Param id path string true "Reciter ID"
// @Success 200 {object} ReciterSuccessResponse "Reciter retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/reciters/{id} [get]
func (h *GetReciterByIdHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetReciterByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid URI parameters", errorDetail, start)
		return
	}

	qry := query.GetReciterByIdQuery{ID: req.ID}
	result, err := h.queryHandler.GetReciterById(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get reciter", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Reciter retrieved successfully", result, start)
}
