package model

import (
	"errors"
	"unicode"
)

type RecordData struct {
	Income              int    `json:"income" validate:"required"`
	NumberOfCreditCards int    `json:"number_of_credit_cards" validate:"required"`
	Age                 int    `json:"age" validate:"required"`
	PoliticallyExposed  bool   `json:"politically_exposed" validate:"required"`
	JobIndustryCode     string `json:"job_industry_code"`
	PhoneNumber         string `json:"phone_number" validate:"required"`
}

func (r *RecordData) Validate() error {
	if r.Income < 0 {
		return errors.New("please input a valid income")
	}
	if r.NumberOfCreditCards < 0 {
		return errors.New("please input a valid credit card number")
	}
	if r.Age <= 0 {
		return errors.New("please input a valid age")
	}
	if len(r.PhoneNumber) < 10 {
		return errors.New("please input a valid phone number")
	}
	for _, char := range r.PhoneNumber {
		if !unicode.IsDigit(char) && char != '-' {
			return errors.New("please input a valid phone number")
		}
	}
	return nil
}
