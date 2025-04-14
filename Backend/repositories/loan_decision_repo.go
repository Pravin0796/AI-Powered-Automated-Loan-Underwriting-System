package repositories

import (
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"context"
	"gorm.io/gorm"
)

type LoanDecisionRepo struct {
	DB *gorm.DB
}

func NewLoanDecisionRepo(DB *gorm.DB) *LoanDecisionRepo {
	return &LoanDecisionRepo{DB: DB}
}

func (r *LoanDecisionRepo) CreateLoanDecision(ctx context.Context, loanDecision models.LoanDecision) error {
	return r.DB.WithContext(ctx).Create(&loanDecision).Error
}

func (r *LoanDecisionRepo) GetLoanDecisionByLoanApplicationID(ctx context.Context, loanApplicationID uint, loanDecision *models.LoanDecision) error {
	return r.DB.WithContext(ctx).Where("loan_application_id = ?", loanApplicationID).First(&loanDecision).Error
}

func (r *LoanDecisionRepo) UpdateLoanDecision(ctx context.Context, loanDecision models.LoanDecision) error {
	return r.DB.WithContext(ctx).Save(&loanDecision).Error
}

func (r *LoanDecisionRepo) DeleteLoanDecision(ctx context.Context, loanDecisionID uint) error {
	return r.DB.WithContext(ctx).Delete(&models.LoanDecision{}, loanDecisionID).Error
}
