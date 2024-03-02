package middleware

import (
	"context"
	"encoding/json"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
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
			if err != nil {
				// handle error
			}
			// Increment the request count
			count, err := rdb.Incr(ctx, ip).Result()
			if err != nil {
				// Create a JsonError instance
				err := model.JsonError{
					Success: false,
					Message: "Error contacting redis db",
				}
				// Convert the JsonError instance into JSON
				errJson, _ := json.Marshal(err)
				// Set the Content-Type header to application/json
				w.Header().Set("Content-Type", "application/json")
				// Write the JSON error to the response with a 429 status code
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(errJson)
				return
			}

			if count == 1 {
				// If this is the first request, set the key to expire after the time window
				rdb.Expire(ctx, ip, window)
			}

			if count > int64(limit) {
				// Create a JsonError instance
				err := model.JsonError{
					Success: false,
					Message: "Rate limit exceeded",
				}
				// Convert the JsonError instance into JSON
				errJson, _ := json.Marshal(err)
				// Set the Content-Type header to application/json
				w.Header().Set("Content-Type", "application/json")
				// Write the JSON error to the response with a 429 status code
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write(errJson)
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
