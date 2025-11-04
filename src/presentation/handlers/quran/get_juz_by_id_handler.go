package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetJuzByIdHandler handles the get juz by id request
type GetJuzByIdHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetJuzByIdHandler creates a new get juz by id handler
func NewGetJuzByIdHandler(queryHandler *query.QueryHandler) *GetJuzByIdHandler {
	return &GetJuzByIdHandler{
		queryHandler: queryHandler,
	}
}

// Handle processes the get juz by id request
// @Summary Get juz by ID
// @Description Get a specific juz by its ID
// @Tags Juz
// @Accept json
// @Produce json
// @Param id path string true "Juz ID"
// @Param include query string false "Include related data" Enums(relations) example(relations)
// @Success 200 {object} JuzSuccessResponse "Juz retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Juz not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /v1/quran/juz/{id} [get]
func (h *GetJuzByIdHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetJuzByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid juz ID", errorDetail, start)
		return
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid query parameters", errorDetail, start)
		return
	}

	qry := query.GetJuzByIdQuery{
		ID:      req.ID,
		Include: req.Include,
	}

	result, err := h.queryHandler.GetJuzById(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get juz", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Juz retrieved successfully", result, start)
}
