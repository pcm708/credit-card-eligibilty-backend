package check

import (
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/stretchr/testify/mock"
	"net/http"
)

type MockNumberPreApprovedCheck struct {
	mock.Mock
	next ICheck
}

func (n *MockNumberPreApprovedCheck) Check(data model.RecordData, uid string) (bool, int, error) {
	return false, http.StatusOK, nil
}

func (n *MockNumberPreApprovedCheck) SetNext(check ICheck) {
	n.next = nil
}
