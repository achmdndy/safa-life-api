package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetAllJuzHandler handles the get all juz request
type GetAllJuzHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetAllJuzHandler creates a new get all juz handler
func NewGetAllJuzHandler(queryHandler *query.QueryHandler) *GetAllJuzHandler {
	return &GetAllJuzHandler{
		queryHandler: queryHandler,
	}
}

// Handle processes the get all juz request
// @Summary Get all juz
// @Description Get all juz with pagination
// @Tags Juz
// @Accept json
// @Produce json
// @Param limit query int false "Limit (0 for all)" default(0)
// @Param offset query int false "Offset" default(0)
// @Param include query string false "Include related data" Enums(relations) example(relations)
// @Success 200 {object} JuzListSuccessResponse "Juz list retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /v1/quran/juz [get]
func (h *GetAllJuzHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetAllJuzRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid query parameters", errorDetail, start)
		return
	}

	qry := query.GetAllJuzQuery{
		Limit:   req.Limit,
		Offset:  req.Offset,
		Include: req.Include,
	}

	result, err := h.queryHandler.GetAllJuz(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get juz list", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Juz list retrieved successfully", result, start)
}
