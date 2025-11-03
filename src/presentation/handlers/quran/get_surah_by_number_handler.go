package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetSurahByNumberHandler handles the get surah by number request
type GetSurahByNumberHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetSurahByNumberHandler creates a new get surah by number handler
func NewGetSurahByNumberHandler(queryHandler *query.QueryHandler) *GetSurahByNumberHandler {
	return &GetSurahByNumberHandler{
		queryHandler: queryHandler,
	}
}

// Handle processes the get surah by number request
// @Summary Get surah by number
// @Description Get a surah by its number from the Quran
// @Tags Surah
// @Accept json
// @Produce json
// @Param number path int true "Surah Number"
// @Param include query string false "Include related data" Enums(ayahs) example(ayahs)
// @Success 200 {object} GetSurahByNumberSuccessResponse "Surah retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Surah not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /api/v1/quran/surahs/number/{number} [get]
func (h *GetSurahByNumberHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetSurahByNumberRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid surah number", errorDetail, start)
		return
	}

	qry := query.GetSurahByNumberQuery{
		Number: req.Number,
	}

	result, err := h.queryHandler.GetSurahByNumber(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get surah", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Surah retrieved successfully", result, start)
}