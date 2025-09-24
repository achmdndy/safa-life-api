package resource

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/resource/command"
	"github.com/achmdndy/safa-life-api/src/application/resource/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// CreateResource handles POST /resources
// @Summary Create Resource
// @Description Add a new downloadable resource.
// @Tags Resources
// @Accept json
// @Produce json
// @Param resource body dto.CreateResourceRequest true "Resource Data"
// @Success 201 {object} core.SuccessResponse{data=dto.ResourceResponse} "Resource created successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid request body"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /resources [post]
func (h *Handler) CreateResource(c *gin.Context) {
	start := time.Now()
	var req dto.CreateResourceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromCreateResourceRequest(req)
	result, err := h.commandHandler.CreateResource(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to create resource", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusCreated, "Resource created successfully", dto.ToResourceResponse(*result), start)
}
