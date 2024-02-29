package cloud

import (
	"fmt"
	"github.com/honestbank/tech-assignment-backend-engineer/constants"
	"io/ioutil"
	"net/http"
	"time"
)

// Define the base URL of your Windows machine API
var baseURL string = constants.SERVER_BASE_URL + ":" + constants.SERVER_PORT

func GetDataFromServer() ([]byte, int, error) {
	client := http.Client{
		Timeout: time.Second * 5, // Timeout after 10 seconds
	}

	resp, err := client.Get(baseURL + "/numbers")
	if err != nil {
		return nil, http.StatusServiceUnavailable, fmt.Errorf("failed to fetch data from file on server: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to read response body: %w", err)
	}
	return data, http.StatusOK, nil
}

// GetDataFromServer retrieves data from a file on the server
//func GetDataFromServer() ([]byte, error) {
//	resp, err := http.Get(baseURL + "/numbers")
//	if err != nil {
//		return nil, fmt.Errorf("failed to fetch data from file on server: %w", err)
//	}
//	defer resp.Body.Close()
//	// Read response body
//	data, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return nil, fmt.Errorf("failed to read response body: %w", err)
//	}
//	return data, nil
//}

// StoreNewNumber sends a PUT request to the server to store a new phone number
func StoreNewNumber(phoneNumber string) (string, int, error) {
	url := fmt.Sprintf("%s/numbers/add?phone_number=%s", baseURL, phoneNumber)
	// Create a new PUT request with an empty body
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return err.Error(), http.StatusInternalServerError, err
	}
	// Send the PUT request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err.Error(), http.StatusInternalServerError, err
	}
	defer resp.Body.Close()
	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return "unexpected error code", resp.StatusCode, err
	}
	return "", http.StatusOK, nil
}
