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

func GetConfig(path string) model.Config {
	config, err := readConfigFile(path)
	if err != nil {
		log.Fatal("error reading config file:", err)
	}
	return config
}

// reading
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
