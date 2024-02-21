package utils

import (
	. "github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"strconv"
)

// calculateCreditRisk calculates the credit risk based on age and number of credit cards.
// It returns a string indicating the risk level.
func calculateCreditRisk(age, numberOfCreditCard int) string {
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

// IsValidAreaCode checks if the area code of the phone number is valid.
// It returns a boolean indicating the validity.
func IsValidAreaCode(data model.RecordData, config model.Config) bool {
	areaCodeStr := string(data.PhoneNumber[0])
	areaCode, err := strconv.Atoi(areaCodeStr)
	if err != nil {
		LogToJSON(data.PhoneNumber, INVALID_AREA_CODE, DECLINED, LOG_LEVEL_ERROR)
		return false
	}
	for _, code := range config.AllowedAreaCodes {
		if areaCode == code {
			return true
		}
	}
	LogToJSON(data.PhoneNumber, INVALID_AREA_CODE, DECLINED, LOG_LEVEL_WARN)
	return false
}

// IfValidNumberOfCreditCards checks if the number of credit cards is valid.
// It returns a boolean indicating the validity.
func IfValidNumberOfCreditCards(data model.RecordData, config model.Config) bool {
	if data.NumberOfCreditCards != nil && *data.NumberOfCreditCards <= config.MinNumberOfCC {
		return true
	}
	LogToJSON(data.PhoneNumber, INVALID_CC_NUMBER, DECLINED, LOG_LEVEL_WARN)
	return false
}

// IfValidAge checks if the age is valid.
// It returns a boolean indicating the validity
func IfValidAge(data model.RecordData, config model.Config) bool {
	if data.Age >= config.MinAge {
		return true
	}
	LogToJSON(data.PhoneNumber, INVALID_AGE, DECLINED, LOG_LEVEL_WARN)
	return false
}

// IfValidIncome checks if the income is valid.
// It returns a boolean indicating the validity.
func IfValidIncome(data model.RecordData, config model.Config) bool {
	if data.Income >= config.MinIncome {
		return true
	}
	LogToJSON(data.PhoneNumber, INVALID_INCOME, DECLINED, LOG_LEVEL_WARN)
	return false
}

// IfCreditRiskScoreLow checks if the credit risk score is low.
// It returns a boolean indicating the result.
func IfCreditRiskScoreLow(data model.RecordData, config model.Config) bool {
	if data.NumberOfCreditCards != nil &&
		config.DesiredCreditRiskScore ==
			calculateCreditRisk(data.Age, *data.NumberOfCreditCards) {
		return true
	}
	LogToJSON(data.PhoneNumber, INVALID_CREDIT_RISK_SCORE, DECLINED, LOG_LEVEL_WARN)
	return false
}

// IfApplicantPoliticallyExposed checks if the applicant is politically exposed.
// It returns a boolean indicating the result.
func IfApplicantPoliticallyExposed(data model.RecordData) bool {
	if data.PoliticallyExposed != nil && *data.PoliticallyExposed {
		LogToJSON(data.PhoneNumber, POLITICALLY_EXPOSED, DECLINED, LOG_LEVEL_WARN)
		return true
	}
	return false
}
