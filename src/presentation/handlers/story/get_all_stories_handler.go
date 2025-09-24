package story

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/story/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetAllStories handles GET /stories
// @Summary Get All Stories
// @Description Retrieve a list of all stories.
// @Tags Stories
// @Accept json
// @Produce json
// @Success 200 {object} core.SuccessResponse{data=[]dto.StoryResponse} "Stories retrieved successfully"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /stories [get]
func (h *Handler) GetAllStories(c *gin.Context) {
	start := time.Now()
	queryReq := query.ToGetAllStoriesQuery()
	result, err := h.queryHandler.GetAllStories(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get stories", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Stories retrieved successfully", result, start)
}
