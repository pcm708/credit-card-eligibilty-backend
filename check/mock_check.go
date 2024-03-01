package check

import (
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/stretchr/testify/mock"
	"net/http"
)

type MockCheck struct {
	mock.Mock
}

func (a *MockCheck) Check(data model.RecordData, config model.Config) (bool, int, error) {
	if data.Age <= config.MinAge &&
		*data.NumberOfCreditCards <= config.MinNumberOfCC &&
		data.Income <= config.MinIncome &&
		data.PoliticallyExposed != nil && *data.PoliticallyExposed == false &&
		data.NumberOfCreditCards != nil &&
		config.DesiredCreditRiskScore == calculateCreditRisk(data.Age, *data.NumberOfCreditCards) {
		return true, http.StatusOK, nil
	}
	return false, http.StatusOK, nil
}

func (m *MockCheck) SetNext(check CheckInterface) {
}
