package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/honestbank/tech-assignment-backend-engineer/services"
	"github.com/honestbank/tech-assignment-backend-engineer/validator"
)

// ProcessData is the main function that handles HTTP requests.
// It takes in a http.ResponseWriter and a http.Request as parameters.
func ProcessData(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		defer req.Body.Close()
		data, err := processRequestBody(req)
		if err != nil {
			resp.Header().Set("Content-Type", "application/json")
			resp.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(resp).Encode(model.JsonError{Success: false, Message: err.Error()})
			return
		}

		// Performing the decision logic here
		status := services.DecisionEngine(data)
		// Return the response
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(model.JSONResponse{Status: status})
		return
	default:
		log.Println("error no 404")
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprint(resp, "not found")
	}
}

// processRequestBody processes the request body
// It checks for any validation errors or malformed json,
// If passes all checks returns a RecordData object.
func processRequestBody(req *http.Request) (model.RecordData, error) {
	var data model.RecordData

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(constants.LOG_LEVEL_ERROR + "error reading request body: " + err.Error())
		return data, err
	}

	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	err = json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		log.Println(constants.LOG_LEVEL_ERROR+"error decoding request body: ", string(bodyBytes), " Error: ", err.Error())
		return data, err
	}

	vErr := validator.ValidateRecordData(data)
	if vErr != "" {
		log.Println(constants.LOG_LEVEL_ERROR + "error validating request body: " + vErr)
		return data, errors.New(vErr)
	}

	return data, nil
}
