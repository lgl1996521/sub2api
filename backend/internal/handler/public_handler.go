package handler

import (
	"context"
	"sort"
	"strings"
	"time"

	"github.com/Wei-Shaw/sub2api/internal/pkg/response"
	"github.com/Wei-Shaw/sub2api/internal/service"

	"github.com/gin-gonic/gin"
)

// PublicHandler exposes unauthenticated homepage data: plans, model health
// snapshots and server-line status. None of the endpoints surface sensitive
// account-level details.
type PublicHandler struct {
	configService      *service.PaymentConfigService
	opsService         *service.OpsService
	serverLinesService *service.ServerLinesService
}

// NewPublicHandler creates a new PublicHandler.
func NewPublicHandler(
	configService *service.PaymentConfigService,
	opsService *service.OpsService,
	serverLinesService *service.ServerLinesService,
) *PublicHandler {
	return &PublicHandler{
		configService:      configService,
		opsService:         opsService,
		serverLinesService: serverLinesService,
	}
}

// publicPlanDTO is the minimal plan shape used on the marketing homepage.
type publicPlanDTO struct {
	ID                 int64    `json:"id"`
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	Price              float64  `json:"price"`
	OriginalPrice      *float64 `json:"original_price,omitempty"`
	ValidityDays       int      `json:"validity_days"`
	ValidityUnit       string   `json:"validity_unit"`
	Features           []string `json:"features"`
	ProductName        string   `json:"product_name"`
	SortOrder          int      `json:"sort_order"`
	GroupID            int64    `json:"group_id"`
	GroupName          string   `json:"group_name"`
	GroupPlatform      string   `json:"group_platform"`
	RateMultiplier     float64  `json:"rate_multiplier"`
	DailyLimitUSD      *float64 `json:"daily_limit_usd,omitempty"`
	WeeklyLimitUSD     *float64 `json:"weekly_limit_usd,omitempty"`
	MonthlyLimitUSD    *float64 `json:"monthly_limit_usd,omitempty"`
	SupportedScopes    []string `json:"supported_model_scopes"`
}

// GetPlans returns public subscription plans (for_sale=true) enriched with the
// owning group's platform / rate / limits so the homepage can render them
// without hitting any authenticated endpoints.
// GET /api/v1/public/plans
func (h *PublicHandler) GetPlans(c *gin.Context) {
	if h.configService == nil {
		response.Success(c, []publicPlanDTO{})
		return
	}
	ctx := c.Request.Context()
	plans, err := h.configService.ListPlansForSale(ctx)
	if err != nil {
		response.ErrorFrom(c, err)
		return
	}

	info := h.configService.GetGroupInfoMap(ctx, plans)
	out := make([]publicPlanDTO, 0, len(plans))
	for _, p := range plans {
		gi := info[p.GroupID]
		out = append(out, publicPlanDTO{
			ID:              int64(p.ID),
			Name:            p.Name,
			Description:     p.Description,
			Price:           p.Price,
			OriginalPrice:   p.OriginalPrice,
			ValidityDays:    p.ValidityDays,
			ValidityUnit:    p.ValidityUnit,
			Features:        splitFeatures(p.Features),
			ProductName:     p.ProductName,
			SortOrder:       p.SortOrder,
			GroupID:         p.GroupID,
			GroupName:       gi.Name,
			GroupPlatform:   gi.Platform,
			RateMultiplier:  gi.RateMultiplier,
			DailyLimitUSD:   gi.DailyLimitUSD,
			WeeklyLimitUSD:  gi.WeeklyLimitUSD,
			MonthlyLimitUSD: gi.MonthlyLimitUSD,
			SupportedScopes: gi.ModelScopes,
		})
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].SortOrder != out[j].SortOrder {
			return out[i].SortOrder < out[j].SortOrder
		}
		return out[i].Price < out[j].Price
	})
	response.Success(c, out)
}

// publicHealthItem is a per-platform health summary for the homepage.
type publicHealthItem struct {
	Platform       string `json:"platform"`
	DisplayName    string `json:"display_name"`
	TotalAccounts  int64  `json:"total_accounts"`
	AvailableCount int64  `json:"available_count"`
	RateLimitCount int64  `json:"rate_limit_count"`
	ErrorCount     int64  `json:"error_count"`
	Status         string `json:"status"`
}

// publicHealthResponse wraps the homepage health payload.
type publicHealthResponse struct {
	Enabled     bool               `json:"enabled"`
	CollectedAt *time.Time         `json:"collected_at,omitempty"`
	Platforms   []publicHealthItem `json:"platforms"`
}

// GetModelHealth returns aggregated platform-level health suitable for an
// unauthenticated homepage banner. It never includes account-level identifiers
// or error messages.
// GET /api/v1/public/model-health
func (h *PublicHandler) GetModelHealth(c *gin.Context) {
	resp := publicHealthResponse{Platforms: []publicHealthItem{}}
	if h.opsService == nil {
		response.Success(c, resp)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	if !h.opsService.IsRealtimeMonitoringEnabled(ctx) {
		response.Success(c, resp)
		return
	}

	platformStats, _, _, collectedAt, err := h.opsService.GetAccountAvailabilityStats(ctx, "", nil)
	if err != nil {
		response.Success(c, resp)
		return
	}

	items := make([]publicHealthItem, 0, len(platformStats))
	for _, p := range platformStats {
		if p == nil {
			continue
		}
		items = append(items, publicHealthItem{
			Platform:       p.Platform,
			DisplayName:    platformDisplayName(p.Platform),
			TotalAccounts:  p.TotalAccounts,
			AvailableCount: p.AvailableCount,
			RateLimitCount: p.RateLimitCount,
			ErrorCount:     p.ErrorCount,
			Status:         classifyAvailability(p.AvailableCount, p.TotalAccounts),
		})
	}
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Platform < items[j].Platform
	})

	resp.Enabled = true
	resp.Platforms = items
	resp.CollectedAt = collectedAt
	response.Success(c, resp)
}

// GetServerLines returns the public list of server lines with probe status.
// GET /api/v1/public/server-lines
func (h *PublicHandler) GetServerLines(c *gin.Context) {
	if h.serverLinesService == nil {
		response.Success(c, []service.ServerLineWithStatus{})
		return
	}
	lines, err := h.serverLinesService.ListWithStatus(c.Request.Context())
	if err != nil {
		response.ErrorFrom(c, err)
		return
	}
	response.Success(c, lines)
}

// classifyAvailability maps a (available, total) pair into a coarse status
// label for homepage display.
func classifyAvailability(available, total int64) string {
	if total <= 0 {
		return "unknown"
	}
	ratio := float64(available) / float64(total)
	switch {
	case available == 0:
		return "down"
	case ratio < 0.5:
		return "degraded"
	default:
		return "healthy"
	}
}

// platformDisplayName returns a friendlier label for a platform identifier.
func platformDisplayName(platform string) string {
	switch strings.ToLower(platform) {
	case "claude", "anthropic":
		return "Claude"
	case "openai":
		return "OpenAI"
	case "codex":
		return "Codex"
	case "gemini":
		return "Gemini"
	case "antigravity":
		return "Antigravity"
	case "bedrock":
		return "Bedrock"
	default:
		if platform == "" {
			return "Unknown"
		}
		return strings.ToUpper(platform[:1]) + platform[1:]
	}
}

// splitFeatures parses the newline-separated feature string from a plan into a
// trimmed list, dropping empty entries.
func splitFeatures(raw string) []string {
	if strings.TrimSpace(raw) == "" {
		return []string{}
	}
	parts := strings.Split(raw, "\n")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		trimmed := strings.TrimSpace(p)
		if trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out
}
