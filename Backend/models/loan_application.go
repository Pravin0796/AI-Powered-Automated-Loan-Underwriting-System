package models

import (
	"time"
)

// LoanApplication represents a loan request
type LoanApplication struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	UserID          uint      `gorm:"not null" json:"user_id"`
	AmountRequested float64   `gorm:"type:decimal(10,2);not null" json:"amount_requested"`
	LoanTerm        int       `gorm:"not null" json:"loan_term"`                        // In months
	Status          string    `gorm:"type:varchar(20);default:'pending'" json:"status"` // pending, approved, rejected
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
