package tests

import (
	"encoding/json"
	"github.com/honestbank/tech-assignment-backend-engineer/check"
	. "github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/honestbank/tech-assignment-backend-engineer/reader"
	"github.com/honestbank/tech-assignment-backend-engineer/services"
	"github.com/honestbank/tech-assignment-backend-engineer/writer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"strconv"
	"testing"
)

func TestDecisionEngineWithJsonRecords(t *testing.T) {
	mockCheck := new(check.MockCheckImpl)
	mockReader := new(reader.MockReaderImpl)
	mockWriter := new(writer.MockWriter)
	services.Check = mockCheck
	services.Reader = mockReader
	services.Writer = mockWriter
	mockCheck.On("IsNumberPreApproved", mock.Anything).Return(false)
	mockReader.On("GetConfig", mock.Anything).Return(model.Config{
		MinAge:                 18,
		MinIncome:              100000,
		MinNumberOfCC:          3,
		AllowedAreaCodes:       []int{0, 2, 5, 8},
		DesiredCreditRiskScore: "LOW",
	})

	// Open the JSON file
	jsonFile, err := os.Open(JSON_RECORDS_5)
	if err != nil {
		t.Fatal(err)
	}
	defer jsonFile.Close()

	// Decode the JSON data into a slice of model.RecordData
	var recordData []model.RecordData
	err = json.NewDecoder(jsonFile).Decode(&recordData)
	if err != nil {
		t.Fatal(err)
	}

	// Iterate over the slice and generate a test for each JSON entry
	for i, data := range recordData {
		t.Run("TestRecord_"+strconv.Itoa(i), func(t *testing.T) {
			//numberOfCreditCards := 3
			//politicallyExposed := false
			dummyData := data

			result := services.DecisionEngine(dummyData)
			assert.Equal(t, DECLINED, result)
		})
	}
}
