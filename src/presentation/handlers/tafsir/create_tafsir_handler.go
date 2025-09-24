package tafsir

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tafsir/command"
	"github.com/achmdndy/safa-life-api/src/application/tafsir/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// CreateTafsir handles POST /tafsirs/editions
// @Summary Create Tafsir Edition
// @Description Add a new tafsir edition.
// @Tags Tafsir Editions
// @Accept json
// @Produce json
// @Param edition body dto.CreateTafsirRequest true "Tafsir Edition Data"
// @Success 201 {object} core.SuccessResponse{data=dto.TafsirResponse}
// @Failure 400 {object} core.ErrorResponse
// @Failure 500 {object} core.ErrorResponse
// @Router /tafsirs/editions [post]
func (h *Handler) CreateTafsir(c *gin.Context) {
	start := time.Now()
	var req dto.CreateTafsirRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromCreateTafsirRequest(req)
	result, err := h.commandHandler.CreateTafsir(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to create tafsir edition", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusCreated, "Tafsir edition created successfully", dto.ToTafsirResponse(*result), start)
}
