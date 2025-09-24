package quran

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/achmdndy/safa-life-api/src/application/quran/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

// GetAyahByID handles GET /ayahs/:surahId/:ayahId
// @Summary Get ayah by ID
// @Description Get a specific ayah by surah ID and ayah ID
// @Tags Ayah
// @Accept json
// @Produce json
// @Param surahId path int true "Surah ID"
// @Param ayahId path int true "Ayah ID"
// @Success 200 {object} core.SuccessResponse{data=dto.AyahResponse} "Ayah retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid surah ID or ayah ID"
// @Failure 500 {object} core.ErrorResponse "Failed to get ayah"
// @Router /ayahs/{surahId}/{ayahId} [get]
func (h *Handler) GetAyahByID(c *gin.Context) {
	start := time.Now()
	
	surahID, err := strconv.Atoi(c.Param("surahId"))
	if err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid surah ID", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	ayahID, err := strconv.Atoi(c.Param("ayahId"))
	if err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid ayah ID", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.GetAyahByIDQuery{SurahID: surahID, AyahID: ayahID}
	result, err := h.queryHandler.GetAyahByID(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get ayah", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result != nil {
		core.Success(c, http.StatusOK, "Ayah retrieved successfully", result, start)
	} else {
		core.Success(c, http.StatusOK, "Ayah retrieved successfully", core.EmptyData{}, start)
	}
}