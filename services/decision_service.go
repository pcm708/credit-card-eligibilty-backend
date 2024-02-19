package services

import (
	"github.com/honestbank/tech-assignment-backend-engineer/check"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/honestbank/tech-assignment-backend-engineer/utils"
)

func DecisionEngine(data model.RecordData) string {
	config := utils.GetConfig()
	return isApplicantEligible(data, config)
}

func isApplicantEligible(data model.RecordData, config model.Config) string {
	// check is number is preapproved or not
	if check.IsNumberPreApproved(data) {
		return config.Approved
	}

	if !check.IfApplicantPoliticallyExposed(data) &&
		check.IfValidAge(data) &&
		check.IfValidIncome(data) &&
		check.IfValidNumberOfCreditCards(data) &&
		check.IsValidAreaCode(data) &&
		check.IfCreditRiskScoreLow(data) {

		utils.StorePreApprovedNumber(data.PhoneNumber)
		utils.LogToJSON(data.PhoneNumber, "number logged", config.Approved)
		return config.Approved
	}

	return config.Declined
}
