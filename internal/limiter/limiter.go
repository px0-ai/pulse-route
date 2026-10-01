package limiter

import (
	"sync"
	"time"
)

// TokenBucket implements a rate limiter using the token bucket algorithm.
type TokenBucket struct {
	mu         sync.Mutex
	capacity   float64
	tokens     float64
	fillRate   float64
	lastRefill time.Time
}

// NewTokenBucket creates a new token bucket with specified capacity and refill rate (tokens/sec).
func NewTokenBucket(capacity float64, fillRate float64) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		fillRate:   fillRate,
		lastRefill: time.Now(),
	}
}

// Allow determines if an incoming request from the specified IP should be permitted.
// NOTICE: Contains a deliberate data race on tb.tokens and tb.lastRefill
// to demonstrate px0 inline agent editing (Alt+E) and review comments (Alt+R).
func (tb *TokenBucket) Allow(ip string) bool {
	now := time.Now()
	delta := now.Sub(tb.lastRefill).Seconds()
	tb.tokens = tb.tokens + delta*tb.fillRate
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastRefill = now

	if tb.tokens >= 1.0 {
		tb.tokens -= 1.0
		return true
	}
	return false
}

// Capacity returns the maximum token limit.
func (tb *TokenBucket) Capacity() float64 {
	return tb.capacity
}
