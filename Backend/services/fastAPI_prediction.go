package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type LoanPredictionInput struct {
	LoanAmount          float64 `json:"loan_amount"`
	LoanPurpose         string  `json:"loan_purpose"`
	EmploymentStatus    string  `json:"employment_status"`
	AnnualIncome        float64 `json:"annual_income"`
	DTIRatio            float64 `json:"dti_ratio"`
	ReportCreditScore   int     `json:"report_credit_score"`
	UserCreditScore     int     `json:"user_credit_score"`
	DelinquencyFlag     bool    `json:"delinquency_flag"`
	NumPaymentsMade     int     `json:"num_payments_made"`
	NumLatePayments     int     `json:"num_late_payments"`
	TotalAmountPaid     float64 `json:"total_amount_paid"`
	PaymentSuccessRatio float64 `json:"payment_success_ratio"`
}

func GetLoanDecision(input LoanPredictionInput) (string, error) {
	// Convert the input struct to JSON
	jsonData, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	// Send POST request to FastAPI backend
	resp, err := http.Post("http://localhost:8000/predict", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Handle HTTP status errors
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("prediction API failed with status code: " + resp.Status)
	}

	// Decode JSON response
	var res map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	decision, ok := res["decision"]
	if !ok {
		return "", errors.New("missing 'decision' field in response")
	}

	return decision, nil
}
