package model

// Config represents the structure of the configuration file
type Config struct {
	AllowedAreaCodes       []int  `json:"allowed_area_codes"`
	MinAge                 int    `json:"min_age"`
	MinIncome              int    `json:"min_income"`
	MinNumberOfCC          int    `json:"min_number_of_credit_cards"`
	DesiredCreditRiskScore string `json:"desired_credit_risk"`
	Approved               string `json:"approved"`
	Declined               string `json:"declined"`
}
