package check

import (
	"github.com/honestbank/tech-assignment-backend-engineer/db"
	"net/http"
	"strconv"

	. "github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/honestbank/tech-assignment-backend-engineer/writer"
)

var Writer writer.IWriter = &writer.WriterImpl{}

// .
// NumberPreApprovedCheck checks if the number is already pre-approved or not
type NumberPreApprovedCheck struct {
	next ICheck
}

func (n *NumberPreApprovedCheck) Check(data model.RecordData, uid string) (bool, int, error) {
	flag, err := db.CheckIfNumberPresent(data.PhoneNumber)
	if err != nil {
		return false, http.StatusServiceUnavailable, err
	}
	return flag, http.StatusOK, err
}
func (n *NumberPreApprovedCheck) SetNext(check ICheck) {
	n.next = nil
}

// .
// AgeCheck checks if the age is valid or not
type AgeCheck struct {
	next ICheck
}

func (a *AgeCheck) Check(data model.RecordData, uid string) (bool, int, error) {
	if data.Age >= MIN_AGE {
		if a.next != nil {
			return a.next.Check(data, uid)
		}
		return true, http.StatusOK, nil
	}
	Writer.LogToJSON(uid, INVALID_AGE, DECLINED, LOG_LEVEL_WARN)
	return false, http.StatusOK, nil
}
func (a *AgeCheck) SetNext(check ICheck) {
	a.next = check
}

// .
// AreaCodeCheck checks if the area code is allowed or not
type AreaCodeCheck struct {
	next ICheck
}

func (a *AreaCodeCheck) Check(data model.RecordData, uid string) (bool, int, error) {
	areaCodeStr := string(data.PhoneNumber[0])
	areaCode, err := strconv.Atoi(areaCodeStr)
	if err != nil {
		Writer.LogToJSON(uid, INVALID_AREA_CODE, DECLINED, LOG_LEVEL_ERROR)
		return false, http.StatusInternalServerError, err
	}
	for _, code := range ALLOWED_AREA_CODE {
		if areaCode == code {
			if a.next != nil {
				return a.next.Check(data, uid)
			}
			return true, http.StatusOK, nil
		}
	}
	Writer.LogToJSON(uid, INVALID_AREA_CODE, DECLINED, LOG_LEVEL_WARN)
	return false, http.StatusOK, nil
}
func (a *AreaCodeCheck) SetNext(check ICheck) {
	a.next = check
}

// .
// IncomeCheck checks if the income is valid or not
type IncomeCheck struct {
	next ICheck
}

func (i *IncomeCheck) Check(data model.RecordData, uid string) (bool, int, error) {
	if data.Income >= MIN_INCOME {
		if i.next != nil {
			return i.next.Check(data, uid)
		}
		return true, http.StatusOK, nil
	}
	Writer.LogToJSON(uid, INVALID_INCOME, DECLINED, LOG_LEVEL_WARN)
	return false, http.StatusOK, nil
}
func (i *IncomeCheck) SetNext(check ICheck) {
	i.next = check
}

// .
// NumberOfCreditCardsCheck checks if the number of credit cards is under the limit or not
type NumberOfCreditCardsCheck struct {
	next ICheck
}

func (n *NumberOfCreditCardsCheck) Check(data model.RecordData, uid string) (bool, int, error) {
	if data.NumberOfCreditCards != nil && *data.NumberOfCreditCards <= MAX_NUMBER_OF_CC {
		if n.next != nil {
			return n.next.Check(data, uid)
		}
		return true, http.StatusOK, nil
	}
	Writer.LogToJSON(uid, INVALID_CC_NUMBER, DECLINED, LOG_LEVEL_WARN)
	return false, http.StatusOK, nil
}
func (n *NumberOfCreditCardsCheck) SetNext(check ICheck) {
	n.next = check
}

// .
// CreditRiskScoreCheck checks if the credit risk score is Low or not
type CreditRiskScoreCheck struct {
	next ICheck
}

func (c *CreditRiskScoreCheck) Check(data model.RecordData, uid string) (bool, int, error) {
	if data.NumberOfCreditCards != nil &&
		DESIRED_CREDIT_RISK_SCORE == calculateCreditRisk(data.Age, *data.NumberOfCreditCards) {
		if c.next != nil {
			return c.next.Check(data, uid)
		}
		return true, http.StatusOK, nil
	}
	Writer.LogToJSON(uid, INVALID_CREDIT_RISK_SCORE, DECLINED, LOG_LEVEL_WARN)
	return false, http.StatusOK, nil
}
func (c *CreditRiskScoreCheck) SetNext(check ICheck) {
	c.next = check
}

// .
// PoliticallyExposedCheck checks if the person is politically exposed or not
type PoliticallyExposedCheck struct {
	next ICheck
}

func (p *PoliticallyExposedCheck) Check(data model.RecordData, uid string) (bool, int, error) {
	if data.PoliticallyExposed != nil && *data.PoliticallyExposed {
		Writer.LogToJSON(uid, POLITICALLY_EXPOSED, DECLINED, LOG_LEVEL_WARN)
		return false, http.StatusOK, nil
	}
	if p.next != nil {
		return p.next.Check(data, uid)
	}
	return true, http.StatusOK, nil
}
func (p *PoliticallyExposedCheck) SetNext(check ICheck) {
	p.next = check
}
