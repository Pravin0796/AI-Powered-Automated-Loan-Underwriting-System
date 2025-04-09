package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a system user
type User struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	FullName    string         `gorm:"type:varchar(255);not null" json:"full_name"`
	Email       string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password    string         `gorm:"column:password_hash;type:text;not null" json:"-"`
	Phone       string         `gorm:"type:varchar(20);uniqueIndex;not null" json:"phone"`
	DateOfBirth time.Time      `json:"date_of_birth"`
	Address     string         `gorm:"type:text" json:"address"`
	CreditScore int            `gorm:"default:0" json:"credit_score"` // Default credit score
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete
}
