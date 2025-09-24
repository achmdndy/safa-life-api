package reciter

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/reciter/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetAllReciters handles GET /reciters
// GetAllReciters handles GET /reciters
// @Summary Get All Reciters
// @Description Retrieve a list of all available reciters.
// @Tags Reciters
// @Accept json
// @Produce json
// @Success 200 {object} core.SuccessResponse{data=[]dto.ReciterResponse} "Reciters retrieved successfully"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /reciters [get]
func (h *Handler) GetAllReciters(c *gin.Context) {
	start := time.Now()
	queryReq := query.ToGetAllRecitersQuery()
	result, err := h.queryHandler.GetAllReciters(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get reciters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Reciters retrieved successfully", result, start)
}
