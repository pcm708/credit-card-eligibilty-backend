package writer

import "github.com/stretchr/testify/mock"

// MockWriter is a mock struct for the WriterInterface.
type MockWriter struct {
	mock.Mock
}

// StorePreApprovedNumber is a mock function for the StorePreApprovedNumber function in the WriterInterface.
func (m *MockWriter) StorePreApprovedNumber(phoneNumber string) {
}

// LogToJSON is a mock function for the LogToJSON function in the WriterInterface.
func (m *MockWriter) LogToJSON(phoneNumber string, reason string, decision string, logLevel string) error {
	return nil
}
