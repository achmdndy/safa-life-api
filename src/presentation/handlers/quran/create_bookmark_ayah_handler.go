package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
	"github.com/safalife/core-api/src/presentation/middlewares"
)

// CreateBookmarkAyahHandler handles create bookmark ayah requests
type CreateBookmarkAyahHandler struct {
	commandHandler *command.CommandHandler
}

// NewCreateBookmarkAyahHandler creates a new CreateBookmarkAyahHandler
func NewCreateBookmarkAyahHandler(commandHandler *command.CommandHandler) *CreateBookmarkAyahHandler {
	return &CreateBookmarkAyahHandler{commandHandler: commandHandler}
}

// Handle creates a bookmark for an ayah
// @Summary Create ayah bookmark
// @Description Bookmark an ayah for the current user
// @Tags Bookmarks
// @Accept json
// @Produce json
// @Param request body dto.CreateBookmarkAyahRequest true "Create bookmark request"
// @Success 201 {object} BookmarkAyahSuccessResponse "Bookmark created successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/bookmarks/ayahs [post]
func (h *CreateBookmarkAyahHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.CreateBookmarkAyahRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid request body", errorDetail, start)
		return
	}

	cmd := command.CreateBookmarkAyahCommand{
		UserID:    middlewares.GetUserID(c),
		AyahID:    req.AyahID,
		CreatedBy: middlewares.GetUserID(c),
	}

	result, err := h.commandHandler.CreateBookmarkAyah(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to create bookmark", errorDetail, start)
		return
	}

	core.Success(c, http.StatusCreated, "Bookmark created successfully", result, start)
}
