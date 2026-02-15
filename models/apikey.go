package models

import "time"

// APIKey represents an API key with metadata
type APIKey struct {
	Key           string            `json:"key"`
	Name          string            `json:"name"`
	CreatedAt     time.Time         `json:"created_at"`
	LastUsedAt    *time.Time        `json:"last_used_at,omitempty"`
	RateLimit     int               `json:"rate_limit"`      // Requests per minute
	RequestCount  int64             `json:"request_count"`   // Total requests made
	IsActive      bool              `json:"is_active"`
	EndpointUsage map[string]int64  `json:"endpoint_usage"`  // Track usage per endpoint
}

// UsageStats provides aggregated usage statistics
type UsageStats struct {
	TotalKeys       int                        `json:"total_keys"`
	ActiveKeys      int                        `json:"active_keys"`
	TotalRequests   int64                      `json:"total_requests"`
	KeyStats        map[string]*KeyUsageStats  `json:"key_stats"`
}

// KeyUsageStats provides per-key usage information
type KeyUsageStats struct {
	Name          string            `json:"name"`
	RequestCount  int64             `json:"request_count"`
	LastUsedAt    *time.Time        `json:"last_used_at,omitempty"`
	RateLimit     int               `json:"rate_limit"`
	EndpointUsage map[string]int64  `json:"endpoint_usage"`
	CreatedAt     time.Time         `json:"created_at"`
}

// RateLimitInfo tracks rate limiting state
type RateLimitInfo struct {
	WindowStart time.Time
	Count       int
}
