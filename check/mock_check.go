package check

import (
	. "github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/stretchr/testify/mock"
	"net/http"
)

type MockCheck struct {
	mock.Mock
}

func (a *MockCheck) Check(data model.RecordData) (bool, int, error) {
	if data.Age <= MIN_AGE &&
		*data.NumberOfCreditCards <= MIN_NUMBER_OF_CC &&
		data.Income <= MIN_INCOME &&
		data.PoliticallyExposed != nil && *data.PoliticallyExposed == false &&
		data.NumberOfCreditCards != nil &&
		DESIRED_CREDIT_RISK_SCORE == calculateCreditRisk(data.Age, *data.NumberOfCreditCards) {
		return true, http.StatusOK, nil
	}
	return false, http.StatusOK, nil
}

func (m *MockCheck) SetNext(check CheckInterface) {
}
