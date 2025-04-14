package repositories

import (
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"context"
	"gorm.io/gorm"
)

type LoanPaymentRepo struct {
	DB *gorm.DB
}

func NewLoanPaymentRepo(DB *gorm.DB) *LoanPaymentRepo {
	return &LoanPaymentRepo{DB: DB}
}

func (r *LoanPaymentRepo) CreateLoanPayment(ctx context.Context, loanPayment models.LoanPayment) error {
	return r.DB.WithContext(ctx).Create(&loanPayment).Error
}

func (r *LoanPaymentRepo) GetLoanPaymentsByLoanApplicationID(ctx context.Context, loanApplicationID uint, loanPayments *[]models.LoanPayment) error {
	return r.DB.WithContext(ctx).Where("loan_application_id = ?", loanApplicationID).Find(&loanPayments).Error
}

func (r *LoanPaymentRepo) UpdateLoanPayment(ctx context.Context, loanPayment models.LoanPayment) error {
	return r.DB.WithContext(ctx).Save(&loanPayment).Error
}

func (r *LoanPaymentRepo) DeleteLoanPayment(ctx context.Context, loanPaymentID uint) error {
	return r.DB.WithContext(ctx).Delete(&models.LoanPayment{}, loanPaymentID).Error
}
