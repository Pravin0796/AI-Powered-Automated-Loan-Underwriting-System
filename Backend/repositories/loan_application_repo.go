package repositories

import (
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"context"

	"gorm.io/gorm"
)

type LoanApplicationRepo struct {
	DB *gorm.DB
}

func NewLoanApplicationRepo(DB *gorm.DB) *LoanApplicationRepo {
	return &LoanApplicationRepo{DB: DB}
}

func (r *LoanApplicationRepo) CreateLoanApplication(ctx context.Context, loanApplication models.LoanApplication) error {
	return r.DB.WithContext(ctx).Create(&loanApplication).Error
}

func (r *LoanApplicationRepo) GetLoanApplicationByID(ctx context.Context, loanID uint, loanApplication *models.LoanApplication) error {
	return r.DB.WithContext(ctx).Preload("User").First(&loanApplication, loanID).Error
}

func (r *LoanApplicationRepo) GetLoanApplicationByUserID(ctx context.Context, userID uint, loanApplications *[]models.LoanApplication) error {
	return r.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&loanApplications).Error
}

func (r *LoanApplicationRepo) UpdateLoanApplication(ctx context.Context, loanApplication *models.LoanApplication) error {
	return r.DB.WithContext(ctx).Save(&loanApplication).Error
}

func (r *LoanApplicationRepo) DeleteLoanApplication(ctx context.Context, loanID uint) error {
	return r.DB.WithContext(ctx).Delete(&models.LoanApplication{}, loanID).Error
}

func (r *LoanApplicationRepo) GetAllLoanApplications(ctx context.Context, loanApplications *[]models.LoanApplication) error {
	return r.DB.WithContext(ctx).Preload("User").Find(&loanApplications).Error
}

func (r *LoanApplicationRepo) GetLoanApplicationBySSN(ctx context.Context, ssn string, loanApplication *models.LoanApplication) error {
	return r.DB.WithContext(ctx).Where("ssn = ?", ssn).First(&loanApplication).Error
}
