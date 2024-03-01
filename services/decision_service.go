package services

import (
	"github.com/honestbank/tech-assignment-backend-engineer/check"
	. "github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/honestbank/tech-assignment-backend-engineer/reader"
	"github.com/honestbank/tech-assignment-backend-engineer/writer"
	"net/http"
)

var Eligibility check.CheckInterface = check.CreateChecks()
var IsNumberPreApproved = check.IsNumberPreApprovedCheck()
var Reader reader.ReaderInterface = &reader.ReaderImpl{}
var Writer writer.WriterInterface = &writer.WriterImpl{}

// DecisionEngine is the main function that decides if an applicant is eligible or not.
// It takes in a RecordData object and returns a string indicating the decision.
func DecisionEngine(data model.RecordData) (string, int, error) {
	// If the number is not pre-approved, check if the applicant is eligible.
	return isApplicantEligible(data)
}

func isApplicantEligible(data model.RecordData) (string, int, error) {
	config := Reader.GetConfig(CONFIG_FILE)

	if ok, res, err := IsNumberPreApproved.Check(data, config); ok {
		// if number is pre-approved, log and return
		Writer.LogToJSON(data.PhoneNumber, PREAPPROVED_NUMBER, APPROVED, LOG_LEVEL_INFO)
		return APPROVED, res, nil
	} else if err != nil {
		return err.Error(), res, err
	}

	//If number is not pre-approved, create other eligibility checks
	// Start the checks
	if ok, res, err := Eligibility.Check(data, config); !ok {
		return DECLINED, res, err
	}

	Writer.StorePreApprovedNumber(data.PhoneNumber)
	Writer.LogToJSON(data.PhoneNumber, NUMBER_LOGGED, APPROVED, LOG_LEVEL_INFO)
	return APPROVED, http.StatusOK, nil
}
