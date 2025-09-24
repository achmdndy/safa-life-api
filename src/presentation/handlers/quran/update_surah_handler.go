package quran

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/achmdndy/safa-life-api/src/application/quran/command"
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

// UpdateSurah handles PUT /surahs/:id
// @Summary Update an existing surah
// @Description Update an existing surah (chapter) in the Quran
// @Tags Surah
// @Accept json
// @Produce json
// @Param id path int true "Surah ID"
// @Param surah body dto.UpdateSurahRequest true "Surah data"
// @Success 200 {object} core.SuccessResponse{data=dto.SurahResponse} "Surah updated successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid surah ID or request body"
// @Failure 500 {object} core.ErrorResponse "Failed to update surah"
// @Router /surahs/{id} [put]
func (h *Handler) UpdateSurah(c *gin.Context) {
	start := time.Now()
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid surah ID", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	var req dto.UpdateSurahRequest
	if bindErr := c.ShouldBindJSON(&req); bindErr != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: bindErr.Error()}, start)
		return
	}

	cmd := command.FromUpdateSurahRequest(req)
	cmd.ID = id
	domainResult, err := h.commandHandler.UpdateSurah(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to update surah", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if domainResult != nil {
		result := dto.ToSurahResponse(*domainResult)
		core.Success(c, http.StatusOK, "Surah updated successfully", result, start)
	} else {
		core.Success(c, http.StatusOK, "Surah updated successfully", core.EmptyData{}, start)
	}
}