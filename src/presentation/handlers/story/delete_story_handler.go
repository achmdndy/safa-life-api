package story

import (
	"net/http"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/story/command"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// DeleteStory handles DELETE /stories/:id
// @Summary Delete Story
// @Description Delete a specific story.
// @Tags Stories
// @Accept json
// @Produce json
// @Param id path string true "Story ID"
// @Success 200 {object} core.SuccessResponse "Story deleted successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid ID"
// @Failure 404 {object} core.ErrorResponse "Story not found"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /stories/{id} [delete]
func (h *Handler) DeleteStory(c *gin.Context) {
	start := time.Now()
	id := c.Param("id")

	cmd := command.DeleteStoryCommand{ID: id}
	if err := h.commandHandler.DeleteStory(c.Request.Context(), cmd); err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Story not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to delete story", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Story deleted successfully", core.EmptyData{}, start)
}
