package model

// RecordData represents the structure of the record data
type RecordData struct {
	Income              int    `json:"income"`
	NumberOfCreditCards int    `json:"number_of_credit_cards"`
	Age                 int    `json:"age"`
	PoliticallyExposed  *bool  `json:"politically_exposed"`
	JobIndustryCode     string `json:"job_industry_code"`
	PhoneNumber         string `json:"phone_number"`
}
