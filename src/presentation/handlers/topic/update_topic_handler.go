package topic

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/topic/command"
	"github.com/achmdndy/safa-life-api/src/application/topic/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// UpdateTopic handles PUT /topics/:id
// @Summary Update Topic
// @Description Update an existing topic.
// @Tags Topics
// @Accept json
// @Produce json
// @Param id path string true "Topic ID"
// @Param topic body dto.UpdateTopicRequest true "Topic Data"
// @Success 200 {object} core.SuccessResponse{data=dto.TopicResponse} "Topic updated successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters or request body"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /topics/{id} [put]
func (h *Handler) UpdateTopic(c *gin.Context) {
	start := time.Now()
	var req dto.UpdateTopicRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid URI parameter", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromUpdateTopicRequest(req)
	result, err := h.commandHandler.UpdateTopic(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to update topic", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Topic updated successfully", dto.ToTopicResponse(*result), start)
}
