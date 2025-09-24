package topic

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/topic/command"
	"github.com/achmdndy/safa-life-api/src/application/topic/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// CreateTopic handles POST /topics
// @Summary Create Topic
// @Description Add a new topic.
// @Tags Topics
// @Accept json
// @Produce json
// @Param topic body dto.CreateTopicRequest true "Topic Data"
// @Success 201 {object} core.SuccessResponse{data=dto.TopicResponse} "Topic created successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid request body"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /topics [post]
func (h *Handler) CreateTopic(c *gin.Context) {
	start := time.Now()
	var req dto.CreateTopicRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromCreateTopicRequest(req)
	result, err := h.commandHandler.CreateTopic(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to create topic", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusCreated, "Topic created successfully", dto.ToTopicResponse(*result), start)
}
