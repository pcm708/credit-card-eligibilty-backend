package check

import (
	"strconv"

	. "github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/honestbank/tech-assignment-backend-engineer/reader"
	"github.com/honestbank/tech-assignment-backend-engineer/writer"
)

type CheckInterface interface {
	// IsNumberPreApproved checks if the phone number is pre-approved.
	IsNumberPreApproved(data model.RecordData) bool

	// IsValidAreaCode checks if the area code of the phone number is valid.
	// It returns a boolean indicating the validity.
	IsValidAreaCode(data model.RecordData, config model.Config) bool

	// IfValidNumberOfCreditCards checks if the number of credit cards is valid.
	// It returns a boolean indicating the validity.
	IfValidNumberOfCreditCards(data model.RecordData, config model.Config) bool

	// IfValidAge checks if the age is valid.
	// It returns a boolean indicating the validity
	IfValidAge(data model.RecordData, config model.Config) bool

	// IfValidIncome checks if the income is valid.
	// It returns a boolean indicating the validity.
	IfValidIncome(data model.RecordData, config model.Config) bool

	// IfCreditRiskScoreLow checks if the credit risk score is low.
	// It returns a boolean indicating the result.
	IfCreditRiskScoreLow(data model.RecordData, config model.Config) bool

	// IfApplicantPoliticallyExposed checks if the applicant is politically exposed.
	// It returns a boolean indicating the result.
	IfApplicantPoliticallyExposed(data model.RecordData) bool
}

var Writer writer.WriterInterface = &writer.WriterImpl{}

type CheckImpl struct{}

func (c *CheckImpl) IsNumberPreApproved(data model.RecordData) bool {
	preApprovedNumbers, _ := reader.ExtractPreApprovedNumbers()
	for _, number := range preApprovedNumbers {
		if number == data.PhoneNumber {
			Writer.LogToJSON(data.PhoneNumber, PREAPPROVED_NUMBER, APPROVED, LOG_LEVEL_INFO)
			return true
		}
	}
	return false
}

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

func (c *CheckImpl) IsValidAreaCode(data model.RecordData, config model.Config) bool {
	areaCodeStr := string(data.PhoneNumber[0])
	areaCode, err := strconv.Atoi(areaCodeStr)
	if err != nil {
		Writer.LogToJSON(data.PhoneNumber, INVALID_AREA_CODE, DECLINED, LOG_LEVEL_ERROR)
		return false
	}
	for _, code := range config.AllowedAreaCodes {
		if areaCode == code {
			return true
		}
	}
	Writer.LogToJSON(data.PhoneNumber, INVALID_AREA_CODE, DECLINED, LOG_LEVEL_WARN)
	return false
}

func (c *CheckImpl) IfValidNumberOfCreditCards(data model.RecordData, config model.Config) bool {
	if data.NumberOfCreditCards != nil && *data.NumberOfCreditCards <= config.MinNumberOfCC {
		return true
	}
	Writer.LogToJSON(data.PhoneNumber, INVALID_CC_NUMBER, DECLINED, LOG_LEVEL_WARN)
	return false
}

func (c *CheckImpl) IfValidAge(data model.RecordData, config model.Config) bool {
	if data.Age >= config.MinAge {
		return true
	}
	Writer.LogToJSON(data.PhoneNumber, INVALID_AGE, DECLINED, LOG_LEVEL_WARN)
	return false
}

func (c *CheckImpl) IfValidIncome(data model.RecordData, config model.Config) bool {
	if data.Income >= config.MinIncome {
		return true
	}
	Writer.LogToJSON(data.PhoneNumber, INVALID_INCOME, DECLINED, LOG_LEVEL_WARN)
	return false
}

func (c *CheckImpl) IfCreditRiskScoreLow(data model.RecordData, config model.Config) bool {
	if data.NumberOfCreditCards != nil &&
		config.DesiredCreditRiskScore == calculateCreditRisk(data.Age, *data.NumberOfCreditCards) {
		return true
	}
	Writer.LogToJSON(data.PhoneNumber, INVALID_CREDIT_RISK_SCORE, DECLINED, LOG_LEVEL_WARN)
	return false
}

func (c *CheckImpl) IfApplicantPoliticallyExposed(data model.RecordData) bool {
	if data.PoliticallyExposed != nil && *data.PoliticallyExposed {
		Writer.LogToJSON(data.PhoneNumber, POLITICALLY_EXPOSED, DECLINED, LOG_LEVEL_WARN)
		return true
	}
	return false
}
