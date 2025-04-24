package experian

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gorm.io/datatypes"
)

// Change this when switching to real API
const ExperianAPIURL = "http://localhost:8081/mock-experian/credit-report"

type ExperianRequest struct {
	SSN              string  `json:"ssn"`
	UserID           string  `json:"userId"`
	LoanAmount       float64 `json:"loanAmount"`
	EmploymentStatus string  `json:"employmentStatus"`
}

type ExperianResponse struct {
	CreditScore     int             `json:"creditScore"`
	FraudIndicators json.RawMessage `json:"fraudIndicators"`
	DelinquencyFlag bool            `json:"delinquencyFlag"`
	ReportData      json.RawMessage `json:"reportData"`
}

// Used in your service logic where you create a CreditReport model instance
type CreditReportData struct {
	UserID            string
	LoanApplicationID string
	CreditScore       int
	DelinquencyFlag   bool
	ReportData        datatypes.JSON
	FraudIndicators   datatypes.JSON
}

func FetchMockCreditReport(req ExperianRequest, loanApplicationID string) (*CreditReportData, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal error: %v", err)
	}

	resp, err := http.Post(ExperianAPIURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("experian API error: %s", string(body))
	}

	var experianResp ExperianResponse
	if err := json.Unmarshal(body, &experianResp); err != nil {
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}

	return &CreditReportData{
		UserID:            req.UserID,
		LoanApplicationID: loanApplicationID,
		CreditScore:       experianResp.CreditScore,
		DelinquencyFlag:   experianResp.DelinquencyFlag,
		ReportData:        datatypes.JSON(experianResp.ReportData),
		FraudIndicators:   datatypes.JSON(experianResp.FraudIndicators),
	}, nil
}
