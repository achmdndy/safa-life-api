package story

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/story/command"
	"github.com/achmdndy/safa-life-api/src/application/story/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// CreateStory handles POST /stories
// @Summary Create Story
// @Description Add a new story.
// @Tags Stories
// @Accept json
// @Produce json
// @Param story body dto.CreateStoryRequest true "Story Data"
// @Success 201 {object} core.SuccessResponse{data=dto.StoryResponse} "Story created successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid request body"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /stories [post]
func (h *Handler) CreateStory(c *gin.Context) {
	start := time.Now()
	var req dto.CreateStoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromCreateStoryRequest(req)
	result, err := h.commandHandler.CreateStory(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to create story", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusCreated, "Story created successfully", dto.ToStoryResponse(*result), start)
}
