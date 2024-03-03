package handler

import (
	"encoding/json"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
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
