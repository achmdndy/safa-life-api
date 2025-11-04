package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetAyahsBySurahHandler handles the get ayahs by surah request
type GetAyahsBySurahHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetAyahsBySurahHandler creates a new get ayahs by surah handler
func NewGetAyahsBySurahHandler(queryHandler *query.QueryHandler) *GetAyahsBySurahHandler {
	return &GetAyahsBySurahHandler{
		queryHandler: queryHandler,
	}
}

// Handle processes the get ayahs by surah request
// @Summary Get ayahs by surah
// @Description Get all ayahs from a specific surah with pagination
// @Tags Ayahs
// @Accept json
// @Produce json
// @Param surahId path string true "Surah ID"
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Param include query string false "Include related data" Enums(surah) example(surah)
// @Success 200 {object} GetAyahsBySurahSuccessResponse "Ayahs retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /v1/quran/surahs/{surahId}/ayahs [get]
func (h *GetAyahsBySurahHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetAyahsBySurahIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid surah ID", errorDetail, start)
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

	qry := query.GetAyahsBySurahQuery{
		SurahID: req.SurahID,
		Limit:   req.Limit,
		Offset:  req.Offset,
	}

	result, err := h.queryHandler.GetAyahsBySurah(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get ayahs", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayahs retrieved successfully", result, start)
}
