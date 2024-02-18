package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/honestbank/tech-assignment-backend-engineer/model"
)

// ReadConfig reads configuration data from a JSON file
func ReadConfigFile() (model.Config, error) {
	var config model.Config
	// Read config from file
	configFilePath, exist := os.LookupEnv("CONFIG_PATH")
	if !exist {
		log.Fatal("no path for config file specified")
	}

	configData, err := os.ReadFile(configFilePath)
	if err != nil {
		return config, fmt.Errorf("error reading config file: %v", err)
	}

	// Unmarshal config data into Config struct
	err = json.Unmarshal(configData, &config)
	if err != nil {
		return config, fmt.Errorf("error parsing config data: %v", err)
	}

	return config, nil
}

// reading
func IsPreApproved(phoneNumber string) bool {

	preApproved_Numbers, exist := os.LookupEnv("PRE_APPROVED_NUMBERS_FILE_PATH")
	if !exist {
		log.Fatal("no path for config file specified")
	}

	file, err := os.Open(preApproved_Numbers)
	if err != nil {
		return false // File doesn't exist or cannot be opened, treat as not pre-approved
	}
	defer file.Close()

	// Check if the provided phone number matches any pre-approved number in the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == phoneNumber {
			return true // Phone number found in pre-approved list
		}
	}

	return false // Phone number not found in pre-approved list
}
