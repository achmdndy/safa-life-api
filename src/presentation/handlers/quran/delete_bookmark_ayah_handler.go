package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
)

// DeleteBookmarkAyahHandler handles deleting a bookmark ayah by ID
type DeleteBookmarkAyahHandler struct {
	commandHandler *command.CommandHandler
}

// NewDeleteBookmarkAyahHandler creates a new handler instance
func NewDeleteBookmarkAyahHandler(commandHandler *command.CommandHandler) *DeleteBookmarkAyahHandler {
	return &DeleteBookmarkAyahHandler{commandHandler: commandHandler}
}

// Handle processes the request
// @Summary Delete bookmark ayah
// @Description Delete a bookmark by its ID
// @Tags Bookmarks
// @Accept json
// @Produce json
// @Param id path string true "Bookmark ID"
// @Success 200 {object} DeleteBookmarkAyahSuccessResponse "Bookmark deleted successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Bookmark not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/bookmarks/ayahs/{id} [delete]
func (h *DeleteBookmarkAyahHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.DeleteBookmarkAyahRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid bookmark ID", errorDetail, start)
		return
	}

	cmd := command.DeleteBookmarkAyahCommand{ID: req.ID}

	err := h.commandHandler.DeleteBookmarkAyah(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to delete bookmark", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Bookmark deleted successfully", map[string]interface{}{}, start)
}
