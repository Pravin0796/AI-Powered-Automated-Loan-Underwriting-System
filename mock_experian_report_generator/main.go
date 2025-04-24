package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type ExperianRequest struct {
	SSN              string  `json:"ssn"`
	UserID           string  `json:"userId"`
	LoanAmount       float64 `json:"loanAmount"`
	EmploymentStatus string  `json:"employmentStatus"`
}

type CreditReportResponse struct {
	CreditScore     int         `json:"creditScore"`
	FraudIndicators interface{} `json:"fraudIndicators"`
	DelinquencyFlag bool        `json:"delinquencyFlag"`
	ReportData      interface{} `json:"reportData"`
}

func mockExperianHandler(w http.ResponseWriter, r *http.Request) {
	var req ExperianRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Printf("Received mock Experian request: %+v\n", req)

	response := CreditReportResponse{
		CreditScore: gofakeit.Number(600, 750),
		FraudIndicators: map[string]bool{
			"syntheticIdentity": gofakeit.Bool(),
			"multipleSSN":       gofakeit.Bool(),
		},
		DelinquencyFlag: gofakeit.Bool(),
		ReportData: map[string]interface{}{
			"tradelines": []map[string]interface{}{
				{
					"accountType":   gofakeit.RandomString([]string{"credit card", "auto loan", "mortgage"}),
					"balance":       gofakeit.Price(100, 10000),
					"creditLimit":   gofakeit.Price(1000, 20000),
					"paymentStatus": gofakeit.RandomString([]string{"current", "late", "default"}),
					"openedDate":    gofakeit.DateRange(time.Now().AddDate(-10, 0, 0), time.Now()).Format("2006-01-02"),
				},
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/mock-experian/credit-report", mockExperianHandler)
	log.Println("Mock Experian server running on :8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
