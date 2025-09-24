package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/achmdndy/safa-life-api/src/application/quran/command"
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

// CreateSurah handles POST /surahs
// @Summary Create a new surah
// @Description Create a new surah (chapter) in the Quran
// @Tags Surah
// @Accept json
// @Produce json
// @Param surah body dto.CreateSurahRequest true "Surah data"
// @Success 201 {object} core.SuccessResponse{data=dto.SurahResponse} "Surah created successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid request body"
// @Failure 500 {object} core.ErrorResponse "Failed to create surah"
// @Router /surahs [post]
func (h *Handler) CreateSurah(c *gin.Context) {
	start := time.Now()
	
	var req dto.CreateSurahRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromCreateSurahRequest(req)
	domainResult, err := h.commandHandler.CreateSurah(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to create surah", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if domainResult != nil {
		result := dto.ToSurahResponse(*domainResult)
		core.Success(c, http.StatusCreated, "Surah created successfully", result, start)
	} else {
		core.Success(c, http.StatusCreated, "Surah created successfully", core.EmptyData{}, start)
	}
}