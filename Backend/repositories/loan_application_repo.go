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

func (r *LoanApplicationRepo) GetLoanStatistics(ctx context.Context) (totalApplications, approved, rejected, pending uint32, err error) {
	var totalApplicationsInt64, approvedInt64, rejectedInt64, pendingInt64 int64

	// Query to get the total number of loan applications
	err = r.DB.WithContext(ctx).Model(&models.LoanApplication{}).Count(&totalApplicationsInt64).Error
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// Query to get the count of approved loans (based on ApplicationStatus)
	err = r.DB.WithContext(ctx).Model(&models.LoanApplication{}).Where("application_status = ?", "APPROVED").Count(&approvedInt64).Error
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// Query to get the count of rejected loans (based on ApplicationStatus)
	err = r.DB.WithContext(ctx).Model(&models.LoanApplication{}).Where("application_status = ?", "REJECTED").Count(&rejectedInt64).Error
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// Query to get the count of pending loans (based on ApplicationStatus)
	err = r.DB.WithContext(ctx).Model(&models.LoanApplication{}).Where("application_status = ?", "PENDING").Count(&pendingInt64).Error
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// Convert int64 to uint32 for returning the result
	return uint32(totalApplicationsInt64), uint32(approvedInt64), uint32(rejectedInt64), uint32(pendingInt64), nil
}

func (r *LoanApplicationRepo) GetLoanApplicationBySSN(ctx context.Context, ssn string, loanApplication *models.LoanApplication) error {
	return r.DB.WithContext(ctx).Where("ssn = ?", ssn).First(&loanApplication).Error
}

func (r *LoanApplicationRepo) CountLoanApplications(ctx context.Context, total *int64) error {
	return r.DB.WithContext(ctx).Model(&models.LoanApplication{}).Count(total).Error
}

func (r *LoanApplicationRepo) GetPaginatedLoanApplications(ctx context.Context, limit, offset int, loans *[]models.LoanApplication) error {
	return r.DB.WithContext(ctx).
		Preload("User").
		Limit(limit).
		Offset(offset).
		Find(loans).Error
}
