package kafka

import (
	"AI-Powered-Automated-Loan-Underwriting-System/config"
	"AI-Powered-Automated-Loan-Underwriting-System/experian"
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"AI-Powered-Automated-Loan-Underwriting-System/services"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func handleLoanApplicationSubmitted(event models.Event) {
	log.Println("[DEBUG] Received LoanApplicationSubmitted event")

	var payload struct {
		LoanID uint   `json:"loan_id"`
		UserID uint   `json:"user_id"`
		Status string `json:"status"`
	}
	err := json.Unmarshal([]byte(event.Payload), &payload)
	if err != nil {
		log.Printf("[ERROR] Unmarshalling event payload: %v", err)
		return
	}
	log.Printf("[DEBUG] Event payload unmarshalled: %+v", payload)

	// Fetch loan application from DB
	loan := models.LoanApplication{}
	if err := config.DB.First(&loan, payload.LoanID).Error; err != nil {
		log.Printf("[ERROR] Fetching loan application from DB: %v", err)
		return
	}
	log.Printf("[DEBUG] Fetched loan application: %+v", loan)

	// Prepare Experian request
	experianReq := experian.ExperianRequest{
		SSN:              loan.SSN,
		UserID:           fmt.Sprint(loan.UserID),
		LoanAmount:       loan.LoanAmount,
		EmploymentStatus: loan.EmploymentStatus,
	}
	log.Printf("[DEBUG] Created Experian request: %+v", experianReq)

	// Fetch credit report
	report, err := experian.FetchMockCreditReport(experianReq, fmt.Sprint(loan.ID))
	if err != nil {
		log.Printf("[ERROR] Fetching credit report: %v", err)
		return
	}
	log.Printf("[DEBUG] Fetched credit report: %+v", report)

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
		log.Printf("[ERROR] Saving credit report to DB: %v", err)
		return
	}
	log.Println("[DEBUG] Saved credit report to DB")

	// Decision logic
	decisionInput := services.LoanPredictionInput{
		LoanAmount:        loan.LoanAmount,
		LoanPurpose:       loan.LoanPurpose,
		EmploymentStatus:  loan.EmploymentStatus,
		AnnualIncome:      loan.GrossMonthlyIncome * 12,
		DTIRatio:          loan.DTIRatio,
		ReportCreditScore: creditReport.CreditScore,
		// UserCreditScore:     report.CreditScore,
		DelinquencyFlag:     creditReport.DelinquencyFlag,
		NumPaymentsMade:     10,   // if available
		NumLatePayments:     1,    // if available
		TotalAmountPaid:     2400, // if available
		PaymentSuccessRatio: 0.91, // if available
	}

	res, err := services.GetLoanDecision(decisionInput)
	if err != nil {
		log.Printf("ML model call failed: %v", err)
		return
	}

	loanDecision := models.LoanDecision{
		LoanApplicationID: loan.ID,
		AiDecision:        res.Decision,
		Reasoning:         res.Reasoning,
	}

	log.Printf("[DEBUG] Loan decision based on credit score (%d): %t", report.CreditScore, loanDecision)

	// Save loan decision
	loanDecisionRecord := models.LoanDecision{
		LoanApplicationID: loan.ID,
		AiDecision:        loanDecision,
		Reasoning:         "Credit score evaluation",
	}
	if err := config.DB.Create(&loanDecisionRecord).Error; err != nil {
		log.Printf("[ERROR] Saving loan decision to DB: %v", err)
		return
	}
	log.Println("[DEBUG] Saved loan decision to DB")

	// Publish event
	event = models.Event{
		EventType: "LoanEvaluated",
		Payload:   fmt.Sprintf(`{"loan_id":%d,"user_id":%d,"status":"%s","decision":"%t"}`, loan.ID, loan.UserID, loan.ApplicationStatus, loanDecision),
		Timestamp: time.Now(),
	}

	kafkaServer := config.GetKafkaServer()
	log.Printf("[DEBUG] Kafka server: %s", kafkaServer)

	producer := NewProducer(kafkaServer, kafkaTopic)
	defer func() {
		log.Println("[DEBUG] Closing Kafka producer")
		producer.Close()
	}()
	if err := producer.SendMessage(event); err != nil {
		log.Printf("[ERROR] Publishing loan evaluated event: %v", err)
	} else {
		log.Printf("[DEBUG] Published loan evaluated event for loan ID %d", loan.ID)
	}
}
