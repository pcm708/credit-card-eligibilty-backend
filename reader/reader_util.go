package reader

import (
	"encoding/json"
	"fmt"
	. "github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

// ExtractPreApprovedNumbers extracts the pre-approved numbers from the numbers file.
func ExtractPreApprovedNumbers() ([]string, error) {
	content, err := ReadTxtFile(NUMBERS_FILE)
	if err != nil {
		return nil, fmt.Errorf("failed to read numbers file: %w", err)
	}
	lines := strings.Split(string(content), "\n")
	var preApprovedNumbers []string
	preApprovedNumbers = append(preApprovedNumbers, lines...)
	return preApprovedNumbers, nil
}
