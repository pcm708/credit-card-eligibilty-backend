package check

import (
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/stretchr/testify/mock"
)

type MockCheckUtil struct {
	mock.Mock
}

func (m *MockCheckUtil) Check(data model.RecordData, config model.Config) (bool, int, error) {
	args := m.Called(data, config)
	return args.Bool(0), args.Int(1), args.Error(2)
}

func (m *MockCheckUtil) SetNext(check EligibilityCheck) {
	m.Called(check)
}
