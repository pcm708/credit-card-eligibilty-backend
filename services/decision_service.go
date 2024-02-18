package services

import (
	"fmt"

	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/honestbank/tech-assignment-backend-engineer/utils"
)

func DecisionEngine(data model.RecordData) string {
	configFile, _ := utils.ReadConfigFile()

	if utils.IsPreApproved(data.PhoneNumber) {
		return "approved,as phone number is already stored"
	}

	//negative checks
	areaCodeValid := utils.IsValidAreaCode(data.PhoneNumber, configFile.AllowedAreaCodes)
	if !areaCodeValid {
		fmt.Println("Area code of the phone number is not valid.")
		return "declined"
	}

	if data.NumberOfCreditCards > configFile.MinNumberOfCC {
		fmt.Println("Number of credit cards exceeds the minimum allowed.")
		return "declined"
	}

	if data.Age < configFile.MinAge {
		fmt.Println("Age is below the minimum allowed.")
		return "declined"
	}

	if data.Income <= configFile.MinIncome {
		fmt.Println("Income is below or equal to the minimum allowed.")
		return "declined"
	}

	creditRisk := utils.CalculateCreditRisk(data.Age, data.NumberOfCreditCards)
	if creditRisk != configFile.DesiredCreditRiskScore {
		fmt.Println("Credit risk score is not at the desired level.")
		return "declined"
	}

	if data.PoliticallyExposed {
		fmt.Println("Applicant is politically exposed.")
		return "declined"
	}

	// happy flow
	utils.StorePreApprovedNumber(data.PhoneNumber)
	return "approved"
}
