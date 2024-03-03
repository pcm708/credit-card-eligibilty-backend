package tests

import (
	"encoding/json"
	"github.com/honestbank/tech-assignment-backend-engineer/check"
	"os"
	"strconv"
	"testing"

	. "github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"github.com/honestbank/tech-assignment-backend-engineer/services"
	"github.com/honestbank/tech-assignment-backend-engineer/writer"
	"github.com/stretchr/testify/assert"
)

func TestDecisionEngineWithJsonRecords(t *testing.T) {
	check.Writer = &writer.MockWriterImpl{}
	services.EligibilityChecker = check.CreateEligibilityChecks()
	services.IsNumberPreApproved = &check.MockNumberPreApprovedCheck{}
	services.Writer = &writer.MockWriterImpl{}

	// Open the JSON file
	jsonFile, err := os.Open(JSON_RECORDS_1000)
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

	// Iterate over the slice and generate a tests for each JSON entry
	for i, data := range recordData {
		t.Run("TestRecord_"+strconv.Itoa(i+1), func(t *testing.T) {
			dummyData := data
			result, _, _, _ := services.DecisionEngine(dummyData)
			assert.Equal(t, DECLINED, result)
		})
	}
}
