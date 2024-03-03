package routes

import (
	"github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/controllers"
	"github.com/honestbank/tech-assignment-backend-engineer/middleware"
	"net/http"
	"time"
)

var rdb = middleware.RedisClient()

// SetupRoutes function sets up the routes for the application
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", controllers.AddNumber)

	rateLimiter := middleware.NewRateLimiter(constants.MAX_REQUESTS, time.Duration(constants.RATE_LIMITER_DURATION)*time.Minute)
	mux.Handle("/process", rateLimiter(http.HandlerFunc(controllers.ProcessData)))
	return mux
}
