package tajweed

import (
	"net/http"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tajweed/command"
	"github.com/achmdndy/safa-life-api/src/application/tajweed/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// UpdateAyahTajweed handles PUT /tajweed/ayah/:tajweed_id/:surah_id/:ayah_id
// UpdateAyahTajweed handles PUT /tajweed/ayah/:tajweed_id/:surah_id/:ayah_id
// @Summary Update Ayah Tajweed
// @Description Update the set of tajweed rules for a specific ayah.
// @Tags Tajweed
// @Accept json
// @Produce json
// @Param tajweed_id path string true "Tajweed ID"
// @Param surah_id path int true "Surah ID"
// @Param ayah_id path int true "Ayah ID"
// @Param ayah_tajweed body dto.UpdateAyahTajweedRequest true "Ayah Tajweed Data"
// @Success 200 {object} core.SuccessResponse{data=dto.AyahTajweedResponse} "Ayah tajweed updated successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters or request body"
// @Failure 404 {object} core.ErrorResponse "Ayah tajweed not found"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /tajweed/ayah/{tajweed_id}/{surah_id}/{ayah_id} [put]
func (h *Handler) UpdateAyahTajweed(c *gin.Context) {
	start := time.Now()

	var uriReq dto.UpdateAyahTajweedRequest
	if err := c.ShouldBindUri(&uriReq); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid URI parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	var jsonReq dto.UpdateAyahTajweedRequest
	if err := c.ShouldBindJSON(&jsonReq); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	uriReq.Words = jsonReq.Words

	cmd := command.FromUpdateAyahTajweedRequest(uriReq)
	domainResult, err := h.commandHandler.UpdateAyahTajweed(c.Request.Context(), cmd)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Ayah tajweed not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to update ayah tajweed", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if domainResult == nil {
		core.Error(c, http.StatusNotFound, "Ayah tajweed not found", &core.ErrorDetail{}, start)
		return
	}

	result := dto.ToAyahTajweedResponse(*domainResult)
	core.Success(c, http.StatusOK, "Ayah tajweed updated successfully", result, start)
}
