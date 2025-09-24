package topic

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/topic/command"
	"github.com/achmdndy/safa-life-api/src/application/topic/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// AddAyahToTopic handles POST /topics/:id/ayahs
// @Summary Add Ayah to Topic
// @Description Link an ayah to a topic.
// @Tags Topics
// @Accept json
// @Produce json
// @Param id path string true "Topic ID"
// @Param ayah_info body dto.AddAyahToTopicRequest true "Ayah Info"
// @Success 201 {object} core.SuccessResponse{data=dto.TopicAyahResponse} "Ayah added to topic successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid request body or parameters"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /topics/{id}/ayahs [post]
func (h *Handler) AddAyahToTopic(c *gin.Context) {
	start := time.Now()
	var req dto.AddAyahToTopicRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid URI parameter", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromAddAyahToTopicRequest(req)
	result, err := h.commandHandler.AddAyahToTopic(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to add ayah to topic", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusCreated, "Ayah added to topic successfully", dto.ToTopicAyahResponse(*result), start)
}
