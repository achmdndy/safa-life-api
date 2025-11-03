package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetJuzByNumberHandler handles the get juz by number request
type GetJuzByNumberHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetJuzByNumberHandler creates a new get juz by number handler
func NewGetJuzByNumberHandler(queryHandler *query.QueryHandler) *GetJuzByNumberHandler {
	return &GetJuzByNumberHandler{
		queryHandler: queryHandler,
	}
}

// Handle processes the get juz by number request
// @Summary Get juz by number
// @Description Get a specific juz by its number
// @Tags Juz
// @Accept json
// @Produce json
// @Param number path int true "Juz Number"
// @Param include query string false "Include related data" Enums(relations) example(relations)
// @Success 200 {object} JuzSuccessResponse "Juz retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Juz not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /api/v1/quran/juz/number/{number} [get]
func (h *GetJuzByNumberHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetJuzByNumberRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid juz number", errorDetail, start)
		return
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid query parameters", errorDetail, start)
		return
	}

	qry := query.GetJuzByNumberQuery{
		Number:  req.Number,
		Include: req.Include,
	}

	result, err := h.queryHandler.GetJuzByNumber(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get juz", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Juz retrieved successfully", result, start)
}