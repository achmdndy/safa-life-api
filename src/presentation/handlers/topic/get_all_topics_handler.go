package topic

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/topic/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetAllTopics handles GET /topics
// @Summary Get All Topics
// @Description Retrieve a list of all topics.
// @Tags Topics
// @Accept json
// @Produce json
// @Success 200 {object} core.SuccessResponse{data=[]dto.TopicResponse} "Topics retrieved successfully"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /topics [get]
func (h *Handler) GetAllTopics(c *gin.Context) {
	start := time.Now()
	queryReq := query.ToGetAllTopicsQuery()
	result, err := h.queryHandler.GetAllTopics(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get topics", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Topics retrieved successfully", result, start)
}
