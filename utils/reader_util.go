package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/honestbank/tech-assignment-backend-engineer/model"
)

// ReadConfig reads configuration data from a JSON file
func readConfigFile() (model.Config, error) {
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

func GetConfig() model.Config {
	config, err := readConfigFile()
	if err != nil {
		log.Fatal("error reading config file:", err)
	}
	return config
}

// reading
func ReadTxtFile() ([]byte, error) {
	filePath, exist := os.LookupEnv("PRE_APPROVED_NUMBERS_FILE_PATH")
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
