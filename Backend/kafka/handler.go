package kafka

import (
	"AI-Powered-Automated-Loan-Underwriting-System/config"
	"AI-Powered-Automated-Loan-Underwriting-System/experian"
	"AI-Powered-Automated-Loan-Underwriting-System/ml_model"
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
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
	const maxRetries = 3

	var report *experian.CreditReportData

	for attempt := 1; attempt <= maxRetries; attempt++ {
		report, err = experian.FetchMockCreditReport(experianReq, fmt.Sprint(loan.ID))
		if err == nil {
			break
		}
		log.Printf("[WARN] Attempt %d: failed fetching credit report: %v", attempt, err)
		time.Sleep(time.Duration(attempt) * time.Second) // exponential backoff
	}

	if err != nil {
		log.Printf("[ERROR] All attempts failed fetching credit report: %v", err)
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
	tx := config.DB.Begin()
	if tx.Error != nil {
		log.Printf("[ERROR] Starting transaction: %v", tx.Error)
		return
	}

	// 1. Save credit report
	if err := tx.Create(&creditReport).Error; err != nil {
		log.Printf("[ERROR] Saving credit report: %v", err)
		tx.Rollback()
		return
	}

	if err := tx.Where("loan_application_id = ?", loan.ID).First(&creditReport).Error; err != nil {
		log.Printf("[ERROR] Fetching credit report from DB: %v", err)
		return
	}

	experianRequestIDStr := strconv.FormatUint(uint64(creditReport.ID), 10)

	// 2. Update loan application
	if err := tx.Model(&loan).Where("id = ?", loan.ID).Updates(map[string]interface{}{
		"credit_report_fetched": true,
		"credit_score":          report.CreditScore,
		"experian_request_id":   experianRequestIDStr,
	}).Error; err != nil {
		log.Printf("[ERROR] Updating loan: %v", err)
		tx.Rollback()
		return
	}

	// If everything succeeded
	if err := tx.Commit().Error; err != nil {
		log.Printf("[ERROR] Committing transaction: %v", err)
		return
	}

	log.Println("[DEBUG] Transaction committed successfully")

	// no_payments_made := 0
	// no_late_payments := 0
	// total_amount_paid := 0.0
	// payment_success_ratio := 0.0

	// // Fetch past payment history for this user, excluding current loan
	// paymentHistory := []models.LoanPayment{}
	// if err := config.DB.
	// 	Where("user_id = ? AND loan_application_id != ?", loan.UserID, loan.ID).
	// 	Find(&paymentHistory).Error; err != nil {
	// 	log.Printf("[ERROR] Fetching past payment history: %v", err)
	// }
	// log.Printf("[DEBUG] Fetched past payment history for user %d: %+v", loan.UserID, paymentHistory)

	// if len(paymentHistory) == 0 {
	// 	log.Println("[DEBUG] No past payment history found")
	// } else {
	// 	for _, payment := range paymentHistory {
	// 		no_payments_made++
	// 		if payment.Status == "successful" {
	// 			total_amount_paid += payment.AmountPaid
	// 		}
	// 		if payment.PaymentDate.After(payment.DueDate) {
	// 			no_late_payments++
	// 		}
	// 	}

	// 	if no_payments_made > 0 {
	// 		payment_success_ratio = total_amount_paid / float64(no_payments_made)
	// 	} else {
	// 		payment_success_ratio = 0.0
	// 	}
	// }

	// log.Printf("[DEBUG] Past payment stats: no_payments_made=%d, no_late_payments=%d, total_amount_paid=%.2f, payment_success_ratio=%.2f",
	// 	no_payments_made, no_late_payments, total_amount_paid, payment_success_ratio)

	// Decision logic
	decisionInput := ml_model.LoanPredictionInput{
		LoanAmount:        loan.LoanAmount,
		LoanPurpose:       loan.LoanPurpose,
		EmploymentStatus:  loan.EmploymentStatus,
		AnnualIncome:      loan.GrossMonthlyIncome * 12,
		DTIRatio:          loan.DTIRatio,
		ReportCreditScore: creditReport.CreditScore,
		DelinquencyFlag:   creditReport.DelinquencyFlag,
		// NumPaymentsMade:     no_payments_made,
		// NumLatePayments:     no_late_payments,
		// TotalAmountPaid:     total_amount_paid,
		// PaymentSuccessRatio: payment_success_ratio,
	}

	res, err := ml_model.GetLoanDecision(decisionInput)
	if err != nil {
		log.Printf("ML model call failed: %v", err)
		return
	}

	// Correct decision handling based on model's response (approved/rejected)
	var aiDecision bool
	if res.Decision == "approved" {
		aiDecision = true
	} else if res.Decision == "rejected" {
		aiDecision = false
	} else {
		log.Printf("[ERROR] Unexpected decision response: %s", res.Decision)
		return
	}

	// Save loan decision
	loanDecisionRecord := models.LoanDecision{
		LoanApplicationID: loan.ID,
		AiDecision:        aiDecision,
		Reasoning:         res.Reasoning,
	}
	log.Printf("[DEBUG] Loan decision based on credit score (%d): %t", creditReport.CreditScore, loanDecisionRecord.AiDecision)

	tx = config.DB.Begin()
	if tx.Error != nil {
		log.Printf("[ERROR] Starting transaction: %v", tx.Error)
		return
	}
	if err := tx.Create(&loanDecisionRecord).Error; err != nil {
		log.Printf("[ERROR] Saving loan decision to DB: %v", err)
		return
	}
	log.Println("[DEBUG] Saved loan decision to DB")

	if loanDecisionRecord.AiDecision {
		loan.ApplicationStatus = "APPROVED"
	} else {
		loan.ApplicationStatus = "REJECTED"
	}
	log.Printf("[DEBUG] Loan application status set to: %s", loan.ApplicationStatus)

	// Update loan application status
	if err := tx.Model(&loan).Where("id = ?", loan.ID).Updates(map[string]interface{}{
		"application_status": loan.ApplicationStatus,
		"reasoning":          loanDecisionRecord.Reasoning,
	}).Error; err != nil {
		log.Printf("[ERROR] Updating loan application status: %v", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	log.Println("[DEBUG] Transaction committed successfully")
	log.Printf("[DEBUG] Loan application status updated to: %s", loan.ApplicationStatus)
	log.Printf("[DEBUG] Loan decision saved: %t", loanDecisionRecord.AiDecision)
	log.Printf("[DEBUG] Reasoning: %s", loanDecisionRecord.Reasoning)

	//Publish event
	// event = models.Event{
	// 	EventType: "LoanEvaluated",
	// 	Payload:   fmt.Sprintf(`{"loan_id":%d,"user_id":%d,"status":"%s","decision":"%t","reasoning":"%s"}`, loan.ID, loan.UserID, loan.ApplicationStatus, loanDecisionRecord.AiDecision, loanDecisionRecord.Reasoning),
	// 	Timestamp: time.Now(),
	// }

	// kafkaServer := config.GetKafkaServer()
	// log.Printf("[DEBUG] Kafka server: %s", kafkaServer)

	// producer := NewProducer(kafkaServer, kafkaTopic)
	// defer func() {
	// 	log.Println("[DEBUG] Closing Kafka producer")
	// 	producer.Close()
	// }()
	// if err := producer.SendMessage(event); err != nil {
	// 	log.Printf("[ERROR] Publishing loan evaluated event: %v", err)
	// } else {
	// 	log.Printf("[DEBUG] Published loan evaluated event for loan ID %d", loan.ID)
	// }
}
