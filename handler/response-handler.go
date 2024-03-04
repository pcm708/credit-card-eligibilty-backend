package handler

import (
	"encoding/json"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"log"
	"net/http"
)

func GetHTTPStatusCode(statusCode int) int {
	switch statusCode {
	case 200:
		return http.StatusOK
	case 400:
		return http.StatusBadRequest
	case 401:
		return http.StatusUnauthorized
	case 403:
		return http.StatusForbidden
	case 404:
		return http.StatusNotFound
	case 500:
		return http.StatusInternalServerError
	case 503:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}

func ErrorHandler(err error, statusCode int, resp http.ResponseWriter, uid string) {
	resp.Header().Set("X-Request-ID", uid)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(statusCode)
	json.NewEncoder(resp).Encode(model.JsonError{Success: false,
		Message: err.Error()})
}

func ResponseHandler(result string, resp http.ResponseWriter, uid string) {
	resp.Header().Set("X-Request-ID", uid)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(model.JSONResponse{Status: result})
}

func DefaultResponseHandler(resp http.ResponseWriter) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusNotFound)
}

func RedisConnectionErrorResponse(err error, w http.ResponseWriter) {
	log.Println(err.Error())
	// Create a JsonError instance
	errResp := model.JsonError{
		Success: false,
		Message: "redis host seems down",
	}
	// Convert the JsonError instance into JSON
	errJson, _ := json.Marshal(errResp)
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON error to the response with a 429 status code
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Write(errJson)
}

func RateLimitExceededResponse(w http.ResponseWriter) {
	// Create a JsonError instance
	errResp := model.JsonError{
		Success: false,
		Message: "Rate limit exceeded",
	}
	// Convert the JsonError instance into JSON
	errJson, _ := json.Marshal(errResp)
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON error to the response with a 429 status code
	w.WriteHeader(http.StatusTooManyRequests)
	w.Write(errJson)
}
