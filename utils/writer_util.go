package utils

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/honestbank/tech-assignment-backend-engineer/model"
)

// Function to store a pre-approved phone number

// Using locks to ensure go-routine safety
var mutex sync.Mutex

func StorePreApprovedNumber(phoneNumber string) {
	mutex.Lock()
	defer mutex.Unlock()

	preApproved_Numbers, exist := os.LookupEnv("PRE_APPROVED_NUMBERS_FILE_PATH")
	if !exist {
		log.Fatal("no path for config file specified")
	}

	file, err := os.OpenFile(preApproved_Numbers, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error reading preApproved_Numbers.txt:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(phoneNumber + "\n"); err != nil {
		log.Println("Error writing to preApproved_Numbers.txt:", err)
		return
	}
}

func LogToJSON(phoneNumber string, message string, status string, loglevel string) error {
	mutex.Lock()
	defer mutex.Unlock()

	Logger(phoneNumber, message, loglevel)

	filePath, exist := os.LookupEnv("LOG_FILE_PATH")
	if !exist {
		log.Fatal("no path for config file specified")
	}

	// Open log file for appending
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decode the existing JSON array if the file isn't empty
	var logs []model.LogEntry
	stat, _ := file.Stat()
	if stat.Size() > 0 {
		if err := json.NewDecoder(file).Decode(&logs); err != nil {
			return err
		}
	}

	// Append the new log entry
	logs = append(logs, model.LogEntry{
		PhoneNumber: phoneNumber,
		Status:      status,
		Message:     message,
	})

	// Rewind the file pointer to the beginning
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

	// Encode the logs to JSON and write to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(logs); err != nil {
		return err
	}

	return nil
}

func Logger(phoneNumber string, message string, logLevel string) {
	log.Println(logLevel + ":: " + phoneNumber + "," + message)
}
