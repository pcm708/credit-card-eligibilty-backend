package utils

import (
	"strconv"
)

func CalculateCreditRisk(age, numberOfCreditCard int) string {
	sum := age + numberOfCreditCard
	mod := sum % 3
	if mod == 0 {
		return "LOW"
	}
	if mod == 1 {
		return "MEDIUM"
	}
	return "HIGH"
}

func IsValidAreaCode(phoneNumber string, allowedAreaCodes []int) bool {
	// Extract the area code (first digit of phone number)
	areaCodeStr := string(phoneNumber[0])
	areaCode, err := strconv.Atoi(areaCodeStr)
	if err != nil {
		// Handle error, perhaps log it
		return false
	}

	// Check if the area code is in the list of allowed area codes
	for _, code := range allowedAreaCodes {
		if areaCode == code {
			return true
		}
	}
	return false
}
