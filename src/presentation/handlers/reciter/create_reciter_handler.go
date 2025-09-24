package reciter

import (
	"net/http"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/reciter/command"
	"github.com/achmdndy/safa-life-api/src/application/reciter/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// CreateReciter handles POST /reciters
// CreateReciter handles POST /reciters
// @Summary Create Reciter
// @Description Add a new reciter to the system.
// @Tags Reciters
// @Accept json
// @Produce json
// @Param reciter body dto.CreateReciterRequest true "Reciter Data"
// @Success 201 {object} core.SuccessResponse{data=dto.ReciterResponse} "Reciter created successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid request body"
// @Failure 409 {object} core.ErrorResponse "Reciter already exists"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /reciters [post]
func (h *Handler) CreateReciter(c *gin.Context) {
	start := time.Now()
	var req dto.CreateReciterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromCreateReciterRequest(req)
	result, err := h.commandHandler.CreateReciter(c.Request.Context(), cmd)
	if err != nil {
		// Check for duplicate key constraint violations
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "UNIQUE constraint failed") {
			core.Error(c, http.StatusConflict, "Reciter already exists", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to create reciter", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusCreated, "Reciter created successfully", dto.ToReciterResponse(*result), start)
}
