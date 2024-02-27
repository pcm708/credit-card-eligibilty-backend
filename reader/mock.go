package reader

import (
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/stretchr/testify/mock"
)

// MockReader is a mock struct for the ReaderInterface.
type MockReaderImpl struct {
	mock.Mock
}

// GetConfig is a mock function for the GetConfig function in the ReaderInterface.
func (m *MockReaderImpl) GetConfig(configFile string) model.Config {
	args := m.Called(configFile)
	return args.Get(0).(model.Config)
}
