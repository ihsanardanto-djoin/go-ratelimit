package goratelimit

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// Client struct represents a params of rate limit PER IP address
type Client struct {
	Limiter  *rate.Limiter // A pointer which is used for rate limiting.
	LastSeen time.Time     // timestamp representing the last time the client interacted with the system
}

// RateLimiter struct represents a params needed to use rate limit
type RateLimiter struct {
	Limit   int // The maximum number of requests allowed per unit of time
	Burst   int // he maximum number of requests that can be processed in a single burst without violating the rate limit
	Mu      sync.Mutex
	Clients map[string]*Client // track clients(ip) and their request rates
}

// NewRateLimiter creates and initializes a new RateLimiter instance with the specified rate limit and burst limit.
func NewRateLimiter(rateLimit, burstLimit int) *RateLimiter {
	return &RateLimiter{
		Limit:   rateLimit,
		Burst:   burstLimit,
		Clients: make(map[string]*Client),
	}
}

// AllowRequest checks if a request from the specified IP address is allowed based on the configured rate limits.
// Parameters:
// - ip: The IP address of the client for which the rate limit is being checked.
func (rl *RateLimiter) AllowRequest(ip string) bool {
	return rl.Clients[ip].Limiter.Allow()
}
