package story

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/story/dto"
	"github.com/achmdndy/safa-life-api/src/application/story/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetStoryByID handles GET /stories/:id
// @Summary Get Story By ID
// @Description Retrieve a specific story by its ID.
// @Tags Stories
// @Accept json
// @Produce json
// @Param id path string true "Story ID"
// @Success 200 {object} core.SuccessResponse{data=dto.StoryResponse} "Story retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid ID"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /stories/{id} [get]
func (h *Handler) GetStoryByID(c *gin.Context) {
	start := time.Now()
	var req dto.GetStoryByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.ToGetStoryByIDQuery(req.ID)
	result, err := h.queryHandler.GetStoryByID(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get story", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Story retrieved successfully", result, start)
}
