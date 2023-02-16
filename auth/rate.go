package auth

import (
	"sync"
	"golang.org/x/time/rate"
)

// RateLimitManager - Customized limit for each IP address
// Parameters:
// 	visitors (map) - Map from ip address string to limiter
// 	mu (Mutex) - for locking and unlocking
// https://www.alexedwards.net/blog/how-to-rate-limit-http-requests
type RateLimitManager struct {
	visitors map[string]*rate.Limiter
	mu sync.Mutex
	maxPerSec rate.Limit
	burst int
}

// CreateRateLimitManager - Create a new rate manager
// Arguments:
// 	maxPerSec (int) - Allowed rate from IP address
// 	burst (int) - Allowed maximum burst rate
func CreateRateLimitManager(maxPerSec int, burst int) (*RateLimitManager, error) {
	manager := &RateLimitManager{
		maxPerSec: rate.Limit(maxPerSec),
		burst: burst,
		visitors: make(map[string]*rate.Limiter),
	}
	return manager, nil
}

// GetVisitor - Retrieve and return the rate limiter for the current 
// visitor if it already exists. Otherwise create a new rate limiter 
// and add it to the visitors map, using the IP address as the key
func (manager *RateLimitManager) GetVisitor(ip string) *rate.Limiter {
	manager.mu.Lock()

	// defer executes after function is done
	defer manager.mu.Unlock()

	limit, ok := manager.visitors[ip]
	if !ok {
		limit = rate.NewLimiter(manager.maxPerSec, manager.burst)
		manager.visitors[ip] = limit
	}
	return limit
}



