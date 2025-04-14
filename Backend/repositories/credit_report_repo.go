package repositories

import (
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"context"
	"gorm.io/gorm"
)

type CreditReportRepo struct {
	DB *gorm.DB
}

func NewCreditReportRepo(DB *gorm.DB) *CreditReportRepo {
	return &CreditReportRepo{DB: DB}
}

func (r *CreditReportRepo) CreateCreditReport(ctx context.Context, creditReport models.CreditReport) error {
	return r.DB.WithContext(ctx).Create(&creditReport).Error
}

func (r *CreditReportRepo) GetCreditReportByLoanApplicationID(ctx context.Context, loanApplicationID uint, creditReport *models.CreditReport) error {
	return r.DB.WithContext(ctx).Where("loan_application_id = ?", loanApplicationID).First(&creditReport).Error
}

func (r *CreditReportRepo) UpdateCreditReport(ctx context.Context, creditReport models.CreditReport) error {
	return r.DB.WithContext(ctx).Save(&creditReport).Error
}

func (r *CreditReportRepo) DeleteCreditReport(ctx context.Context, creditReportID uint) error {
	return r.DB.WithContext(ctx).Delete(&models.CreditReport{}, creditReportID).Error
}
