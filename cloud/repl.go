package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/exceptions"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// Define the base URL of your Windows machine API
var baseURL string = constants.SERVER_BASE_URL + ":" + constants.SERVER_PORT
var ctx = context.Background()
var rdb = RedisClient()

func GetDataFromServer() ([]string, int, error) {
	// Fetch data from Redis
	result, err := rdb.Get(ctx, "numbers").Result()
	if err != nil {
		exceptions.HandleRedisServerError(err)
		//fallback to return data from cloud
		return GetDataFromCloudServer()
	}

	preApprovedNumbers := strings.Split(result, "\n")
	return preApprovedNumbers, http.StatusOK, nil

	//ExtractPreApprovedNumbers_Cloud()
}

func GetDataFromCloudServer() ([]string, int, error) {
	client := http.Client{
		Timeout: time.Second * 5, // Timeout after 10 seconds
	}

	resp, err := client.Get(baseURL + "/numbers")
	exceptions.HandleError(err, http.StatusServiceUnavailable)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	exceptions.HandleError(err, http.StatusInternalServerError)

	var numbers []string

	// If the request timed out, log a warning and return an empty list
	if err != nil {
		exceptions.HandleError(err, http.StatusInternalServerError)
		return []string{}, http.StatusOK, err
	}

	err = json.Unmarshal(data, &numbers)
	exceptions.HandleError(err, http.StatusInternalServerError)

	return numbers, http.StatusOK, nil
}

// StoreNewNumber sends a PUT request to the server to store a new phone number
func StoreNewNumber(phoneNumber string) (string, int, error) {
	url := fmt.Sprintf("%s/numbers/add?phone_number=%s", baseURL, phoneNumber)
	// Create a new PUT request with an empty body
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return err.Error(), http.StatusInternalServerError, err
	}
	//exceptions.HandleError(err, http.StatusInternalServerError)

	// Send the PUT request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err.Error(), http.StatusInternalServerError, err
	}
	defer resp.Body.Close()
	//exceptions.HandleError(err, http.StatusInternalServerError)

	// Check the response status
	//exceptions.HandleStatusCodeNot200Error(resp, err)
	if resp.StatusCode != http.StatusOK {
		return "unexpected error code", resp.StatusCode, err
	}
	log.Println(constants.LOG_LEVEL_INFO + "Successfully stored new number on the database")
	return "", http.StatusOK, nil
}
