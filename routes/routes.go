package routes

import (
	"net/http"

	"github.com/honestbank/tech-assignment-backend-engineer/controllers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	// route for handling our post request
	mux.HandleFunc("/process", controllers.ProcessData)

	return mux
}
