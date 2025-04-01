package models

import (
	"time"
)

// LoanApplication represents a loan request
type LoanApplication struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	UserID         uint      `gorm:"not null" json:"user_id"`
	AmountRequested float64  `gorm:"type:decimal(10,2);not null" json:"amount_requested"`
	LoanTerm       int       `gorm:"not null" json:"loan_term"` // In months
	Status         string    `gorm:"type:varchar(20);default:'pending'" json:"status"` // pending, approved, rejected
	SubmittedAt    time.Time `gorm:"autoCreateTime" json:"submitted_at"`
}
