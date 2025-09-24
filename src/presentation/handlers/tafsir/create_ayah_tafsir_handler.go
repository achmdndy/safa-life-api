package tafsir

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tafsir/command"
	"github.com/achmdndy/safa-life-api/src/application/tafsir/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// CreateAyahTafsir handles POST /tafsirs/ayahs
// @Summary Create Ayah Tafsir
// @Description Add a new tafsir for a specific ayah.
// @Tags Ayah Tafsirs
// @Accept json
// @Produce json
// @Param ayah_tafsir body dto.CreateAyahTafsirRequest true "Ayah Tafsir Data"
// @Success 201 {object} core.SuccessResponse{data=dto.AyahTafsirResponse}
// @Failure 400 {object} core.ErrorResponse
// @Failure 500 {object} core.ErrorResponse
// @Router /tafsirs/ayahs [post]
func (h *Handler) CreateAyahTafsir(c *gin.Context) {
	start := time.Now()
	var req dto.CreateAyahTafsirRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromCreateAyahTafsirRequest(req)
	result, err := h.commandHandler.CreateAyahTafsir(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to create ayah tafsir", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusCreated, "Ayah tafsir created successfully", dto.ToAyahTafsirResponse(*result), start)
}
