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
		// decoder := json.NewDecoder(req.Body)
		defer req.Body.Close()

		var data model.RecordData

		err := json.NewDecoder(req.Body).Decode(&data)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusBadRequest)
			return
		}

		// Perform decision logic here
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
