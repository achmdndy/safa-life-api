package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/achmdndy/safa-life-api/src/application/quran/command"
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

// CreateAyah handles POST /ayahs
// @Summary Create a new ayah
// @Description Create a new ayah (verse) in the Quran
// @Tags Ayah
// @Accept json
// @Produce json
// @Param ayah body dto.CreateAyahRequest true "Ayah data"
// @Success 201 {object} core.SuccessResponse{data=dto.AyahResponse} "Ayah created successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid request body"
// @Failure 500 {object} core.ErrorResponse "Failed to create ayah"
// @Router /ayahs [post]
func (h *Handler) CreateAyah(c *gin.Context) {
	start := time.Now()
	
	var req dto.CreateAyahRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromCreateAyahRequest(req)
	domainResult, err := h.commandHandler.CreateAyah(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to create ayah", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if domainResult != nil {
		result := dto.ToAyahResponse(*domainResult)
		core.Success(c, http.StatusCreated, "Ayah created successfully", result, start)
	} else {
		core.Success(c, http.StatusCreated, "Ayah created successfully", core.EmptyData{}, start)
	}
}