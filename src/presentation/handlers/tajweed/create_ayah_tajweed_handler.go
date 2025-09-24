package tajweed

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tajweed/command"
	"github.com/achmdndy/safa-life-api/src/application/tajweed/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// CreateAyahTajweed handles POST /tajweed/ayah
// CreateAyahTajweed handles POST /tajweed/ayah
// @Summary Create Ayah Tajweed
// @Description Add a new set of tajweed rules for a specific ayah.
// @Tags Tajweed
// @Accept json
// @Produce json
// @Param ayah_tajweed body dto.CreateAyahTajweedRequest true "Ayah Tajweed Data"
// @Success 201 {object} core.SuccessResponse{data=dto.AyahTajweedResponse} "Ayah tajweed created successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid request body"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /tajweed/ayah [post]
func (h *Handler) CreateAyahTajweed(c *gin.Context) {
	start := time.Now()

	var req dto.CreateAyahTajweedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromCreateAyahTajweedRequest(req)
	domainResult, err := h.commandHandler.CreateAyahTajweed(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to create ayah tajweed", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	result := dto.ToAyahTajweedResponse(*domainResult)
	core.Success(c, http.StatusCreated, "Ayah tajweed created successfully", result, start)
}
