package tajweed

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tajweed/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetAllTajweedRules handles GET /tajweed/rules
// GetAllTajweedRules handles GET /tajweed/rules
// @Summary Get All Tajweed Rules
// @Description Retrieve a list of all available tajweed rules.
// @Tags Tajweed Rules
// @Accept json
// @Produce json
// @Success 200 {object} core.SuccessResponse{data=[]dto.TajweedRuleResponse} "Tajweed rules retrieved successfully"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /tajweed/rules [get]
func (h *Handler) GetAllTajweedRules(c *gin.Context) {
	start := time.Now()

	queryReq := query.ToGetAllTajweedRulesQuery()
	result, err := h.queryHandler.GetAllTajweedRules(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get tajweed rules", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Tajweed rules retrieved successfully", result, start)
}
