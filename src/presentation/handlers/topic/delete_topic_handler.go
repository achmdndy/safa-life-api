package topic

import (
	"net/http"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/topic/command"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// DeleteTopic handles DELETE /topics/:id
// @Summary Delete Topic
// @Description Delete a specific topic.
// @Tags Topics
// @Accept json
// @Produce json
// @Param id path string true "Topic ID"
// @Success 200 {object} core.SuccessResponse "Topic deleted successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid ID"
// @Failure 404 {object} core.ErrorResponse "Topic not found"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /topics/{id} [delete]
func (h *Handler) DeleteTopic(c *gin.Context) {
	start := time.Now()
	id := c.Param("id")

	cmd := command.DeleteTopicCommand{ID: id}
	if err := h.commandHandler.DeleteTopic(c.Request.Context(), cmd); err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Topic not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to delete topic", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Topic deleted successfully", core.EmptyData{}, start)
}
