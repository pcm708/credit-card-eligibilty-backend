package controllers

import (
	"github.com/honestbank/tech-assignment-backend-engineer/handler"
	"github.com/honestbank/tech-assignment-backend-engineer/validator"
	"net/http"

	"github.com/honestbank/tech-assignment-backend-engineer/services"
)

func ProcessData(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		defer req.Body.Close()
		// Sanitizing and validating the request body
		data, status, err := validator.ProcessRequestBody(req)
		if err != nil {
			handler.ErrorHandler(err, status, resp)
			return
		}

		// Performing the decision logic here
		result, status, err := services.DecisionEngine(data)
		if err != nil {
			handler.ErrorHandler(err, status, resp)
			return
		}

		handler.ResponseHandler(result, resp)
		return
	default:
		handler.DefaultResponseHandler(resp)
	}
}
