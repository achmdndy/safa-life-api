package story

import (
	"net/http"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/story/command"
	"github.com/achmdndy/safa-life-api/src/application/story/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// RemoveAyahFromStory handles DELETE /stories/:id/ayahs/:surah_id/:ayah_id
// @Summary Remove Ayah from Story
// @Description Unlink an ayah from a story.
// @Tags Stories
// @Accept json
// @Produce json
// @Param id path string true "Story ID"
// @Param surah_id path int true "Surah ID"
// @Param ayah_id path int true "Ayah ID"
// @Success 200 {object} core.SuccessResponse "Ayah removed from story successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters"
// @Failure 404 {object} core.ErrorResponse "Story or ayah not found"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /stories/{id}/ayahs/{surah_id}/{ayah_id} [delete]
func (h *Handler) RemoveAyahFromStory(c *gin.Context) {
	start := time.Now()
	var req dto.RemoveAyahFromStoryRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.RemoveAyahFromStoryCommand(req)
	if err := h.commandHandler.RemoveAyahFromStory(c.Request.Context(), cmd); err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Story or ayah not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to remove ayah from story", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah removed from story successfully", core.EmptyData{}, start)
}
