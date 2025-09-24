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

// GetTajweedRuleByID handles GET /tajweed/rules/id/:id
// GetTajweedRuleByID handles GET /tajweed/rules/id/:id
// @Summary Get Tajweed Rule By ID
// @Description Retrieve a specific tajweed rule by its ID.
// @Tags Tajweed Rules
// @Accept json
// @Produce json
// @Param id path string true "Rule ID"
// @Success 200 {object} core.SuccessResponse{data=dto.TajweedRuleResponse} "Tajweed rule retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid rule ID"
// @Failure 404 {object} core.ErrorResponse "Tajweed rule not found"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /tajweed/rules/id/{id} [get]
func (h *Handler) GetTajweedRuleByID(c *gin.Context) {
	start := time.Now()
	var req dto.GetTajweedRuleByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.ToGetTajweedRuleByIDQuery(req.ID)
	result, err := h.queryHandler.GetTajweedRuleByID(c.Request.Context(), queryReq)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Tajweed rule not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get tajweed rule", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result == nil {
		core.Error(c, http.StatusNotFound, "Tajweed rule not found", &core.ErrorDetail{}, start)
		return
	}

	core.Success(c, http.StatusOK, "Tajweed rule retrieved successfully", result, start)
}
