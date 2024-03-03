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
		data, status, err, uid := validator.ProcessRequestBody(req)
		if err != nil {
			handler.ErrorHandler(err, status, resp, uid)
			return
		}

		// Performing the decision logic here
		result, status, err, uid := services.DecisionEngine(data)
		if err != nil {
			handler.ErrorHandler(err, status, resp, uid)
			return
		}

		handler.ResponseHandler(result, resp, uid)
		return
	default:
		handler.DefaultResponseHandler(resp)
	}
}
