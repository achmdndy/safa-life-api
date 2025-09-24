package topic

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/topic/dto"
	"github.com/achmdndy/safa-life-api/src/application/topic/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetTopicByID handles GET /topics/:id
// @Summary Get Topic By ID
// @Description Retrieve a specific topic by its ID.
// @Tags Topics
// @Accept json
// @Produce json
// @Param id path string true "Topic ID"
// @Success 200 {object} core.SuccessResponse{data=dto.TopicResponse} "Topic retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid ID"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /topics/{id} [get]
func (h *Handler) GetTopicByID(c *gin.Context) {
	start := time.Now()
	var req dto.GetTopicByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.ToGetTopicByIDQuery(req.ID)
	result, err := h.queryHandler.GetTopicByID(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get topic", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Topic retrieved successfully", result, start)
}
