package writer

import (
	"fmt"
	"github.com/stretchr/testify/mock"
)

// MockWriter is a mock struct for the IWriter.
type MockWriterImpl struct {
	mock.Mock
}

// StorePreApprovedNumber is a mock function for the StorePreApprovedNumber function in the IWriter.
func (m *MockWriterImpl) StorePreApprovedNumber(phoneNumber string) error {
	fmt.Print("mocking LogToJSON function in writer")
	return nil
}

// LogToJSON is a mock function for the LogToJSON function in the IWriter.
func (m *MockWriterImpl) LogToJSON(phoneNumber string, reason string, decision string, logLevel string) error {
	fmt.Print("mocking LogToJSON function in writer")
	return nil
}
