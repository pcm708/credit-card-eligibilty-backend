package routes

import (
	"net/http"

	"github.com/honestbank/tech-assignment-backend-engineer/controllers"
	"github.com/honestbank/tech-assignment-backend-engineer/middleware"
)

var rdb = middleware.RedisClient()

// SetupRoutes function sets up the routes for the application
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", controllers.AddNumber)

	// rateLimiter := middleware.NewRateLimiter(constants.MAX_REQUESTS, time.Minute)
	// mux.Handle("/process", rateLimiter(http.HandlerFunc(controllers.ProcessData)))
	mux.HandleFunc("/process", controllers.ProcessData)
	return mux
}
