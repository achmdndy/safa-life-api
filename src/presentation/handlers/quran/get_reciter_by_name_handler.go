package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetReciterByNameHandler handles the get reciter by name request
type GetReciterByNameHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetReciterByNameHandler creates a new get reciter by name handler
func NewGetReciterByNameHandler(queryHandler *query.QueryHandler) *GetReciterByNameHandler {
	return &GetReciterByNameHandler{queryHandler: queryHandler}
}

// Handle processes the get reciter by name request
// @Summary Get reciter by name
// @Description Get a reciter by name
// @Tags Reciters
// @Accept json
// @Produce json
// @Param name path string true "Reciter name"
// @Success 200 {object} ReciterSuccessResponse "Reciter retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/reciters/name/{name} [get]
func (h *GetReciterByNameHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetReciterByNameRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid URI parameters", errorDetail, start)
		return
	}

	qry := query.GetReciterByNameQuery{Name: req.Name}
	result, err := h.queryHandler.GetReciterByName(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get reciter", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Reciter retrieved successfully", result, start)
}
