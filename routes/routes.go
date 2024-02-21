package routes

import (
	"net/http"

	"github.com/honestbank/tech-assignment-backend-engineer/controllers"
)

// SetupRoutes function sets up the routes for the application
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	// Adding a route for handling our post request
	// When "/process" is hit, the ProcessData function from the controllers package will be called
	mux.HandleFunc("/process", controllers.ProcessData)

	return mux
}
