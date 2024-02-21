package utils

import (
	"encoding/json"
	"fmt"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"io/ioutil"
	"log"
	"os"
)

// readConfigFile reads configuration data from a JSON file.
// It takes in a string representing the path to the file and returns a Config object and an error.
func ReadConfigFile(path string) (model.Config, error) {
	var config model.Config
	// Read config from file
	configFilePath, exist := os.LookupEnv(path)
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

// GetConfig is a wrapper function for readConfigFile.
// It takes in a string representing the path to the file and returns a Config object.
func GetConfig(path string) model.Config {
	config, err := ReadConfigFile(path)
	if err != nil {
		log.Fatal("error reading config file:", err)
	}
	return config
}

// ReadTxtFile reads a text file and returns its content as a byte slice and an error.
// It takes in a string representing the path to the file.
func ReadTxtFile(path string) ([]byte, error) {
	filePath, exist := os.LookupEnv(path)
	if !exist {
		log.Fatal("no path for config file specified")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}

//// ExtractPreApprovedNumbers extracts pre-approved numbers from a file.
//// It returns a slice of strings containing the numbers and an error if any.
//func ExtractPreApprovedNumbers() ([]string, error) {
//	content, err := ReadTxtFile(NUMBERS_FILE)
//	if err != nil {
//		return nil, fmt.Errorf("failed to read numbers file: %w", err)
//	}
//	lines := strings.Split(string(content), "\n")
//	var preApprovedNumbers []string
//	preApprovedNumbers = append(preApprovedNumbers, lines...)
//	return preApprovedNumbers, nil
//}
