package echo

import (
	"net/http"
	"time"

	goiprequest "github.com/ihsanardanto-djoin/go-ip-request"
	goratelimit "github.com/ihsanardanto-djoin/go-ratelimit"
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

// RateLimitMiddleware is a middleware function for rate limiting in an Echo framework application.
//
// This function returns a middleware function that wraps the next handler function in the Echo
// framework's middleware chain. It measures each client making the request and calculating
// if it violating the rate limit set by system.
func RateLimitMiddleware(rl *goratelimit.RateLimiter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		go func() {
			for {
				time.Sleep(time.Minute)
				// Lock the mutex to protect this section from race conditions.
				rl.Mu.Lock()
				for ip, client := range rl.Clients {
					if time.Since(client.LastSeen) > 3*time.Minute {
						delete(rl.Clients, ip)
					}
				}
				rl.Mu.Unlock()
			}
		}()

		return func(c echo.Context) error {
			ip, _ := goiprequest.GetClientIP(c.Request())
			rl.Mu.Lock()
			if _, found := rl.Clients[ip]; !found {
				rl.Clients[ip] = &goratelimit.Client{Limiter: rate.NewLimiter(rate.Limit(rl.Limit), rl.Burst)}
			}
			rl.Clients[ip].LastSeen = time.Now()

			if !rl.AllowRequest(ip) {
				rl.Mu.Unlock()
				return c.JSON(http.StatusTooManyRequests, map[string]string{"error": "Rate limit exceeded"})
			}

			rl.Mu.Unlock()
			return next(c)

		}
	}
}
