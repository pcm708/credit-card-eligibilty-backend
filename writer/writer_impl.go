package writer

import (
	"encoding/json"
	. "github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/db"
	"log"
	"os"
	"time"

	"github.com/honestbank/tech-assignment-backend-engineer/model"
)

type WriterImpl struct{}

// StorePreApprovedNumber stores a pre-approved phone number into the db server
func (c *WriterImpl) StorePreApprovedNumber(phoneNumber string) error {
	err := db.StoreNewNumber(phoneNumber)
	if err != nil {
		return err
	}
	return nil
}

// LogToJSON logs the decision to a JSON file
func (c *WriterImpl) LogToJSON(phoneNumber string, message string, status string, loglevel string) error {
	log.Println(loglevel, phoneNumber+", "+message)

	filePath := LOG_FILE_PATH
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
