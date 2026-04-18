package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	infraerrors "github.com/Wei-Shaw/sub2api/internal/pkg/errors"
)

// SettingKeyServerLines stores the JSON-encoded list of ServerLine entries.
const SettingKeyServerLines = "server_lines"

// ServerLineStatus is the probed status of a server line.
type ServerLineStatus string

const (
	ServerLineStatusHealthy  ServerLineStatus = "healthy"
	ServerLineStatusDegraded ServerLineStatus = "degraded"
	ServerLineStatusDown     ServerLineStatus = "down"
	ServerLineStatusUnknown  ServerLineStatus = "unknown"
)

// ServerLine represents a user-selectable upstream line (region-specific
// gateway deployment) advertised on the public homepage.
type ServerLine struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Region      string `json:"region"`
	URL         string `json:"url"`
	ProbePath   string `json:"probe_path"`
	Description string `json:"description"`
	SortOrder   int    `json:"sort_order"`
	Enabled     bool   `json:"enabled"`
}

// ServerLineWithStatus enriches a ServerLine with the latest probe result.
type ServerLineWithStatus struct {
	ServerLine
	Status    ServerLineStatus `json:"status"`
	LatencyMS int64            `json:"latency_ms"`
	CheckedAt time.Time        `json:"checked_at"`
	Error     string           `json:"error,omitempty"`
}

// serverLineProbeCache caches the most recent probe snapshot for a given line.
type serverLineProbeCache struct {
	status    ServerLineStatus
	latency   time.Duration
	err       string
	checkedAt time.Time
}

// ServerLinesService persists server line definitions via the generic Setting
// key/value store so that no schema migration is required.
type ServerLinesService struct {
	settingRepo SettingRepository
	httpClient  *http.Client

	mu         sync.RWMutex
	probeCache map[string]serverLineProbeCache
	probeTTL   time.Duration
}

// NewServerLinesService constructs a ServerLinesService with default timeouts.
func NewServerLinesService(settingRepo SettingRepository) *ServerLinesService {
	return &ServerLinesService{
		settingRepo: settingRepo,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		probeCache: map[string]serverLineProbeCache{},
		probeTTL:   60 * time.Second,
	}
}

// List returns the raw list of configured server lines. The result is sorted
// by SortOrder then Name for stable ordering.
func (s *ServerLinesService) List(ctx context.Context) ([]ServerLine, error) {
	raw, err := s.settingRepo.GetValue(ctx, SettingKeyServerLines)
	if err != nil && !errors.Is(err, ErrSettingNotFound) {
		return nil, err
	}
	lines := make([]ServerLine, 0)
	if strings.TrimSpace(raw) != "" {
		if err := json.Unmarshal([]byte(raw), &lines); err != nil {
			return nil, infraerrors.InternalServer("SERVER_LINES_PARSE", "failed to decode server_lines setting")
		}
	}
	sortServerLines(lines)
	return lines, nil
}

// ListEnabled returns only lines that the admin has enabled.
func (s *ServerLinesService) ListEnabled(ctx context.Context) ([]ServerLine, error) {
	all, err := s.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]ServerLine, 0, len(all))
	for _, l := range all {
		if l.Enabled {
			out = append(out, l)
		}
	}
	return out, nil
}

// Save replaces the entire set of server lines. It validates each entry,
// normalises defaults (probe_path), and enforces unique IDs.
func (s *ServerLinesService) Save(ctx context.Context, lines []ServerLine) ([]ServerLine, error) {
	normalised := make([]ServerLine, 0, len(lines))
	seen := make(map[string]struct{}, len(lines))
	for i := range lines {
		l := lines[i]
		l.ID = strings.TrimSpace(l.ID)
		l.Name = strings.TrimSpace(l.Name)
		l.Region = strings.TrimSpace(l.Region)
		l.URL = strings.TrimSpace(l.URL)
		l.ProbePath = strings.TrimSpace(l.ProbePath)
		l.Description = strings.TrimSpace(l.Description)

		if l.ID == "" {
			return nil, infraerrors.BadRequest("SERVER_LINE_ID_REQUIRED", "server line id is required")
		}
		if _, dup := seen[l.ID]; dup {
			return nil, infraerrors.BadRequest("SERVER_LINE_ID_DUPLICATE", "server line id must be unique: "+l.ID)
		}
		seen[l.ID] = struct{}{}

		if l.Name == "" {
			return nil, infraerrors.BadRequest("SERVER_LINE_NAME_REQUIRED", "server line name is required")
		}
		if l.URL == "" {
			return nil, infraerrors.BadRequest("SERVER_LINE_URL_REQUIRED", "server line url is required")
		}
		if !strings.HasPrefix(l.URL, "http://") && !strings.HasPrefix(l.URL, "https://") {
			return nil, infraerrors.BadRequest("SERVER_LINE_URL_INVALID", "server line url must start with http(s)://")
		}
		if l.ProbePath == "" {
			l.ProbePath = "/health"
		}
		if l.Region == "" {
			l.Region = "intl"
		}
		normalised = append(normalised, l)
	}
	sortServerLines(normalised)

	encoded, err := json.Marshal(normalised)
	if err != nil {
		return nil, infraerrors.InternalServer("SERVER_LINES_ENCODE", "failed to encode server_lines setting")
	}
	if err := s.settingRepo.Set(ctx, SettingKeyServerLines, string(encoded)); err != nil {
		return nil, err
	}
	s.mu.Lock()
	s.probeCache = map[string]serverLineProbeCache{}
	s.mu.Unlock()
	return normalised, nil
}

// ListWithStatus returns enabled server lines enriched with their latest probe
// status. Probes are executed concurrently and results cached for probeTTL.
func (s *ServerLinesService) ListWithStatus(ctx context.Context) ([]ServerLineWithStatus, error) {
	lines, err := s.ListEnabled(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]ServerLineWithStatus, len(lines))
	var wg sync.WaitGroup
	for i := range lines {
		wg.Add(1)
		go func(idx int, line ServerLine) {
			defer wg.Done()
			probe := s.probe(ctx, line)
			out[idx] = ServerLineWithStatus{
				ServerLine: line,
				Status:     probe.status,
				LatencyMS:  probe.latency.Milliseconds(),
				CheckedAt:  probe.checkedAt,
				Error:      probe.err,
			}
		}(i, lines[i])
	}
	wg.Wait()
	return out, nil
}

// probe performs a HEAD/GET request against ProbePath and returns a cached
// snapshot when a recent one is available.
func (s *ServerLinesService) probe(ctx context.Context, line ServerLine) serverLineProbeCache {
	now := time.Now()

	s.mu.RLock()
	if cached, ok := s.probeCache[line.ID]; ok && now.Sub(cached.checkedAt) < s.probeTTL {
		s.mu.RUnlock()
		return cached
	}
	s.mu.RUnlock()

	target := strings.TrimRight(line.URL, "/") + "/" + strings.TrimLeft(line.ProbePath, "/")
	probeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(probeCtx, http.MethodGet, target, nil)
	if err != nil {
		result := serverLineProbeCache{
			status:    ServerLineStatusUnknown,
			err:       "invalid url",
			checkedAt: now,
		}
		s.storeProbe(line.ID, result)
		return result
	}
	req.Header.Set("User-Agent", "sub2api-server-line-probe/1.0")

	start := time.Now()
	resp, err := s.httpClient.Do(req)
	latency := time.Since(start)

	result := serverLineProbeCache{
		latency:   latency,
		checkedAt: now,
	}

	switch {
	case err != nil:
		result.status = ServerLineStatusDown
		result.err = truncateProbeError(err.Error())
	case resp.StatusCode >= 500:
		result.status = ServerLineStatusDegraded
		result.err = resp.Status
	case resp.StatusCode >= 400:
		// 4xx from the probe URL usually means the path is wrong but the
		// service is reachable, so consider the line degraded rather than down.
		result.status = ServerLineStatusDegraded
		result.err = resp.Status
	case latency > 2*time.Second:
		result.status = ServerLineStatusDegraded
	default:
		result.status = ServerLineStatusHealthy
	}
	if resp != nil {
		_ = resp.Body.Close()
	}

	s.storeProbe(line.ID, result)
	return result
}

func (s *ServerLinesService) storeProbe(id string, p serverLineProbeCache) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.probeCache[id] = p
}

func sortServerLines(lines []ServerLine) {
	sort.SliceStable(lines, func(i, j int) bool {
		if lines[i].SortOrder != lines[j].SortOrder {
			return lines[i].SortOrder < lines[j].SortOrder
		}
		return strings.ToLower(lines[i].Name) < strings.ToLower(lines[j].Name)
	})
}

func truncateProbeError(msg string) string {
	const maxLen = 200
	if len(msg) <= maxLen {
		return msg
	}
	return msg[:maxLen] + "…"
}
