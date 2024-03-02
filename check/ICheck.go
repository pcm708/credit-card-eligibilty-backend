package check

import "github.com/honestbank/tech-assignment-backend-engineer/model"

type ICheck interface {
	Check(data model.RecordData) (bool, int, error)
	SetNext(check ICheck)
}
