package check

import (
	"fmt"
	"net/http"
	"strconv"

	. "github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/honestbank/tech-assignment-backend-engineer/reader"
	"github.com/honestbank/tech-assignment-backend-engineer/writer"
)

var Writer writer.WriterInterface = &writer.WriterImpl{}

type CheckInterface interface {
	Check(data model.RecordData, config model.Config) (bool, int, error)
	SetNext(check CheckInterface)
}

type NumberPreApprovedCheck struct {
	next CheckInterface
}

type AgeCheck struct {
	next CheckInterface
}

type AreaCodeCheck struct {
	next CheckInterface
}

type IncomeCheck struct {
	next CheckInterface
}

type NumberOfCreditCardsCheck struct {
	next CheckInterface
}

type CreditRiskScoreCheck struct {
	next CheckInterface
}

type PoliticallyExposedCheck struct {
	next CheckInterface
}

func (n *NumberPreApprovedCheck) Check(data model.RecordData, config model.Config) (bool, int, error) {
	//preApprovedNumbers, code, err := reader.ExtractPreApprovedNumbers_Cloud()
	preApprovedNumbers, code, err := reader.ExtractPreApprovedNumbers_Local()
	if err != nil {
		return false, code, err
	}
	for _, number := range preApprovedNumbers {
		if number == data.PhoneNumber {
			return true, code, nil
		}
	}
	return false, code, nil
}

func (a *AgeCheck) Check(data model.RecordData, config model.Config) (bool, int, error) {
	if data.Age >= config.MinAge {
		if a.next != nil {
			return a.next.Check(data, config)
		}
		return true, http.StatusOK, nil
	}
	Writer.LogToJSON(data.PhoneNumber, INVALID_AGE, DECLINED, LOG_LEVEL_WARN)
	return false, http.StatusOK, nil
}

func (a *AreaCodeCheck) Check(data model.RecordData, config model.Config) (bool, int, error) {
	areaCodeStr := string(data.PhoneNumber[0])
	areaCode, err := strconv.Atoi(areaCodeStr)
	if err != nil {
		Writer.LogToJSON(data.PhoneNumber, INVALID_AREA_CODE, DECLINED, LOG_LEVEL_ERROR)
		return false, http.StatusInternalServerError, fmt.Errorf("Area code check failed")
	}
	for _, code := range config.AllowedAreaCodes {
		if areaCode == code {
			if a.next != nil {
				return a.next.Check(data, config)
			}
			return true, http.StatusOK, nil
		}
	}
	Writer.LogToJSON(data.PhoneNumber, INVALID_AREA_CODE, DECLINED, LOG_LEVEL_WARN)
	return false, http.StatusOK, nil
}

func (n *NumberOfCreditCardsCheck) Check(data model.RecordData, config model.Config) (bool, int, error) {
	if data.NumberOfCreditCards != nil && *data.NumberOfCreditCards <= config.MinNumberOfCC {
		if n.next != nil {
			return n.next.Check(data, config)
		}
		return true, http.StatusOK, nil
	}
	Writer.LogToJSON(data.PhoneNumber, INVALID_CC_NUMBER, DECLINED, LOG_LEVEL_WARN)
	return false, http.StatusOK, nil
}

func (i *IncomeCheck) Check(data model.RecordData, config model.Config) (bool, int, error) {
	if data.Income >= config.MinIncome {
		if i.next != nil {
			return i.next.Check(data, config)
		}
		return true, http.StatusOK, nil
	}
	Writer.LogToJSON(data.PhoneNumber, INVALID_INCOME, DECLINED, LOG_LEVEL_WARN)
	return false, http.StatusOK, nil
}

func (c *CreditRiskScoreCheck) Check(data model.RecordData, config model.Config) (bool, int, error) {
	if data.NumberOfCreditCards != nil &&
		config.DesiredCreditRiskScore == calculateCreditRisk(data.Age, *data.NumberOfCreditCards) {
		if c.next != nil {
			return c.next.Check(data, config)
		}
		return true, http.StatusOK, nil
	}
	Writer.LogToJSON(data.PhoneNumber, INVALID_CREDIT_RISK_SCORE, DECLINED, LOG_LEVEL_WARN)
	return false, http.StatusOK, nil
}

func (p *PoliticallyExposedCheck) Check(data model.RecordData, config model.Config) (bool, int, error) {
	if data.PoliticallyExposed != nil && *data.PoliticallyExposed {
		Writer.LogToJSON(data.PhoneNumber, POLITICALLY_EXPOSED, DECLINED, LOG_LEVEL_WARN)
		return false, http.StatusOK, nil
	}
	if p.next != nil {
		return p.next.Check(data, config)
	}
	return true, http.StatusOK, nil
}

func (n *NumberPreApprovedCheck) SetNext(check CheckInterface) {
	n.next = check
}
func (a *AgeCheck) SetNext(check CheckInterface) {
	a.next = check
}
func (a *AreaCodeCheck) SetNext(check CheckInterface) {
	a.next = check
}
func (n *NumberOfCreditCardsCheck) SetNext(check CheckInterface) {
	n.next = check
}
func (i *IncomeCheck) SetNext(check CheckInterface) {
	i.next = check
}
func (c *CreditRiskScoreCheck) SetNext(check CheckInterface) {
	c.next = check
}
func (p *PoliticallyExposedCheck) SetNext(check CheckInterface) {
	p.next = check
}

// IsNumberPreApprovedCheck returns an instance of NumberPreApprovedCheck
func IsNumberPreApprovedCheck() CheckInterface {
	preApprovedCheck := &NumberPreApprovedCheck{}
	preApprovedCheck.SetNext(nil)
	return preApprovedCheck
}

// CreateChecks creates the instances of all checks and sets up the chain of responsibility.
// It returns the first check in the chain
func CreateChecks() CheckInterface {
	// instance creation
	ageCheck := &AgeCheck{}
	areaCodeCheck := &AreaCodeCheck{}
	incomeCheck := &IncomeCheck{}
	numberOfCreditCardsCheck := &NumberOfCreditCardsCheck{}
	creditRiskScoreCheck := &CreditRiskScoreCheck{}
	politicallyExposedCheck := &PoliticallyExposedCheck{}

	// Set up the chain
	ageCheck.SetNext(areaCodeCheck)
	areaCodeCheck.SetNext(incomeCheck)
	incomeCheck.SetNext(numberOfCreditCardsCheck)
	numberOfCreditCardsCheck.SetNext(creditRiskScoreCheck)
	creditRiskScoreCheck.SetNext(politicallyExposedCheck)

	// Return the first check in the chain
	return ageCheck
}

// calculateCreditRisk calculates the credit risk score based on the age and number of credit cards.
func calculateCreditRisk(age, numberOfCreditCard int) string {
	sum := age + numberOfCreditCard
	mod := sum % 3
	if mod == 0 {
		return "LOW"
	}
	if mod == 1 {
		return "MEDIUM"
	}
	return "HIGH"
}
