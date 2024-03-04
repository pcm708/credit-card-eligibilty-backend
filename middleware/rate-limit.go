package middleware

import (
	"context"
	"github.com/honestbank/tech-assignment-backend-engineer/handler"
	"net"
	"net/http"
	"strconv"
	"time"
)

var ctx = context.Background()
var rdb = RedisClient()

func NewRateLimiter(limit int, window time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			// Increment the request count
			count, err := rdb.Incr(ctx, ip).Result()
			if err != nil {
				handler.RedisConnectionErrorResponse(err, w)
				return
			}

			if count == 1 {
				// If this is the first request, set the key to expire after the time window
				rdb.Expire(ctx, ip, window)
			}

			if count > int64(limit) {
				handler.RateLimitExceededResponse(w)
				return
			}

			// Add the remaining requests and reset time to the headers
			w.Header().Add("X-RateLimit-Remaining", strconv.Itoa(limit-int(count)))
			resetTime := time.Now().Add(window).Unix()
			w.Header().Add("X-RateLimit-Reset", strconv.FormatInt(resetTime, 10))
			next.ServeHTTP(w, r)
		})
	}
}
