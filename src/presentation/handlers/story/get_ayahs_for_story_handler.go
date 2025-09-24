package story

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/story/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetAyahsForStory handles GET /stories/:id/ayahs
// @Summary Get Ayahs for a Story
// @Description Retrieve all ayahs linked to a specific story.
// @Tags Stories
// @Accept json
// @Produce json
// @Param id path string true "Story ID"
// @Success 200 {object} core.SuccessResponse{data=[]dto.StoryAyahResponse} "Ayahs for story retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid ID"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /stories/{id}/ayahs [get]
func (h *Handler) GetAyahsForStory(c *gin.Context) {
	start := time.Now()
	storyID := c.Param("id")

	queryReq := query.ToGetAyahsForStoryQuery(storyID)
	result, err := h.queryHandler.GetAyahsForStory(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get ayahs for story", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayahs for story retrieved successfully", result, start)
}
