package cloud

import (
	"fmt"
	"github.com/honestbank/tech-assignment-backend-engineer/constants"
	"io/ioutil"
	"net/http"
)

// Define the base URL of your Windows machine API
var baseURL string = constants.SERVER_BASE_URL + ":" + constants.SERVER_PORT

// GetDataFromServer retrieves data from a file on the server
func GetDataFromServer() ([]byte, error) {
	resp, err := http.Get(baseURL + "/numbers")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from file on server: %w", err)
	}
	defer resp.Body.Close()
	// Read response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	return data, nil
}

// StoreNewNumber sends a PUT request to the server to store a new phone number
func StoreNewNumber(phoneNumber string) error {
	url := fmt.Sprintf("%s/numbers/add?phone_number=%s", baseURL, phoneNumber)
	// Create a new PUT request with an empty body
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create PUT request: %w", err)
	}
	// Send the PUT request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("PUT request failed: %w", err)
	}
	defer resp.Body.Close()
	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}
