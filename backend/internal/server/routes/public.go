package routes

import (
	"github.com/Wei-Shaw/sub2api/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterPublicRoutes registers unauthenticated endpoints used by the
// marketing homepage (plans, model health, server lines).
func RegisterPublicRoutes(v1 *gin.RouterGroup, h *handler.Handlers) {
	public := v1.Group("/public")
	{
		public.GET("/plans", h.Public.GetPlans)
		public.GET("/model-health", h.Public.GetModelHealth)
		public.GET("/server-lines", h.Public.GetServerLines)
	}
}
