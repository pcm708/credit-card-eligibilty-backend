package writer

import (
	"encoding/json"
	. "github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/db"
	"github.com/honestbank/tech-assignment-backend-engineer/exceptions"
	"log"
	"os"
	"time"

	"github.com/honestbank/tech-assignment-backend-engineer/model"
)

// Function to store a pre-approved phone number
// Using locks to ensure go-routine safety

type WriterInterface interface {
	StorePreApprovedNumber(phoneNumber string) error
	LogToJSON(phoneNumber string, message string, status string, loglevel string) error
}

type WriterImpl struct{}

// StorePreApprovedNumber stores a pre-approved phone number into the db server
func (c *WriterImpl) StorePreApprovedNumber(phoneNumber string) error {
	err := db.StoreNewNumber(phoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (c *WriterImpl) LogToJSON(phoneNumber string, message string, status string, loglevel string) error {
	Logger(loglevel, phoneNumber+", "+message)

	filePath, exist := os.LookupEnv(LOG_FILE_PATH)
	exceptions.HandleFilePathError(exist, "no path for config file specified")
	// Open log file for appending
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	exceptions.HandleOSOpenFileError(err, "Error reading numbers.txt:")
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
func Logger(logLevel string, msg string) {
	log.Println(logLevel + msg)
}
