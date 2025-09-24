package quran

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/quran/command"
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// UpdateAyah handles PUT /ayahs/:surahId/:ayahId
// @Summary Update an ayah
// @Description Update an existing ayah (verse) in the Quran
// @Tags Ayah
// @Accept json
// @Produce json
// @Param surahId path int true "Surah ID"
// @Param ayahId path int true "Ayah ID"
// @Param ayah body dto.UpdateAyahRequest true "Updated ayah data"
// @Success 200 {object} core.SuccessResponse{data=dto.AyahResponse} "Ayah updated successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid surah ID, ayah ID, or request body"
// @Failure 404 {object} core.ErrorResponse "Ayah not found"
// @Failure 500 {object} core.ErrorResponse "Failed to update ayah"
// @Router /ayahs/{surahId}/{ayahId} [put]
func (h *Handler) UpdateAyah(c *gin.Context) {
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

	var req dto.UpdateAyahRequest
	if bindErr := c.ShouldBindJSON(&req); bindErr != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: bindErr.Error()}, start)
		return
	}

	cmd := command.FromUpdateAyahRequest(req)
	cmd.SurahID = surahID
	cmd.AyahID = ayahID
	domainResult, err := h.commandHandler.UpdateAyah(c.Request.Context(), cmd)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Ayah not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to update ayah", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if domainResult == nil {
		core.Error(c, http.StatusNotFound, "Ayah not found", &core.ErrorDetail{}, start)
		return
	}

	result := dto.ToAyahResponse(*domainResult)
	core.Success(c, http.StatusOK, "Ayah updated successfully", result, start)
}
