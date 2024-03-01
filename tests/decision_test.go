package tests

import (
	"github.com/honestbank/tech-assignment-backend-engineer/check"
	"github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/honestbank/tech-assignment-backend-engineer/reader"
	"github.com/honestbank/tech-assignment-backend-engineer/services"
	"github.com/honestbank/tech-assignment-backend-engineer/writer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDecisionEngineWhenNumberIsNotPreApproved(t *testing.T) {
	mockCheck := new(check.MockCheck)
	mockReader := new(reader.MockReaderImpl)
	mockWriter := new(writer.MockWriter)

	mockReader.On("GetConfig", mock.Anything).Return(model.Config{
		MinAge:                 18,
		MinIncome:              100000,
		MinNumberOfCC:          3,
		AllowedAreaCodes:       []int{0, 2, 5, 8},
		DesiredCreditRiskScore: "LOW",
	})
	mockWriter.On("LogToJSON", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	mockWriter.On("StorePreApprovedNumber", mock.Anything)

	services.IsNumberPreApproved = mockCheck
	services.Eligibility = mockCheck
	services.Reader = mockReader
	services.Writer = mockWriter

	t.Run("ShouldDeclineTheApplicant_IfApplicantAgeIsLessThan18Years", func(t *testing.T) {
		numberOfCreditCards := 3
		politicallyExposed := false
		dummyData := model.RecordData{
			Income:              100000,
			Age:                 17,
			PhoneNumber:         "023-456-780",
			NumberOfCreditCards: &numberOfCreditCards,
			JobIndustryCode:     "1-000 - Purpose",
			PoliticallyExposed:  &politicallyExposed,
		}

		result, _, _ := services.DecisionEngine(dummyData)
		assert.Equal(t, constants.DECLINED, result)
	})

	t.Run("ShouldDeclineTheApplicant_IfApplicantAgeEarnsLessThan100000", func(t *testing.T) {
		numberOfCreditCards := 3
		politicallyExposed := false
		dummyData := model.RecordData{
			Income:              90000,
			Age:                 32,
			PhoneNumber:         "023-456-780",
			NumberOfCreditCards: &numberOfCreditCards,
			JobIndustryCode:     "3 - Concrete",
			PoliticallyExposed:  &politicallyExposed,
		}

		result, _, _ := services.DecisionEngine(dummyData)
		assert.Equal(t, constants.DECLINED, result)
	})

	t.Run("ShouldDeclineTheApplicant_IfAreaCodeIsNotValid", func(t *testing.T) {
		numberOfCreditCards := 3
		politicallyExposed := false
		dummyData := model.RecordData{
			Income:              100000,
			Age:                 32,
			PhoneNumber:         "123-456-780",
			NumberOfCreditCards: &numberOfCreditCards,
			JobIndustryCode:     "5 - Concrete",
			PoliticallyExposed:  &politicallyExposed,
		}

		result, _, _ := services.DecisionEngine(dummyData)
		assert.Equal(t, constants.DECLINED, result)
	})

	t.Run("ShouldDeclineTheApplicant_IfItIsPoliticallyExposed", func(t *testing.T) {
		numberOfCreditCards := 3
		politicallyExposed := true
		dummyData := model.RecordData{
			Income:              100000,
			Age:                 33,
			PhoneNumber:         "023-456-780",
			NumberOfCreditCards: &numberOfCreditCards,
			JobIndustryCode:     "5 - Concrete",
			PoliticallyExposed:  &politicallyExposed,
		}

		result, _, _ := services.DecisionEngine(dummyData)
		assert.Equal(t, constants.DECLINED, result)
	})

	t.Run("ShouldDeclineTheApplicant_IfNumberOfCreditCardsAreMoreThan3", func(t *testing.T) {
		numberOfCreditCards := 4
		politicallyExposed := false
		dummyData := model.RecordData{
			Income:              100000,
			Age:                 33,
			PhoneNumber:         "023-456-780",
			NumberOfCreditCards: &numberOfCreditCards,
			JobIndustryCode:     "5 - Concrete",
			PoliticallyExposed:  &politicallyExposed,
		}

		result, _, _ := services.DecisionEngine(dummyData)
		assert.Equal(t, constants.DECLINED, result)
	})

	t.Run("ShouldDeclineTheApplicant_CreditRiskScoreIsNotLOW", func(t *testing.T) {
		numberOfCreditCards := 4
		politicallyExposed := false
		dummyData := model.RecordData{
			Income:              100000,
			Age:                 33,
			PhoneNumber:         "023-456-780",
			NumberOfCreditCards: &numberOfCreditCards,
			JobIndustryCode:     "5 - Concrete",
			PoliticallyExposed:  &politicallyExposed,
		}

		result, _, _ := services.DecisionEngine(dummyData)
		assert.Equal(t, constants.DECLINED, result)
	})

	t.Run("ShouldApproveTheApplicant_IfApplicantMeetsAllRequirements", func(t *testing.T) {
		numberOfCreditCards := 3
		politicallyExposed := false
		dummyData := model.RecordData{
			Income:              100000,
			Age:                 30,
			PhoneNumber:         "023-456-780",
			NumberOfCreditCards: &numberOfCreditCards,
			JobIndustryCode:     "5 - Concrete",
			PoliticallyExposed:  &politicallyExposed,
		}

		result, _, _ := services.DecisionEngine(dummyData)
		assert.Equal(t, constants.APPROVED, result)
	})
}

func TestDecisionEngineWhenNumberIsPreApproved(t *testing.T) {
	mockCheck := new(check.MockCheck)
	mockReader := new(reader.MockReaderImpl)
	mockWriter := new(writer.MockWriter)

	mockReader.On("GetConfig", mock.Anything).Return(model.Config{
		MinAge:                 18,
		MinIncome:              100000,
		MinNumberOfCC:          3,
		AllowedAreaCodes:       []int{0, 2, 5, 8},
		DesiredCreditRiskScore: "LOW",
	})
	mockCheck.On("IsNumberPreApproved", mock.Anything, mock.Anything).Return(true, 0, nil)
	mockWriter.On("LogToJSON", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	mockWriter.On("StorePreApprovedNumber", mock.Anything)

	services.IsNumberPreApproved = mockCheck
	services.Eligibility = mockCheck
	services.Reader = mockReader
	services.Writer = mockWriter

	t.Run("ShouldApproveTheApplicant_IfPhoneNumberIsPreApproved", func(t *testing.T) {
		numberOfCreditCards := 3
		politicallyExposed := false
		dummyData := model.RecordData{
			Income:              100000,
			Age:                 30,
			PhoneNumber:         "023-456-780",
			JobIndustryCode:     "5 - Concrete",
			NumberOfCreditCards: &numberOfCreditCards,
			PoliticallyExposed:  &politicallyExposed,
		}

		result, _, _ := services.DecisionEngine(dummyData)
		assert.Equal(t, constants.APPROVED, result)
	})

	t.Run("ShouldApproveTheApplicant_IfPhoneNumberIsPreApprovedIfEvenApplicantIsPoliticallyExposed", func(t *testing.T) {
		numberOfCreditCards := 3
		politicallyExposed := true
		dummyData := model.RecordData{
			Income:              100000,
			Age:                 30,
			PhoneNumber:         "023-456-780",
			NumberOfCreditCards: &numberOfCreditCards,
			JobIndustryCode:     "5 - Concrete",
			PoliticallyExposed:  &politicallyExposed,
		}

		result, _, _ := services.DecisionEngine(dummyData)
		assert.Equal(t, constants.APPROVED, result)
	})

	t.Run("ShouldApproveTheApplicant_IfPhoneNumberIsPreApprovedIfEvenApplicantHoldsMoreThan3CreditCard", func(t *testing.T) {
		numberOfCreditCards := 6
		politicallyExposed := true
		dummyData := model.RecordData{
			Income:              100000,
			Age:                 30,
			PhoneNumber:         "023-456-780",
			NumberOfCreditCards: &numberOfCreditCards,
			JobIndustryCode:     "5 - Concrete",
			PoliticallyExposed:  &politicallyExposed,
		}

		result, _, _ := services.DecisionEngine(dummyData)
		assert.Equal(t, constants.APPROVED, result)
	})

	t.Run("ShouldApproveTheApplicant_IfPhoneNumberIsPreApprovedIfEvenApplicantSAgeIsLessThan18Years", func(t *testing.T) {
		numberOfCreditCards := 3
		politicallyExposed := false
		dummyData := model.RecordData{
			Income:              100000,
			Age:                 12,
			PhoneNumber:         "023-456-780",
			NumberOfCreditCards: &numberOfCreditCards,
			PoliticallyExposed:  &politicallyExposed,
		}

		result, _, _ := services.DecisionEngine(dummyData)
		assert.Equal(t, constants.APPROVED, result)
	})

	t.Run("ShouldApproveTheApplicant_IfPhoneNumberIsPreApprovedIfEvenApplicantsAgeIsLessThan18Years", func(t *testing.T) {
		numberOfCreditCards := 3
		politicallyExposed := false
		dummyData := model.RecordData{
			Income:              100000,
			Age:                 12,
			PhoneNumber:         "023-456-780",
			NumberOfCreditCards: &numberOfCreditCards,
			PoliticallyExposed:  &politicallyExposed,
		}

		result, _, _ := services.DecisionEngine(dummyData)
		assert.Equal(t, constants.APPROVED, result)
	})

	t.Run("ShouldApproveTheApplicant_IfPhoneNumberIsPreApprovedEvenIfApplicantEarnsLessThan10000", func(t *testing.T) {
		numberOfCreditCards := 3
		politicallyExposed := false
		dummyData := model.RecordData{
			Income:              90000,
			Age:                 12,
			PhoneNumber:         "023-456-780",
			NumberOfCreditCards: &numberOfCreditCards,
			PoliticallyExposed:  &politicallyExposed,
		}

		result, _, _ := services.DecisionEngine(dummyData)
		assert.Equal(t, constants.APPROVED, result)
	})
}
