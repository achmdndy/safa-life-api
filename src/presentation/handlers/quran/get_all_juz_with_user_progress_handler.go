package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/dto"
	query "github.com/safalife/core-api/src/application/quran/query"
	"github.com/safalife/core-api/src/presentation/core"
	"github.com/safalife/core-api/src/presentation/middlewares"
)

type GetAllJuzWithUserProgressHandler struct {
	queryHandler *query.QueryHandler
}

func NewGetAllJuzWithUserProgressHandler(q *query.QueryHandler) *GetAllJuzWithUserProgressHandler {
	return &GetAllJuzWithUserProgressHandler{queryHandler: q}
}

// Handle godoc
// @Summary      List all Juz with user ProgressHatam
// @Description  Returns all Juz combined with the authenticated user's hatam progress per Juz
// @Tags         Juz
// @Accept       json
// @Produce      json
// @Param        limit   query   int     false  "Limit"
// @Param        offset  query   int     false  "Offset"
// @Param        include query   string  false  "Include related data" Enums(relations) example(relations)
// @Success      200     {object} JuzWithProgressListSuccessResponse
// @Failure      400     {object} QuranErrorResponse "Bad request"
// @Failure      500     {object} QuranErrorResponse "Internal server error"
// @Router       /quran/juz/with-progress [get]
func (h *GetAllJuzWithUserProgressHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.GetAllJuzRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid query parameters", errorDetail, start)
		return
	}

	result, err := h.queryHandler.GetAllJuzWithUserProgress(ctx, query.GetAllJuzWithUserProgressQuery{
		Limit:   req.Limit,
		Offset:  req.Offset,
		Include: req.Include,
		UserID:  middlewares.GetUserID(c),
	})
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to get juz with progress list", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Juz with progress list retrieved successfully", result, start)
}
