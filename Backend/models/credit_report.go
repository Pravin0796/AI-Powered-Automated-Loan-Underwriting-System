package models

import "time"

// CreditReport stores financial data from Equifax/Experian
type CreditReport struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	ReportData  string    `gorm:"type:jsonb;not null" json:"report_data"`
	GeneratedAt time.Time `gorm:"autoCreateTime" json:"generated_at"`
}
