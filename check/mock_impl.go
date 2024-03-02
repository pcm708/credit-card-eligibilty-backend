package check

import (
	. "github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/stretchr/testify/mock"
	"net/http"
)

type MockEligibilityCheck struct {
	mock.Mock
	next ICheck
}

func (n *MockEligibilityCheck) SetNext(check ICheck) {
	n.next = check
}

func (a *MockEligibilityCheck) Check(data model.RecordData) (bool, int, error) {
	if data.Age >= MIN_AGE &&
		data.NumberOfCreditCards >= MIN_NUMBER_OF_CC &&
		data.Income >= MIN_INCOME &&
		data.PoliticallyExposed != nil && *data.PoliticallyExposed == false &&
		DESIRED_CREDIT_RISK_SCORE == calculateCreditRisk(data.Age, data.NumberOfCreditCards) {
		return true, http.StatusOK, nil
	}
	return false, http.StatusOK, nil
}

type MockNumberPreApprovedCheck struct {
	mock.Mock
	next ICheck
}

func (n *MockNumberPreApprovedCheck) Check(data model.RecordData) (bool, int, error) {
	return false, http.StatusOK, nil
}

func (n *MockNumberPreApprovedCheck) SetNext(check ICheck) {
	n.next = nil
}
