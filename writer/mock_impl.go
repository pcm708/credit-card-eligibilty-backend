package writer

import (
	"github.com/stretchr/testify/mock"
)

// MockWriter is a mock struct for the IWriter.
type MockWriterImpl struct {
	mock.Mock
}

// StorePreApprovedNumber is a mock function for the StorePreApprovedNumber function in the IWriter.
func (m *MockWriterImpl) StorePreApprovedNumber(phoneNumber string) error {
	return nil
}

// LogToJSON is a mock function for the LogToJSON function in the IWriter.
func (m *MockWriterImpl) LogToJSON(phoneNumber string, reason string, decision string, logLevel string) error {
	return nil
}
