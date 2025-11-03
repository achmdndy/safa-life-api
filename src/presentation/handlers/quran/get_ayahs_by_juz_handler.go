package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetAyahsByJuzHandler handles the get ayahs by juz request
type GetAyahsByJuzHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetAyahsByJuzHandler creates a new get ayahs by juz handler
func NewGetAyahsByJuzHandler(queryHandler *query.QueryHandler) *GetAyahsByJuzHandler {
	return &GetAyahsByJuzHandler{
		queryHandler: queryHandler,
	}
}

// Handle processes the get ayahs by juz request
// @Summary Get ayahs by juz
// @Description Get all ayahs from a specific juz with pagination
// @Tags Ayahs
// @Accept json
// @Produce json
// @Param juzNumber path int true "Juz Number"
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} GetAyahsByJuzSuccessResponse "Ayahs retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /api/v1/quran/juz/{juzNumber}/ayahs [get]
func (h *GetAyahsByJuzHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetAyahsByJuzNumberRequest
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

	// Set default values if not provided
	if req.Limit == 0 {
		req.Limit = 10
	}

	qry := query.GetAyahsByJuzQuery{
		JuzNumber: req.JuzNumber,
		Limit:     req.Limit,
		Offset:    req.Offset,
	}

	result, err := h.queryHandler.GetAyahsByJuz(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get ayahs", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayahs retrieved successfully", result, start)
}
