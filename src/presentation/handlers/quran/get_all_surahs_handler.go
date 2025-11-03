package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetAllSurahsHandler handles the get all surahs request
type GetAllSurahsHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetAllSurahsHandler creates a new get all surahs handler
func NewGetAllSurahsHandler(queryHandler *query.QueryHandler) *GetAllSurahsHandler {
	return &GetAllSurahsHandler{
		queryHandler: queryHandler,
	}
}

// Handle processes the get all surahs request
// @Summary Get all surahs
// @Description Get all surahs from the Quran with pagination
// @Tags Surahs
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Param include query string false "Include related data" Enums(ayahs) example(ayahs)
// @Success 200 {object} GetAllSurahsSuccessResponse "Surahs retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /api/v1/quran/surahs [get]
func (h *GetAllSurahsHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetAllSurahsRequest
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

	qry := query.GetAllSurahsQuery{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	result, err := h.queryHandler.GetAllSurahs(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get surahs", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Surahs retrieved successfully", result, start)
}
