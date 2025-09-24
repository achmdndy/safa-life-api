package topic

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/topic/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetAyahsForTopic handles GET /topics/:id/ayahs
// @Summary Get Ayahs for a Topic
// @Description Retrieve all ayahs linked to a specific topic.
// @Tags Topics
// @Accept json
// @Produce json
// @Param id path string true "Topic ID"
// @Success 200 {object} core.SuccessResponse{data=[]dto.TopicAyahResponse} "Ayahs for topic retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid ID"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /topics/{id}/ayahs [get]
func (h *Handler) GetAyahsForTopic(c *gin.Context) {
	start := time.Now()
	topicID := c.Param("id")

	queryReq := query.ToGetAyahsForTopicQuery(topicID)
	result, err := h.queryHandler.GetAyahsForTopic(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get ayahs for topic", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayahs for topic retrieved successfully", result, start)
}
