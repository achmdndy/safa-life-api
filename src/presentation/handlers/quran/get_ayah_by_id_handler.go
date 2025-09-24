package quran

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/quran/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
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
// @Failure 404 {object} core.ErrorResponse "Ayah not found"
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
		// Check if the error is a validation error (invalid surah/ayah ID) or "not found" error
		if strings.Contains(err.Error(), "invalid surah ID") ||
			strings.Contains(err.Error(), "invalid ayah ID") ||
			strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Ayah not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get ayah", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result == nil {
		core.Error(c, http.StatusNotFound, "Ayah not found", &core.ErrorDetail{}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah retrieved successfully", result, start)
}
