package middleware

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// ipEntry stores the number of requests and the reset time for each IP
type ipEntry struct {
	Mu       sync.Mutex
	Requests int
	ResetAt  time.Time
}

// ipStore keeps track of rate-limiting data per IP address
var ipStore sync.Map

// IPRateLimiter returns a middleware that limits the number of requests per IP within a given time window
func IPRateLimiter(limit int, window time.Duration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ip := extractIP(c)
			now := time.Now()

			val, _ := ipStore.LoadOrStore(ip, &ipEntry{
				Requests: 0,
				ResetAt:  now.Add(window),
			})

			entry := val.(*ipEntry)

			entry.Mu.Lock()
			defer entry.Mu.Unlock()

			// If current time exceeds the reset time, start a new window
			if now.After(entry.ResetAt) {
				entry.Requests = 1
				entry.ResetAt = now.Add(window)
			} else {
				if entry.Requests >= limit {
					retry := int(entry.ResetAt.Sub(now).Seconds())
					c.Response().Header().Set("Retry-After", fmt.Sprintf("%d", retry))
					return c.JSON(http.StatusTooManyRequests, echo.Map{
						"error":       "rate limit exceeded",
						"retry_after": retry,
					})
				}
				entry.Requests++
			}
			return next(c)
		}
	}
}

func extractIP(c echo.Context) string {
	ip := c.RealIP()
	if ip == "" {
		var err error
		ip, _, err = net.SplitHostPort(c.Request().RemoteAddr)
		if err != nil {
			log.Warnf("failed to extract ip address %v", err)
		}
	}
	return ip
}
