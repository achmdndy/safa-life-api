package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/achmdndy/safa-life-api/src/application/quran/command"
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

// CreateJuz handles POST /juz
// @Summary Create a new juz
// @Description Create a new juz (part) in the Quran
// @Tags Juz
// @Accept json
// @Produce json
// @Param juz body dto.CreateJuzRequest true "Juz data"
// @Success 201 {object} core.SuccessResponse{data=dto.JuzResponse} "Juz created successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid request body"
// @Failure 500 {object} core.ErrorResponse "Failed to create juz"
// @Router /juz [post]
func (h *Handler) CreateJuz(c *gin.Context) {
	start := time.Now()
	
	var req dto.CreateJuzRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromCreateJuzRequest(req)
	domainResult, err := h.commandHandler.CreateJuz(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to create juz", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if domainResult != nil {
		result := dto.ToJuzResponse(*domainResult)
		core.Success(c, http.StatusCreated, "Juz created successfully", result, start)
	} else {
		core.Success(c, http.StatusCreated, "Juz created successfully", core.EmptyData{}, start)
	}
}