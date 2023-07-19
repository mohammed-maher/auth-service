package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

func RateLimited() gin.HandlerFunc {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}
	var (
		mutex   sync.Mutex
		clients = make(map[string]*client)
	)

	// launch a goroutine to watch clients last seen and invalidate limits
	go func() {
		for {
			time.Sleep(time.Minute)
			// Lock the mutex to protect this section from race conditions.
			mutex.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mutex.Unlock()
		}
	}()

	return func(ctx *gin.Context) {
		limitKey := keyFunc(ctx)
		mutex.Lock()
		if _, found := clients[limitKey]; !found {
			clients[limitKey] = &client{limiter: rate.NewLimiter(2, 4)}
		}
		clients[limitKey].lastSeen = time.Now()

		if !clients[limitKey].limiter.Allow() {
			mutex.Unlock()
			ctx.AbortWithError(http.StatusTooManyRequests, errors.New("requests limit hit"))
			return
		}
		mutex.Unlock()

		ctx.Next()
	}
}

// use auth token to limit rate if exists, otherwise IP
func keyFunc(ctx *gin.Context) string {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader != "" {
		return authHeader
	}
	return ctx.RemoteIP()
}
