package auth

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/daiwikmh/origami/models"
)

// KeyStore manages API keys in memory
type KeyStore struct {
	keys      map[string]*models.APIKey
	rateLimit map[string]*models.RateLimitInfo
	mu        sync.RWMutex
}

// NewKeyStore creates a new key store
func NewKeyStore() *KeyStore {
	store := &KeyStore{
		keys:      make(map[string]*models.APIKey),
		rateLimit: make(map[string]*models.RateLimitInfo),
	}

	// Create a default API key for testing
	defaultKey := store.GenerateKey("Default Test Key", 100)
	store.keys[defaultKey.Key] = defaultKey

	return store
}

// GenerateKey creates a new API key
func (ks *KeyStore) GenerateKey(name string, rateLimit int) *models.APIKey {
	ks.mu.Lock()
	defer ks.mu.Unlock()

	key := generateRandomKey()
	apiKey := &models.APIKey{
		Key:           key,
		Name:          name,
		CreatedAt:     time.Now(),
		RateLimit:     rateLimit,
		RequestCount:  0,
		IsActive:      true,
		EndpointUsage: make(map[string]int64),
	}

	ks.keys[key] = apiKey
	return apiKey
}

// ValidateKey checks if a key exists and is active
func (ks *KeyStore) ValidateKey(key string) (*models.APIKey, bool) {
	ks.mu.RLock()
	defer ks.mu.RUnlock()

	apiKey, exists := ks.keys[key]
	if !exists || !apiKey.IsActive {
		return nil, false
	}

	return apiKey, true
}

// RevokeKey deactivates an API key
func (ks *KeyStore) RevokeKey(key string) bool {
	ks.mu.Lock()
	defer ks.mu.Unlock()

	apiKey, exists := ks.keys[key]
	if !exists {
		return false
	}

	apiKey.IsActive = false
	return true
}

// ListKeys returns all API keys
func (ks *KeyStore) ListKeys() []*models.APIKey {
	ks.mu.RLock()
	defer ks.mu.RUnlock()

	keys := make([]*models.APIKey, 0, len(ks.keys))
	for _, key := range ks.keys {
		keys = append(keys, key)
	}

	return keys
}

// UpdateLastUsed updates the last used timestamp for a key
func (ks *KeyStore) UpdateLastUsed(key string) {
	ks.mu.Lock()
	defer ks.mu.Unlock()

	if apiKey, exists := ks.keys[key]; exists {
		now := time.Now()
		apiKey.LastUsedAt = &now
		apiKey.RequestCount++
	}
}

// TrackEndpointUsage increments usage counter for a specific endpoint
func (ks *KeyStore) TrackEndpointUsage(key string, endpoint string) {
	ks.mu.Lock()
	defer ks.mu.Unlock()

	if apiKey, exists := ks.keys[key]; exists {
		apiKey.EndpointUsage[endpoint]++
	}
}

// CheckRateLimit checks if the key has exceeded rate limit
func (ks *KeyStore) CheckRateLimit(key string, limit int) bool {
	ks.mu.Lock()
	defer ks.mu.Unlock()

	now := time.Now()
	info, exists := ks.rateLimit[key]

	// If no record or window expired, create new window
	if !exists || now.Sub(info.WindowStart) >= time.Minute {
		ks.rateLimit[key] = &models.RateLimitInfo{
			WindowStart: now,
			Count:       1,
		}
		return true
	}

	// Check if within limit
	if info.Count >= limit {
		return false
	}

	// Increment counter
	info.Count++
	return true
}

// GetUsageStats returns aggregated usage statistics
func (ks *KeyStore) GetUsageStats() *models.UsageStats {
	ks.mu.RLock()
	defer ks.mu.RUnlock()

	stats := &models.UsageStats{
		TotalKeys:     len(ks.keys),
		ActiveKeys:    0,
		TotalRequests: 0,
		KeyStats:      make(map[string]*models.KeyUsageStats),
	}

	for key, apiKey := range ks.keys {
		if apiKey.IsActive {
			stats.ActiveKeys++
		}

		stats.TotalRequests += apiKey.RequestCount

		stats.KeyStats[key] = &models.KeyUsageStats{
			Name:          apiKey.Name,
			RequestCount:  apiKey.RequestCount,
			LastUsedAt:    apiKey.LastUsedAt,
			RateLimit:     apiKey.RateLimit,
			EndpointUsage: apiKey.EndpointUsage,
			CreatedAt:     apiKey.CreatedAt,
		}
	}

	return stats
}

// generateRandomKey creates a cryptographically secure random API key
func generateRandomKey() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return "og_" + hex.EncodeToString(bytes)
}
