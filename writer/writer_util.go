package writer

import (
	"encoding/json"
	"github.com/honestbank/tech-assignment-backend-engineer/cloud"
	. "github.com/honestbank/tech-assignment-backend-engineer/constants"
	"log"
	"os"
	"sync"
	"time"

	"github.com/honestbank/tech-assignment-backend-engineer/model"
)

// Function to store a pre-approved phone number
// Using locks to ensure go-routine safety
var mutex sync.Mutex

type WriterInterface interface {
	// StorePreApprovedNumber stores a pre-approved phone number in numbers.txt on a server or local
	StorePreApprovedNumber(phoneNumber string)
	// LogToJSON logs a message to a JSON file
	LogToJSON(phoneNumber string, message string, status string, loglevel string) error
}

type WriterImpl struct{}

func (c *WriterImpl) StorePreApprovedNumber(phoneNumber string) {
	//storePreApprovedNumber_Cloud(phoneNumber)
	storePreApprovedNumber_Local(phoneNumber)
}

// StorePreApprovedNumber stores a pre-approved phone number into the cloud server
func storePreApprovedNumber_Cloud(phoneNumber string) {
	mutex.Lock()
	defer mutex.Unlock()
	cloud.StoreNewNumber(phoneNumber)
}

// StorePreApprovedNumber_Local stores a pre-approved phone number in numbers.txt
func storePreApprovedNumber_Local(phoneNumber string) {
	mutex.Lock()
	defer mutex.Unlock()

	preApproved_Numbers, exist := os.LookupEnv(NUMBERS_FILE)
	if !exist {
		log.Fatal("no path for config file specified")
	}

	file, err := os.OpenFile(preApproved_Numbers, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error reading numbers.txt:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(phoneNumber + "\n"); err != nil {
		log.Println("Error writing to numbers.txt:", err)
		return
	}
}

func (c *WriterImpl) LogToJSON(phoneNumber string, message string, status string, loglevel string) error {
	Logger(phoneNumber, message, loglevel)

	filePath, exist := os.LookupEnv(LOG_FILE_PATH)
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
		Timestamp:   time.Now().Format("2006-01-02 15:04:05"),
	})
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

// Logger Log to console
func Logger(phoneNumber string, message string, logLevel string) {
	log.Println(logLevel + phoneNumber + " , " + message)
}
