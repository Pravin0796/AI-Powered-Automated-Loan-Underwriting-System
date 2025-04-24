package models

import (
	"time"

	"gorm.io/gorm"
)

// LoanApplication represents a loan application submitted by a user
	type LoanApplication struct {
		ID                      uint           `gorm:"primaryKey" json:"id"`
		UserID                  uint           `gorm:"not null" json:"user_id"`
		User                    User           `gorm:"foreignKey:UserID" json:"user"`
		SSN                     string         `gorm:"type:varchar(20);not null" json:"ssn"`
		AddressArea             string         `gorm:"type:varchar(10);not null" json:"address_area"` // "urban" or "rural"
		LoanAmount              float64        `gorm:"type:decimal(15,2);not null" json:"loan_amount"`
		LoanPurpose             string         `gorm:"type:varchar(255);not null" json:"loan_purpose"`
		EmploymentStatus        string         `gorm:"type:varchar(50);not null" json:"employment_status"`
		GrossMonthlyIncome      float64        `gorm:"type:decimal(10,2);not null" json:"gross_monthly_income"` // Replace AnnualIncome
		TotalMonthlyDebtPayment float64        `gorm:"type:decimal(10,2);not null" json:"total_monthly_debt_payment"`
		DTIRatio                float64        `gorm:"type:decimal(5,2);default:0.00" json:"dti_ratio"`
		ApplicationStatus       string         `gorm:"type:varchar(20);default:'PENDING'" json:"application_status"`
		Reasoning               string         `gorm:"type:text" json:"reasoning"`
		CreditReportFetched     bool           `gorm:"default:false" json:"credit_report_fetched"`
		ExperianRequestID       string         `gorm:"type:varchar(100);default:''" json:"experian_request_id,omitempty"`
		CreditScore             int            `gorm:"default:0" json:"credit_score"`
		CreatedAt               time.Time      `gorm:"autoCreateTime" json:"created_at"`
		UpdatedAt               time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
		DeletedAt               gorm.DeletedAt `gorm:"index" json:"-"`
	}
