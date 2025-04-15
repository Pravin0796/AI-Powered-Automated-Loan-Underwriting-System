package models

import (
	"time"

	"gorm.io/gorm"
)

// Notification represents a message sent to a user about loan-related updates
type Notification struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	Title     string         `gorm:"type:varchar(255);not null" json:"title"`
	Message   string         `gorm:"type:text;not null" json:"message"`
	IsRead    bool           `gorm:"default:false" json:"is_read"`
	Type      string         `gorm:"type:varchar(50);not null" json:"type"` // e.g., "loan_update", "payment_reminder"
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
