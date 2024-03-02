package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/honestbank/tech-assignment-backend-engineer/exceptions"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/honestbank/tech-assignment-backend-engineer/validator"
	"log"
	"net/http"

	"github.com/honestbank/tech-assignment-backend-engineer/services"
)

// ProcessData is the main function that handles HTTP requests.

func ProcessData(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		defer req.Body.Close()
		data, err := validator.ProcessRequestBody(req)
		if err != nil {
			exceptions.HandleValidationOrBadJsonInputError(err, resp)
			return
		}

		// Performing the decision logic here
		result, status, err := services.DecisionEngine(data)
		if err != nil {
			exceptions.HandleDecisionServiceError(err, status, resp)
			return
		}

		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(model.JSONResponse{Status: result})
		return
	default:
		log.Println("error no 404")
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprint(resp, "not found")
	}
}
