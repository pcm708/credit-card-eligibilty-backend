package check

import (
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/stretchr/testify/mock"
	"strconv"
)

type MockCheckImpl struct {
	mock.Mock
}

// IsNumberPreApproved is a mock function for the IsNumberPreApproved function in the CheckInterface.
func (m *MockCheckImpl) IsNumberPreApproved(data model.RecordData) (bool, error) {
	args := m.Called(data)
	return args.Bool(0), args.Error(1)
}

// IfApplicantPoliticallyExposed is a mock function for the IfApplicantPoliticallyExposed function in the CheckInterface.
func (m *MockCheckImpl) IfApplicantPoliticallyExposed(data model.RecordData) bool {
	if data.PoliticallyExposed != nil && *data.PoliticallyExposed {
		return true
	}
	return false
}

// IfValidAge is a mock function for the IfValidAge function in the CheckInterface.
func (m *MockCheckImpl) IfValidAge(data model.RecordData, config model.Config) bool {
	if data.Age >= config.MinAge {
		return true
	}
	return false
}

// IfValidIncome is a mock function for the IfValidIncome function in the CheckInterface.
func (m *MockCheckImpl) IfValidIncome(data model.RecordData, config model.Config) bool {
	if data.Income >= config.MinIncome {
		return true
	}
	return false
}

// IfValidNumberOfCreditCards is a mock function for the IfValidNumberOfCreditCards function in the CheckInterface.
func (m *MockCheckImpl) IfValidNumberOfCreditCards(data model.RecordData, config model.Config) bool {
	if data.NumberOfCreditCards != nil && *data.NumberOfCreditCards <= config.MinNumberOfCC {
		return true
	}
	return false
}

// IsValidAreaCode is a mock function for the IsValidAreaCode function in the CheckInterface.
func (m *MockCheckImpl) IsValidAreaCode(data model.RecordData, config model.Config) bool {
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
	return false
}

// IfCreditRiskScoreLow is a mock function for the IfCreditRiskScoreLow function in the CheckInterface.
func (m *MockCheckImpl) IfCreditRiskScoreLow(data model.RecordData, config model.Config) bool {
	if data.NumberOfCreditCards != nil &&
		config.DesiredCreditRiskScore == "LOW" {
		return true
	}
	return false
}
