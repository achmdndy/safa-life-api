package topic

import (
	"net/http"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/topic/command"
	"github.com/achmdndy/safa-life-api/src/application/topic/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// RemoveAyahFromTopic handles DELETE /topics/:id/ayahs/:surah_id/:ayah_id
// @Summary Remove Ayah from Topic
// @Description Unlink an ayah from a topic.
// @Tags Topics
// @Accept json
// @Produce json
// @Param id path string true "Topic ID"
// @Param surah_id path int true "Surah ID"
// @Param ayah_id path int true "Ayah ID"
// @Success 200 {object} core.SuccessResponse "Ayah removed from topic successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters"
// @Failure 404 {object} core.ErrorResponse "Topic or ayah not found"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /topics/{id}/ayahs/{surah_id}/{ayah_id} [delete]
func (h *Handler) RemoveAyahFromTopic(c *gin.Context) {
	start := time.Now()
	var req dto.RemoveAyahFromTopicRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.RemoveAyahFromTopicCommand(req)
	if err := h.commandHandler.RemoveAyahFromTopic(c.Request.Context(), cmd); err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Topic or ayah not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to remove ayah from topic", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah removed from topic successfully", core.EmptyData{}, start)
}
