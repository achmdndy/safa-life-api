package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetSurahByIdHandler handles the get surah by ID request
type GetSurahByIdHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetSurahByIdHandler creates a new get surah by ID handler
func NewGetSurahByIdHandler(queryHandler *query.QueryHandler) *GetSurahByIdHandler {
	return &GetSurahByIdHandler{
		queryHandler: queryHandler,
	}
}

// Handle processes the get surah by ID request
// @Summary Get surah by ID
// @Description Get a surah by its ID from the Quran
// @Tags Surahs
// @Accept json
// @Produce json
// @Param id path string true "Surah ID"
// @Param include query string false "Include related data" Enums(ayahs) example(ayahs)
// @Success 200 {object} GetSurahByIdSuccessResponse "Surah retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Surah not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/surahs/{id} [get]
func (h *GetSurahByIdHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetSurahByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid surah ID", errorDetail, start)
		return
	}

	qry := query.GetSurahByIdQuery{
		ID: req.ID,
	}

	result, err := h.queryHandler.GetSurahById(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get surah", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Surah retrieved successfully", result, start)
}
