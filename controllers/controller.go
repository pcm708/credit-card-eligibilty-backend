package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/honestbank/tech-assignment-backend-engineer/services"
	"github.com/honestbank/tech-assignment-backend-engineer/validator"
	"log"
	"net/http"
)

// ProcessData is the main function that handles HTTP requests.
// It takes in a http.ResponseWriter and a http.Request as parameters.
func ProcessData(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		// Ensure the request body is closed after the function returns
		defer req.Body.Close()
		var data model.RecordData

		// Decode the request body into the RecordData struct
		err := json.NewDecoder(req.Body).Decode(&data)
		if err != nil {
			log.Println(constants.LOG_LEVEL_ERROR + ":: error decoding request body")
			resp.Header().Set("Content-Type", "application/json")
			resp.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(resp).Encode(model.JsonFor4XX{Success: false, Message: err.Error()})
			return
		}

		// Validate the request body
		validationError := validator.ValidateRecordData(data)
		if validationError != "" {
			log.Println(constants.LOG_LEVEL_ERROR + ":: error validating request body")
			resp.Header().Set("Content-Type", "application/json")
			resp.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(resp).Encode(model.JsonFor4XX{Success: false, Message: validationError})
			return
		}

		// Performing the decision logic here
		status := services.DecisionEngine(data)

		// Return the response
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(model.JSONResponse{Status: status})
	default:
		log.Println("error no 404")
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprint(resp, "not found")
	}
}
