package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type LoanPredictionInput struct {
	LoanAmount       float64 `json:"loan_amount"`
	LoanPurpose      string  `json:"loan_purpose"`
	EmploymentStatus string  `json:"employment_status"`
	AnnualIncome     float64 `json:"annual_income"`
	DTIRatio         float64 `json:"dti_ratio"`
	CreditScore      int     `json:"credit_score"`
	UserCreditScore  int     `json:"user_credit_score"`
	DelinquencyFlag  bool    `json:"delinquency_flag"`
}

func GetLoanDecision(input LoanPredictionInput) (string, error) {
	jsonData, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	resp, err := http.Post("http://localhost:8000/predict", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("prediction API failed")
	}

	var res map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	return res["decision"], nil
}
