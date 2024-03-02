package services

import (
	"github.com/honestbank/tech-assignment-backend-engineer/check"
	. "github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/honestbank/tech-assignment-backend-engineer/writer"
	"net/http"
)

var IsNumberPreApproved = check.CreatePhoneNumberCheck()
var EligibilityChecker = check.CreateEligibilityChecks()
var Writer writer.IWriter = &writer.WriterImpl{}

func DecisionEngine(data model.RecordData) (string, int, error) {
	return isApplicantEligible(data)
}

func isApplicantEligible(data model.RecordData) (string, int, error) {
	// Check if the number is pre-approved
	flag, _, err := IsNumberPreApproved.Check(data)
	if err != nil {
		return DECLINED, http.StatusServiceUnavailable, err
	}
	if flag {
		err = Writer.LogToJSON(data.PhoneNumber, PREAPPROVED_NUMBER, APPROVED, LOG_LEVEL_INFO)
		if err != nil {
			return DECLINED, http.StatusInternalServerError, err
		}
		return APPROVED, http.StatusOK, nil
	}

	//Else Start other eligibility checks
	eligibilityChecks, status, err := EligibilityChecker.Check(data)
	if err != nil {
		return DECLINED, status, err
	}
	if !eligibilityChecks {
		return DECLINED, status, err
	}

	// If the applicant passes all the checks, store the number and return "approved"
	err = Writer.StorePreApprovedNumber(data.PhoneNumber)
	if err != nil {
		return DECLINED, http.StatusServiceUnavailable, err
	}

	// Log the decision
	err = Writer.LogToJSON(data.PhoneNumber, NUMBER_LOGGED, APPROVED, LOG_LEVEL_INFO)
	if err != nil {
		return DECLINED, http.StatusInternalServerError, err
	}

	return APPROVED, http.StatusOK, nil
}
