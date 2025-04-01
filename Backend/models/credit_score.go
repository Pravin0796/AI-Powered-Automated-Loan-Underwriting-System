package models

import "time"

// CreditScore represents a user's credit score
type CreditScore struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;unique" json:"user_id"`
	Score     int       `gorm:"not null" json:"score"`
	ReportID  string    `gorm:"unique;not null" json:"report_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
