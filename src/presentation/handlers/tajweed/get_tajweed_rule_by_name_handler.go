package tajweed

import (
	"net/http"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tajweed/dto"
	"github.com/achmdndy/safa-life-api/src/application/tajweed/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetTajweedRuleByName handles GET /tajweed/rules/name/:rule_name
// GetTajweedRuleByName handles GET /tajweed/rules/name/:rule_name
// @Summary Get Tajweed Rule By Name
// @Description Retrieve a specific tajweed rule by its name.
// @Tags Tajweed Rules
// @Accept json
// @Produce json
// @Param rule_name path string true "Rule Name"
// @Success 200 {object} core.SuccessResponse{data=dto.TajweedRuleResponse} "Tajweed rule retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid rule name"
// @Failure 404 {object} core.ErrorResponse "Tajweed rule not found"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /tajweed/rules/name/{rule_name} [get]
func (h *Handler) GetTajweedRuleByName(c *gin.Context) {
	start := time.Now()
	var req dto.GetTajweedRuleRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.ToGetTajweedRuleByNameQuery(req.RuleName)
	result, err := h.queryHandler.GetTajweedRuleByName(c.Request.Context(), queryReq)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Tajweed rule not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get tajweed rule", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result == nil {
		core.Error(c, http.StatusNotFound, "Tajweed rule not found", &core.ErrorDetail{Reason: "Tajweed rule not found"}, start)
		return
	}

	core.Success(c, http.StatusOK, "Tajweed rule retrieved successfully", result, start)
}
