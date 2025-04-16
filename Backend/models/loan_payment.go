package models

import "time"

// LoanPayment represents a loan repayment transaction
type LoanPayment struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	LoanApplicationID uint      `gorm:"not null" json:"loan_application_id"`
	AmountPaid        float64   `gorm:"type:decimal(10,2);not null" json:"amount_paid"`
	PaymentDate       time.Time `gorm:"not null" json:"payment_date"`
	DueDate           time.Time `gorm:"not null" json:"due_date"`                         // Added: due date for the payment
	Status            string    `gorm:"type:varchar(20);default:'pending'" json:"status"` // pending, successful, failed
}
