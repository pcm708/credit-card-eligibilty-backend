package reader

import (
	"encoding/json"
	"fmt"
	"github.com/honestbank/tech-assignment-backend-engineer/cloud"
	"github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/exceptions"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

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
	exceptions.HandleFileReadingError(err, "error reading config file:")
	return config
}

// readConfigFile reads the config file and returns the config.
func readConfigFile(path string) (model.Config, error) {
	var config model.Config
	// Read config from file
	configFilePath, exist := os.LookupEnv(path)
	exceptions.HandleFilePathError(exist, "no path for config file specified")

	configData, err := os.ReadFile(configFilePath)
	exceptions.HandleFileReadingError(err, "error reading config file: %v")

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
	exceptions.HandleFilePathError(exist, "no path for config file specified")

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

// ExtractPreApprovedNumbers_Local extracts the pre-approved numbers from the local.
func ExtractPreApprovedNumbers_Local() ([]string, int, error) {
	content, err := ReadTxtFile(constants.NUMBERS_FILE)
	exceptions.HandleError(err, http.StatusInternalServerError)
	lines := strings.Split(string(content), "\n")
	var preApprovedNumbers []string
	preApprovedNumbers = append(preApprovedNumbers, lines...)
	return preApprovedNumbers, http.StatusOK, nil
}

func ExtractPreApprovedNumbers_Cloud() ([]string, int, error) {
	return cloud.GetDataFromServer()
}
