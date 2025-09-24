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

// UpdateStory handles PUT /stories/:id
// @Summary Update Story
// @Description Update an existing story.
// @Tags Stories
// @Accept json
// @Produce json
// @Param id path string true "Story ID"
// @Param story body dto.UpdateStoryRequest true "Story Data"
// @Success 200 {object} core.SuccessResponse{data=dto.StoryResponse} "Story updated successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters or request body"
// @Failure 404 {object} core.ErrorResponse "Story not found"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /stories/{id} [put]
func (h *Handler) UpdateStory(c *gin.Context) {
	start := time.Now()
	var req dto.UpdateStoryRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid URI parameter", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromUpdateStoryRequest(req)
	result, err := h.commandHandler.UpdateStory(c.Request.Context(), cmd)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Story not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to update story", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result == nil {
		core.Error(c, http.StatusNotFound, "Story not found", &core.ErrorDetail{}, start)
		return
	}

	core.Success(c, http.StatusOK, "Story updated successfully", dto.ToStoryResponse(*result), start)
}
