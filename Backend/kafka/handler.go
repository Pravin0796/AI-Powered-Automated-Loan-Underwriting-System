package kafka

import (
	"AI-Powered-Automated-Loan-Underwriting-System/config"
	"AI-Powered-Automated-Loan-Underwriting-System/experian"
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func handleLoanApplicationSubmitted(event models.Event) {
	var payload struct {
		LoanID uint   `json:"loan_id"`
		UserID uint   `json:"user_id"`
		Status string `json:"status"`
	}
	err := json.Unmarshal([]byte(event.Payload), &payload)
	if err != nil {
		log.Printf("Error unmarshalling event payload: %v", err)
		return
	}

	// Fetch loan application from DB
	loan := models.LoanApplication{}
	if err := config.DB.First(&loan, payload.LoanID).Error; err != nil {
		log.Printf("Error fetching loan application: %v", err)
		return
	}

	// Prepare Experian request
	experianReq := experian.ExperianRequest{
		SSN:              loan.SSN,
		UserID:           fmt.Sprint(loan.UserID),
		LoanAmount:       loan.LoanAmount,
		EmploymentStatus: loan.EmploymentStatus,
	}

	// Fetch credit report from mock Experian API
	report, err := experian.FetchMockCreditReport(experianReq, fmt.Sprint(loan.ID))
	if err != nil {
		log.Printf("Error fetching credit report: %v", err)
		return
	}

	// Save credit report to DB
	creditReport := models.CreditReport{
		UserID:            loan.UserID,
		LoanApplicationID: loan.ID,
		CreditScore:       report.CreditScore,
		DelinquencyFlag:   report.DelinquencyFlag,
		ReportData:        report.ReportData,
		FraudIndicators:   report.FraudIndicators,
	}
	if err := config.DB.Create(&creditReport).Error; err != nil {
		log.Printf("Error saving credit report to DB: %v", err)
		return
	}

	// Optional: Process loan decision (AI model, rules, etc.)
	// For simplicity, let's assume we approve loans with a score above 700
	loanDecision := false
	if report.CreditScore >= 700 {
		loanDecision = true
	}

	// Save loan decision
	loanDecisionRecord := models.LoanDecision{
		LoanApplicationID: loan.ID,
		AiDecision:        loanDecision,
		Reasoning:         "Credit score evaluation",
	}
	if err := config.DB.Create(&loanDecisionRecord).Error; err != nil {
		log.Printf("Error saving loan decision to DB: %v", err)
		return
	}

	// Publish loan evaluated event (optional)
	event = models.Event{
		EventType: "LoanEvaluated",
		Payload:   fmt.Sprintf(`{"loan_id":%d,"user_id":%d,"status":"%s","decision":"%t"}`, loan.ID, loan.UserID, loan.ApplicationStatus, loanDecision),
		Timestamp: time.Now(),
	}

	kafkaServer := config.GetKafkaServer()

	producer := NewProducer(kafkaServer, kafkaTopic)
	defer producer.Close() 
	if err := producer.SendMessage(event); err != nil {
		log.Printf("Error publishing loan evaluated event: %v", err)
	}
}
