package tajweed

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tajweed/command"
	"github.com/achmdndy/safa-life-api/src/application/tajweed/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// CreateTajweedRule handles POST /tajweed/rules
// CreateTajweedRule handles POST /tajweed/rules
// @Summary Create Tajweed Rule
// @Description Add a new tajweed rule to the system.
// @Tags Tajweed Rules
// @Accept json
// @Produce json
// @Param tajweed_rule body dto.CreateTajweedRuleRequest true "Tajweed Rule Data"
// @Success 201 {object} core.SuccessResponse{data=dto.TajweedRuleResponse} "Tajweed rule created successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid request body"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /tajweed/rules [post]
func (h *Handler) CreateTajweedRule(c *gin.Context) {
	start := time.Now()

	var req dto.CreateTajweedRuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromCreateTajweedRuleRequest(req)
	domainResult, err := h.commandHandler.CreateTajweedRule(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to create tajweed rule", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	result := dto.ToTajweedRuleResponse(*domainResult)
	core.Success(c, http.StatusCreated, "Tajweed rule created successfully", result, start)
}
