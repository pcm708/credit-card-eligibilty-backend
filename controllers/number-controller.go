package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/honestbank/tech-assignment-backend-engineer/db"
	"log"
	"net/http"
)

func AddNumber(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		// Define a slice to hold the numbers
		var numbers []string

		// Decode the JSON body into the numbers slice
		err := json.NewDecoder(req.Body).Decode(&numbers)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusBadRequest)
			return
		}

		// Iterate over the numbers and store each one in the DB
		for _, number := range numbers {
			err := db.StoreNewNumber(number)
			if err != nil {
				http.Error(resp, fmt.Sprintf("Failed to store number %s: %v", number, err), http.StatusInternalServerError)
				return
			}
		}
		// If everything goes well, respond with a success message
		resp.WriteHeader(http.StatusOK)
		fmt.Fprint(resp, "Numbers stored successfully")
		return
	default:
		log.Println("error no 404")
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprint(resp, "not found")
	}
}
