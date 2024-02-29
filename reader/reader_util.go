package reader

import (
	"encoding/json"
	"fmt"
	"github.com/honestbank/tech-assignment-backend-engineer/cloud"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/honestbank/tech-assignment-backend-engineer/constants"
	"log"
	"os"

	"github.com/honestbank/tech-assignment-backend-engineer/model"
)

// ReaderInterface is an interface for the ReaderImpl.
type ReaderInterface interface {
	// GetConfig returns the config from the config file.
	GetConfig(configFile string) model.Config
}

// ReaderImpl is a struct for the ReaderInterface.
type ReaderImpl struct{}

func (c *ReaderImpl) GetConfig(configFile string) model.Config {
	config, err := readConfigFile(configFile)
	if err != nil {
		log.Fatal("error reading config file:", err)
	}
	return config
}

// readConfigFile reads the config file and returns the config.
func readConfigFile(path string) (model.Config, error) {
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

// ReadTxtFile reads a txt file and returns the content.
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

func ExtractPreApprovedNumbers_Cloud() ([]string, int, error) {
	content, statusCode, err := cloud.GetDataFromServer()
	var numbers []string

	// If the request timed out, log a warning and return an empty list
	if err != nil {
		log.Println(constants.LOG_LEVEL_WARN, err.Error())
		return []string{}, statusCode, err
	}

	if err := json.Unmarshal(content, &numbers); err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}
	return numbers, statusCode, nil
}

// ExtractPreApprovedNumbers_Local extracts the pre-approved numbers from the local.
func ExtractPreApprovedNumbers_Local() ([]string, int, error) {
	content, err := ReadTxtFile(constants.NUMBERS_FILE)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to read numbers file: %w", err)
	}
	lines := strings.Split(string(content), "\n")
	var preApprovedNumbers []string
	preApprovedNumbers = append(preApprovedNumbers, lines...)
	return preApprovedNumbers, http.StatusOK, nil
}
