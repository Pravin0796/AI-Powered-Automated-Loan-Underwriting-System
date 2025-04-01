package models

import "time"

// LoanDecision stores the AI's decision on a loan application
type LoanDecision struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	LoanApplicationID uint      `gorm:"not null" json:"loan_application_id"`
	AIDecision        string    `gorm:"type:varchar(20);not null" json:"ai_decision"` // approved, rejected
	Reasoning         string    `gorm:"type:text" json:"reasoning"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
}
