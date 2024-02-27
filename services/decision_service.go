package services

import (
	"github.com/honestbank/tech-assignment-backend-engineer/check"
	. "github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/honestbank/tech-assignment-backend-engineer/reader"
	"github.com/honestbank/tech-assignment-backend-engineer/writer"
)

var Check check.CheckInterface = &check.CheckImpl{}
var Reader reader.ReaderInterface = &reader.ReaderImpl{}
var Writer writer.WriterInterface = &writer.WriterImpl{}

// DecisionEngine is the main function that decides if an applicant is eligible or not.
// It takes in a RecordData object and returns a string indicating the decision.
func DecisionEngine(data model.RecordData) string {
	// If the number is not pre-approved, check if the applicant is eligible.
	return isApplicantEligible(data)
}

func isApplicantEligible(data model.RecordData) string {
	// If the number is pre-approved, log the phone number and return "APPROVED".
	if Check.IsNumberPreApproved(data) {
		return APPROVED
	}

	config := Reader.GetConfig(CONFIG_FILE)

	if !Check.IfApplicantPoliticallyExposed(data) &&
		Check.IfValidAge(data, config) &&
		Check.IfValidIncome(data, config) &&
		Check.IfValidNumberOfCreditCards(data, config) &&
		Check.IsValidAreaCode(data, config) &&
		Check.IfCreditRiskScoreLow(data, config) {

		Writer.StorePreApprovedNumber(data.PhoneNumber)
		Writer.LogToJSON(data.PhoneNumber, NUMBER_LOGGED, APPROVED, LOG_LEVEL_INFO)
		return APPROVED
	}

	return DECLINED
}
