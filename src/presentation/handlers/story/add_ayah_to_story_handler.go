package story

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/story/command"
	"github.com/achmdndy/safa-life-api/src/application/story/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// AddAyahToStory handles POST /stories/:id/ayahs
// @Summary Add Ayah to Story
// @Description Link an ayah to a story.
// @Tags Stories
// @Accept json
// @Produce json
// @Param id path string true "Story ID"
// @Param ayah_info body dto.AddAyahToStoryRequest true "Ayah Info"
// @Success 201 {object} core.SuccessResponse{data=dto.StoryAyahResponse} "Ayah added to story successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid request body or parameters"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /stories/{id}/ayahs [post]
func (h *Handler) AddAyahToStory(c *gin.Context) {
	start := time.Now()
	var req dto.AddAyahToStoryRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid URI parameter", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromAddAyahToStoryRequest(req)
	result, err := h.commandHandler.AddAyahToStory(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to add ayah to story", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusCreated, "Ayah added to story successfully", dto.ToStoryAyahResponse(*result), start)
}
