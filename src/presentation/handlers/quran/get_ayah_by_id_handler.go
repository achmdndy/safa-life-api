package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
)

// GetAyahByIdHandler handles the get ayah by ID request
type GetAyahByIdHandler struct {
	queryHandler *query.QueryHandler
}

// NewGetAyahByIdHandler creates a new get ayah by ID handler
func NewGetAyahByIdHandler(queryHandler *query.QueryHandler) *GetAyahByIdHandler {
	return &GetAyahByIdHandler{
		queryHandler: queryHandler,
	}
}

// Handle processes the get ayah by ID request
// @Summary Get ayah by ID
// @Description Get an ayah by its ID from the Quran
// @Tags Ayahs
// @Accept json
// @Produce json
// @Param id path string true "Ayah ID"
// @Param include query string false "Include related data" Enums(surah) example(surah)
// @Success 200 {object} GetAyahByIdSuccessResponse "Ayah retrieved successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Ayah not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /api/v1/quran/ayahs/{id} [get]
func (h *GetAyahByIdHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetAyahByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid ayah ID", errorDetail, start)
		return
	}

	qry := query.GetAyahByIdQuery{
		ID: req.ID,
	}

	result, err := h.queryHandler.GetAyahById(ctx, qry)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get ayah", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah retrieved successfully", result, start)
}
