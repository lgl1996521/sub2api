package admin

import (
	"github.com/Wei-Shaw/sub2api/internal/pkg/response"
	"github.com/Wei-Shaw/sub2api/internal/service"

	"github.com/gin-gonic/gin"
)

// ServerLinesHandler provides admin CRUD for the public server-line list.
type ServerLinesHandler struct {
	service *service.ServerLinesService
}

// NewServerLinesHandler creates a new ServerLinesHandler.
func NewServerLinesHandler(svc *service.ServerLinesService) *ServerLinesHandler {
	return &ServerLinesHandler{service: svc}
}

// List returns all configured server lines (including disabled entries).
// GET /api/v1/admin/server-lines
func (h *ServerLinesHandler) List(c *gin.Context) {
	lines, err := h.service.List(c.Request.Context())
	if err != nil {
		response.ErrorFrom(c, err)
		return
	}
	response.Success(c, lines)
}

// ListWithStatus returns server lines enriched with current probe status.
// GET /api/v1/admin/server-lines/status
func (h *ServerLinesHandler) ListWithStatus(c *gin.Context) {
	lines, err := h.service.ListWithStatus(c.Request.Context())
	if err != nil {
		response.ErrorFrom(c, err)
		return
	}
	response.Success(c, lines)
}

// SaveRequest is the admin save payload for the server lines list.
type SaveRequest struct {
	Lines []service.ServerLine `json:"lines"`
}

// Save replaces the stored list of server lines.
// PUT /api/v1/admin/server-lines
func (h *ServerLinesHandler) Save(c *gin.Context) {
	var req SaveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request: "+err.Error())
		return
	}
	saved, err := h.service.Save(c.Request.Context(), req.Lines)
	if err != nil {
		response.ErrorFrom(c, err)
		return
	}
	response.Success(c, saved)
}
