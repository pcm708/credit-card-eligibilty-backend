package check

import (
	"strconv"
	"strings"

	"github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/honestbank/tech-assignment-backend-engineer/utils"
)

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

func IsValidAreaCode(data model.RecordData) bool {
	config := utils.GetConfig()
	areaCodeStr := string(data.PhoneNumber[0])
	areaCode, err := strconv.Atoi(areaCodeStr)
	if err != nil {
		return false
	}
	for _, code := range config.AllowedAreaCodes {
		if areaCode == code {
			return true
		}
	}
	utils.LogToJSON(data.PhoneNumber, constants.INVALID_AREA_CODE,
		constants.DECLINED, constants.LOG_LEVEL_WARN)
	return false
}

func IfValidNumberOfCreditCards(data model.RecordData) bool {
	config := utils.GetConfig()
	if data.NumberOfCreditCards <= config.MinNumberOfCC {
		return true
	}
	utils.LogToJSON(data.PhoneNumber, constants.INVALID_CC_NUMBER,
		constants.DECLINED, constants.LOG_LEVEL_WARN)
	return false
}

func IfValidAge(data model.RecordData) bool {
	config := utils.GetConfig()
	if data.Age >= config.MinAge {
		return true
	}
	utils.LogToJSON(data.PhoneNumber, constants.INVALID_AGE,
		constants.DECLINED, constants.LOG_LEVEL_WARN)
	return false
}

func IfValidIncome(data model.RecordData) bool {
	config := utils.GetConfig()
	if data.Income >= config.MinIncome {
		return true
	}
	utils.LogToJSON(data.PhoneNumber, constants.INVALID_INCOME,
		constants.DECLINED, constants.LOG_LEVEL_WARN)
	return false
}

func IfCreditRiskScoreLow(data model.RecordData) bool {
	config := utils.GetConfig()
	if config.DesiredCreditRiskScore ==
		calculateCreditRisk(data.Age, data.NumberOfCreditCards) {
		return true
	}
	utils.LogToJSON(data.PhoneNumber, constants.INVALID_CREDIT_RISK_SCORE,
		constants.DECLINED, constants.LOG_LEVEL_WARN)
	return false
}

func IfApplicantPoliticallyExposed(data model.RecordData) bool {
	if data.PoliticallyExposed {
		utils.LogToJSON(data.PhoneNumber, constants.POLITICALLY_EXPOSED,
			constants.DECLINED, constants.LOG_LEVEL_WARN)
		return true
	}
	return false
}

func IsNumberPreApproved(data model.RecordData) bool {
	preApprovedNumbers := extractPreApprovedNumbers()
	for _, number := range preApprovedNumbers {
		if number == data.PhoneNumber {
			utils.Logger(number, constants.PREAPPROVED_NUMBER,
				constants.LOG_LEVEL_INFO)
			return true
		}
	}
	return false
}

func extractPreApprovedNumbers() []string {
	content, _ := utils.ReadTxtFile()
	lines := strings.Split(string(content), "\n")
	var preApprovedNumbers []string
	preApprovedNumbers = append(preApprovedNumbers, lines...)
	return preApprovedNumbers
}
