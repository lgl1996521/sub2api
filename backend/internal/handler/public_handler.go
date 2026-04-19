package handler

import (
	"context"
	"math/rand"
	"os"
	"sort"
	"strings"
	"sync"
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

// publicModelPricing describes the unit pricing for a single model on the
// /models pricing catalog. Prices are per 1 million tokens (MTOK). CNY is the
// display price; USD is shown as the smaller secondary number.
type publicModelPricing struct {
	InputCNY     float64 `json:"input_cny"`
	InputUSD     float64 `json:"input_usd"`
	OutputCNY    float64 `json:"output_cny"`
	OutputUSD    float64 `json:"output_usd"`
	CacheReadCNY float64 `json:"cache_read_cny"`
	CacheReadUSD float64 `json:"cache_read_usd"`
}

// publicModelItem is a per-model snapshot shown on the /models catalog. It
// mixes static catalog info (name, platform, tags, release date, pricing) with
// a health summary (status, availability, latency, rolling samples).
type publicModelItem struct {
	ModelName       string             `json:"model_name"`
	Platform        string             `json:"platform"`
	DisplayName     string             `json:"display_name"`
	Description     string             `json:"description,omitempty"`
	ReleaseDate     string             `json:"release_date,omitempty"`
	Tags            []string           `json:"tags,omitempty"`
	Pricing         publicModelPricing `json:"pricing"`
	Status          string             `json:"status"`
	AvailabilityPct int                `json:"availability_pct"`
	AvgLatencyMS    int64              `json:"avg_latency_ms"`
	TTFTMs          int64              `json:"ttft_ms,omitempty"`
	RecentSamples   []int              `json:"recent_samples"`
}

// publicHealthResponse wraps the homepage health payload. Models is only
// populated when the operator has opted in (e.g. SEED_DEMO=true for local
// preview).
type publicHealthResponse struct {
	Enabled     bool               `json:"enabled"`
	CollectedAt *time.Time         `json:"collected_at,omitempty"`
	Platforms   []publicHealthItem `json:"platforms"`
	Models      []publicModelItem  `json:"models,omitempty"`
}

// GetModelHealth returns aggregated platform-level health suitable for an
// unauthenticated homepage banner. It never includes account-level identifiers
// or error messages.
// GET /api/v1/public/model-health
func (h *PublicHandler) GetModelHealth(c *gin.Context) {
	resp := publicHealthResponse{Platforms: []publicHealthItem{}}
	demo := demoModelHealthEnabled()

	if h.opsService == nil {
		if demo {
			resp.Enabled = true
			resp.Platforms = demoPlatformSnapshots()
			resp.Models = demoModelSnapshots()
		}
		response.Success(c, resp)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	if !h.opsService.IsRealtimeMonitoringEnabled(ctx) {
		if demo {
			resp.Enabled = true
			resp.Platforms = demoPlatformSnapshots()
			resp.Models = demoModelSnapshots()
		}
		response.Success(c, resp)
		return
	}

	platformStats, _, _, collectedAt, err := h.opsService.GetAccountAvailabilityStats(ctx, "", nil)
	if err != nil {
		if demo {
			resp.Enabled = true
			resp.Platforms = demoPlatformSnapshots()
			resp.Models = demoModelSnapshots()
		}
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

	if demoModelHealthEnabled() {
		resp.Models = demoModelSnapshots()
	}
	response.Success(c, resp)
}

// demoModelHealthEnabled reports whether SEED_DEMO=true is set in the env.
// When enabled, the homepage /models feed is populated with deterministic
// sample data so operators can preview the UI without real traffic.
func demoModelHealthEnabled() bool {
	v := strings.ToLower(strings.TrimSpace(os.Getenv("SEED_DEMO")))
	return v == "1" || v == "true" || v == "yes" || v == "on"
}

var (
	demoModelRand   = rand.New(rand.NewSource(42))
	demoModelRandMu sync.Mutex
)

// demoModelDef describes one fake model row for the SEED_DEMO preview.
type demoModelDef struct {
	name        string
	platform    string
	displayName string
	description string
	releaseDate string
	tags        []string
	baseLatency int64
	baseTTFT    int64
	// downRate is the share of the 60-sample window that is "down" (0).
	downRate float64
	// Pricing per 1M tokens (USD). CNY is derived by multiplying by the
	// demoUSDRate exchange rate.
	inputUSD     float64
	outputUSD    float64
	cacheReadUSD float64
}

// demoUSDRate is the display exchange rate used to convert USD pricing into
// the large CNY headline number on each model card. Tunable so operators can
// experiment without restarting the backend.
const demoUSDRate = 7.2

func demoModelCatalog() []demoModelDef {
	caching := []string{"Prompt Caching"}
	tiered := []string{"Prompt Caching", "分段定价"}
	return []demoModelDef{
		{"claude-opus-4-6", "anthropic", "Claude Opus 4.6", "Anthropic 旗舰模型，长上下文与推理", "2026-02-14", caching, 2800, 1100, 0.0, 15.0, 75.0, 1.5},
		{"claude-sonnet-4-6", "anthropic", "Claude Sonnet 4.6", "性价比最高的 Claude 模型，适合代码与写作", "2026-01-28", caching, 1600, 650, 0.0, 3.0, 15.0, 0.3},
		{"claude-haiku-4-5", "anthropic", "Claude Haiku 4.5", "极低延迟，短回复场景推荐", "2025-12-06", caching, 720, 280, 0.02, 0.8, 4.0, 0.08},
		{"gpt-5", "openai", "GPT-5", "OpenAI 新一代旗舰对话模型", "2026-03-05", tiered, 2100, 1300, 0.0, 2.5, 15.0, 0.25},
		{"gpt-5-codex", "openai", "GPT-5 Codex", "代码能力增强，专为 Codex/IDE 优化", "2026-02-24", caching, 2300, 1200, 0.0, 1.75, 14.0, 0.175},
		{"gpt-4o", "openai", "GPT-4o", "多模态，图像/语音输入支持", "2025-05-13", caching, 1400, 720, 0.02, 2.5, 10.0, 1.25},
		{"gemini-2.5-pro", "gemini", "Gemini 2.5 Pro", "Google 长上下文旗舰，支持 2M tokens", "2026-01-20", tiered, 3100, 1500, 0.05, 1.25, 10.0, 0.31},
		{"gemini-2.5-flash", "gemini", "Gemini 2.5 Flash", "低延迟快速响应版本", "2025-11-14", caching, 980, 420, 0.02, 0.3, 2.5, 0.075},
		{"antigravity-claude", "antigravity", "Antigravity Claude", "Antigravity 平台代理的 Claude 通道", "2026-03-01", caching, 2450, 1100, 0.1, 2.8, 14.0, 0.28},
		{"antigravity-gemini", "antigravity", "Antigravity Gemini", "Antigravity 平台代理的 Gemini 通道", "2026-02-20", caching, 2650, 1350, 0.1, 1.2, 9.0, 0.3},
	}
}

// demoPlatformSnapshots returns aggregated per-platform sample data derived
// from the SEED_DEMO model catalog.
func demoPlatformSnapshots() []publicHealthItem {
	catalog := demoModelCatalog()
	byPlatform := map[string]*publicHealthItem{}
	for _, d := range catalog {
		p, ok := byPlatform[d.platform]
		if !ok {
			p = &publicHealthItem{
				Platform:    d.platform,
				DisplayName: platformDisplayName(d.platform),
			}
			byPlatform[d.platform] = p
		}
		p.TotalAccounts += 3
		switch {
		case d.downRate >= 0.15:
			p.AvailableCount += 1
			p.RateLimitCount += 1
			p.ErrorCount += 1
		case d.downRate >= 0.05:
			p.AvailableCount += 2
			p.RateLimitCount += 1
		default:
			p.AvailableCount += 3
		}
	}
	out := make([]publicHealthItem, 0, len(byPlatform))
	for _, p := range byPlatform {
		p.Status = classifyAvailability(p.AvailableCount, p.TotalAccounts)
		out = append(out, *p)
	}
	sort.SliceStable(out, func(i, j int) bool {
		return out[i].Platform < out[j].Platform
	})
	return out
}

func demoModelSnapshots() []publicModelItem {
	demoModelRandMu.Lock()
	defer demoModelRandMu.Unlock()

	catalog := demoModelCatalog()
	out := make([]publicModelItem, 0, len(catalog))
	for _, d := range catalog {
		samples := make([]int, 60)
		downCount := 0
		for i := range samples {
			if demoModelRand.Float64() < d.downRate {
				samples[i] = 0
				downCount++
			} else {
				samples[i] = 1
			}
		}
		availability := (60 - downCount) * 100 / 60
		status := "healthy"
		switch {
		case availability < 50:
			status = "down"
		case availability < 95:
			status = "degraded"
		}
		jitterAvg := d.baseLatency + int64(demoModelRand.Intn(400)-200)
		jitterTTFT := d.baseTTFT + int64(demoModelRand.Intn(200)-100)
		if jitterAvg < 0 {
			jitterAvg = d.baseLatency
		}
		if jitterTTFT < 0 {
			jitterTTFT = d.baseTTFT
		}
		pricing := publicModelPricing{
			InputUSD:     d.inputUSD,
			OutputUSD:    d.outputUSD,
			CacheReadUSD: d.cacheReadUSD,
			InputCNY:     roundTwoDecimal(d.inputUSD * demoUSDRate),
			OutputCNY:    roundTwoDecimal(d.outputUSD * demoUSDRate),
			CacheReadCNY: roundTwoDecimal(d.cacheReadUSD * demoUSDRate),
		}
		out = append(out, publicModelItem{
			ModelName:       d.name,
			Platform:        d.platform,
			DisplayName:     d.displayName,
			Description:     d.description,
			ReleaseDate:     d.releaseDate,
			Tags:            append([]string(nil), d.tags...),
			Pricing:         pricing,
			Status:          status,
			AvailabilityPct: availability,
			AvgLatencyMS:    jitterAvg,
			TTFTMs:          jitterTTFT,
			RecentSamples:   samples,
		})
	}
	return out
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

// roundTwoDecimal rounds to two decimal places for display-friendly CNY prices.
func roundTwoDecimal(v float64) float64 {
	return float64(int64(v*100+0.5)) / 100
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
