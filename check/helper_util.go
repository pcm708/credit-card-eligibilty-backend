package check

import (
	"log"
	"strconv"
	"strings"

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
	utils.LogToJSON(data.PhoneNumber, "Area code is not valid", config.Declined)
	return false
}

func IfValidNumberOfCreditCards(data model.RecordData) bool {
	config := utils.GetConfig()
	if data.NumberOfCreditCards <= config.MinNumberOfCC {
		return true
	}
	utils.LogToJSON(data.PhoneNumber, "Number of credit cards are not valid", config.Declined)
	return false
}

func IfValidAge(data model.RecordData) bool {
	config := utils.GetConfig()
	if data.Age >= config.MinAge {
		return true
	}
	utils.LogToJSON(data.PhoneNumber, "Age is not valid", config.Declined)
	return false
}

func IfValidIncome(data model.RecordData) bool {
	config := utils.GetConfig()
	if data.Income >= config.MinIncome {
		return true
	}
	utils.LogToJSON(data.PhoneNumber, "Income is not valid", config.Declined)
	return false
}

func IfCreditRiskScoreLow(data model.RecordData) bool {
	config := utils.GetConfig()
	if config.DesiredCreditRiskScore ==
		calculateCreditRisk(data.Age, data.NumberOfCreditCards) {
		return true
	}
	utils.LogToJSON(data.PhoneNumber, "Credit risk score is not low", config.Declined)
	return false
}

func IfApplicantPoliticallyExposed(data model.RecordData) bool {
	config := utils.GetConfig()
	if data.PoliticallyExposed {
		utils.LogToJSON(data.PhoneNumber, "Applicant is politically exposed", config.Declined)
		return true
	}
	return false
}

func IsNumberPreApproved(data model.RecordData) bool {
	preApprovedNumbers := extractPreApprovedNumbers()
	for _, number := range preApprovedNumbers {
		if number == data.PhoneNumber {
			log.Println("[" + number + "]:: number is preapproved")
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
