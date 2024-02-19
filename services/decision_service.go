package services

import (
	"github.com/honestbank/tech-assignment-backend-engineer/check"
	"github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/honestbank/tech-assignment-backend-engineer/utils"
)

func DecisionEngine(data model.RecordData) string {
	return isApplicantEligible(data)
}

func isApplicantEligible(data model.RecordData) string {
	// check is number is preapproved or not
	if check.IsNumberPreApproved(data) {
		return constants.APPROVED
	}

	if !check.IfApplicantPoliticallyExposed(data) &&
		check.IfValidAge(data) &&
		check.IfValidIncome(data) &&
		check.IfValidNumberOfCreditCards(data) &&
		check.IsValidAreaCode(data) &&
		check.IfCreditRiskScoreLow(data) {

		utils.StorePreApprovedNumber(data.PhoneNumber)
		utils.LogToJSON(data.PhoneNumber, constants.NUMBER_LOGGED,
			constants.APPROVED, constants.LOG_LEVEL_INFO)
		return constants.APPROVED
	}

	return constants.DECLINED
}
