package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/honestbank/tech-assignment-backend-engineer/services"
)

func ProcessData(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		defer req.Body.Close()

		var data model.RecordData

		if err := decoder.Decode(&data); err != nil {
			resp.Header().Set("Content-Type", "application/json")
			resp.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(resp).Encode(model.JSONFor400{Success: false, Message: err.Error()})
			return
		}

		if err := data.Validate(); err != nil {
			resp.Header().Set("Content-Type", "application/json")
			resp.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(resp).Encode(model.JSONFor400{Success: false, Message: err.Error()})
			return
		}

		// Perform validation and decision logic here
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
