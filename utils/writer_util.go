package utils

import (
	"log"
	"os"
	"sync"
)

// Function to store a pre-approved phone number

// Using locks to ensure go-routine safety
var preApprovedNumbersMutex sync.Mutex

func StorePreApprovedNumber(phoneNumber string) {
	preApprovedNumbersMutex.Lock()
	defer preApprovedNumbersMutex.Unlock()

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
