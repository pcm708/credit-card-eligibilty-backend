package validator

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"regexp"

	. "github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
)

// processRequestBody processes the request body
// It checks for any validation errors or malformed json,
// If passes all checks returns a RecordData object.
func ProcessRequestBody(req *http.Request) (model.RecordData, error) {
	var data model.RecordData

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(LOG_LEVEL_ERROR + "error reading request body: " + err.Error())
		return data, err
	}

	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	err = json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		log.Println(LOG_LEVEL_ERROR, "error decoding request body: ", err.Error(), "\n", string(bodyBytes))
		return data, err
	}

	vErr := ValidateRecordData(data)
	if vErr != "" {
		log.Println(LOG_LEVEL_ERROR + "error validating request body: " + vErr + "\n" + string(bodyBytes))
		return data, errors.New(vErr)
	}

	return data, nil
}

// ValidateRecordData validates the data in a RecordData object.
// It checks if any fields are blank, if the income is negative, if the number of credit cards is negative,
// if the age is negative, and if the phone number is too long.
// It returns a string indicating the validation error, or an empty string if there are no errors.

func ValidateRecordData(data model.RecordData) string {
	// Check if any fields are blank
	if data.Income == 0 || data.NumberOfCreditCards == 0 ||
		data.Age == 0 || data.PoliticallyExposed == nil ||
		data.PhoneNumber == "" {
		return NO_FIELDS_BLANK
	}
	// Check if the income is negative
	if data.Income < 0 {
		return INCOME_NEGATIVE
	}
	// Check if the number of credit cards is negative
	if data.NumberOfCreditCards < 0 {
		return CC_NEGATIVE
	}
	// Check if the age is negative
	if data.Age < 0 {
		return AGE_NEGATIVE
	}
	// Check if the phone number is equal to 12 and should only contain numbers and '-'
	// Also, Check if the phone number is equal to 10 and should only contain numbers
	// If not, return an error
	phoneNumberPattern := `^(?:[0-9]{3}-[0-9]{3}-[0-9]{4}|[0-9]{10})$`
	match, _ := regexp.MatchString(phoneNumberPattern, data.PhoneNumber)
	if !match {
		return INVALID_PHONE
	}
	// If there are no validation errors, return an empty string
	return ""
}
