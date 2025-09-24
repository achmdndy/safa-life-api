package quran

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/achmdndy/safa-life-api/src/application/quran/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

// SearchHandler handles GET /search
// @Summary Search in Quran
// @Description Search for ayahs (verses) in the Quran by text content
// @Tags Search
// @Accept json
// @Produce json
// @Param q query string true "Search query"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} core.SuccessResponse{data=dto.SearchResponse} "Search completed successfully"
// @Failure 400 {object} core.ErrorResponse "Search query is required"
// @Failure 500 {object} core.ErrorResponse "Failed to perform search"
// @Router /search [get]
func (h *Handler) SearchHandler(c *gin.Context) {
	start := time.Now()
	
	searchQuery := c.Query("q")
	if searchQuery == "" {
		core.Error(c, http.StatusBadRequest, "Search query is required", &core.ErrorDetail{Reason: "Search query parameter 'q' is required"}, start)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	queryReq := query.SearchQuery{
		Query: searchQuery,
		Page:  page,
		Limit: limit,
	}

	result, err := h.queryHandler.Search(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to perform search", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result != nil {
		core.Success(c, http.StatusOK, "Search completed successfully", result, start)
	} else {
		core.Success(c, http.StatusOK, "Search completed successfully", core.EmptyData{}, start)
	}
}