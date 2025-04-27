package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type CreditReport struct {
	ID                uint            `gorm:"primaryKey" json:"id"`
	UserID            uint            `gorm:"not null" json:"user_id"`
	LoanApplicationID uint            `gorm:"not null" json:"loan_application_id"`
	ReportData        datatypes.JSON  `gorm:"type:jsonb;not null" json:"report_data"`
	CreditScore       int             `json:"credit_score"`
	FraudIndicators   datatypes.JSON  `json:"fraud_indicators,omitempty"`
	DelinquencyFlag   bool            `json:"delinquency_flag"`
	CreatedAt         time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"` // Soft delete
	User              User            `gorm:"foreignKey:UserID" json:"user"`
	LoanApplication   LoanApplication `gorm:"foreignKey:LoanApplicationID" json:"loan_application"`
}

// Hook to update User and LoanApplication credit scores
func (c *CreditReport) AfterCreate(tx *gorm.DB) (err error) {
	// Update User Credit Score
	if err = tx.Model(&User{}).
		Where("id = ?", c.UserID).
		Update("credit_score", c.CreditScore).Error; err != nil {
		return err
	}

	// Update LoanApplication Credit Score
	if err = tx.Model(&LoanApplication{}).
		Where("id = ?", c.LoanApplicationID).
		Update("credit_score", c.CreditScore).Error; err != nil {
		return err
	}

	return nil
}
