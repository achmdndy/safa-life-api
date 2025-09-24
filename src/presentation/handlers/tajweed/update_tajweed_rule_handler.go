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

// UpdateTajweedRule handles PUT /tajweed/rules/:id
// UpdateTajweedRule handles PUT /tajweed/rules/:id
// @Summary Update Tajweed Rule
// @Description Update an existing tajweed rule.
// @Tags Tajweed Rules
// @Accept json
// @Produce json
// @Param id path string true "Rule ID"
// @Param tajweed_rule body dto.UpdateTajweedRuleRequest true "Tajweed Rule Data"
// @Success 200 {object} core.SuccessResponse{data=dto.TajweedRuleResponse} "Tajweed rule updated successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters or request body"
// @Failure 404 {object} core.ErrorResponse "Tajweed rule not found"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /tajweed/rules/{id} [put]
func (h *Handler) UpdateTajweedRule(c *gin.Context) {
	start := time.Now()

	var uriReq dto.UpdateTajweedRuleRequest
	if err := c.ShouldBindUri(&uriReq); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid URI parameter", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	var jsonReq dto.UpdateTajweedRuleRequest
	if err := c.ShouldBindJSON(&jsonReq); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	uriReq.Rule = jsonReq.Rule
	uriReq.Explanation = jsonReq.Explanation
	uriReq.Color = jsonReq.Color

	cmd := command.FromUpdateTajweedRuleRequest(uriReq)
	domainResult, err := h.commandHandler.UpdateTajweedRule(c.Request.Context(), cmd)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Tajweed rule not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to update tajweed rule", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if domainResult == nil {
		core.Error(c, http.StatusNotFound, "Tajweed rule not found", &core.ErrorDetail{}, start)
		return
	}

	result := dto.ToTajweedRuleResponse(*domainResult)
	core.Success(c, http.StatusOK, "Tajweed rule updated successfully", result, start)
}
