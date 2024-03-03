package handler

import (
	"encoding/json"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"net/http"
)

func ErrorHandler(err error, statusCode int, resp http.ResponseWriter) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(statusCode)
	json.NewEncoder(resp).Encode(model.JsonError{Success: false, Message: err.Error()})
}

func ResponseHandler(result string, resp http.ResponseWriter) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(model.JSONResponse{Status: result})
}

func DefaultResponseHandler(resp http.ResponseWriter) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusNotFound)
}
