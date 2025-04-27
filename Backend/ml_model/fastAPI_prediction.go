package ml_model

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// LoanPredictionInput holds the features sent to the ML model
type LoanPredictionInput struct {
	LoanAmount          float64 `json:"loan_amount"`
	LoanPurpose         string  `json:"loan_purpose"`
	EmploymentStatus    string  `json:"employment_status"`
	AnnualIncome        float64 `json:"annual_income"`
	DTIRatio            float64 `json:"dti_ratio"`
	ReportCreditScore   int     `json:"report_credit_score"`
	// UserCreditScore     int     `json:"user_credit_score"`
	DelinquencyFlag     bool    `json:"delinquency_flag"`
	NumPaymentsMade     int     `json:"num_payments_made"`
	NumLatePayments     int     `json:"num_late_payments"`
	TotalAmountPaid     float64 `json:"total_amount_paid"`
	PaymentSuccessRatio float64 `json:"payment_success_ratio"`
}

// LoanDecisionResponse holds the AI decision
type LoanDecisionResponse struct {
	Decision string `json:"decision"` // only decision
	Reasoning string `json:"reasoning"`
}

// GetLoanDecision sends input to the ML model and returns the result
func GetLoanDecision(input LoanPredictionInput) (*LoanDecisionResponse, error) {
	jsonData, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://localhost:8000/predict", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("prediction API failed with status code: " + resp.Status)
	}

	var res LoanDecisionResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	if res.Decision == "" {
		return nil, errors.New("missing 'decision' in ML response")
	}

	return &res, nil
}
