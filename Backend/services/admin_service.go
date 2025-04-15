package services

import (
	protos "AI-Powered-Automated-Loan-Underwriting-System/created_proto/loan"
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type AdminService struct {
	protos.UnimplementedLoanServiceServer
	DB *gorm.DB
}

// GetAllLoanApplications fetches all loan applications
func (s *AdminService) GetAllLoanApplications(ctx context.Context, req *protos.Empty) (*protos.LoanApplicationList, error) {
	var applications []models.LoanApplication
	if err := s.DB.Preload("User").Find(&applications).Error; err != nil {
		return nil, err
	}

	// Map models to proto message
	protoApplications := make([]*protos.LoanApplicationResponse, len(applications))
	for i, app := range applications {
		protoApplications[i] = &protos.LoanApplicationResponse{
			LoanId:              uint64(app.ID),
			UserId:              uint64(app.UserID),
			LoanAmount:          app.LoanAmount,
			LoanPurpose:         app.LoanPurpose,
			EmploymentStatus:    app.EmploymentStatus,
			AnnualIncome:        app.AnnualIncome,
			ApplicationStatus:   app.ApplicationStatus,
			CreditReportFetched: app.CreditReportFetched,
			ExperianRequestId:   app.ExperianRequestID,
			CreditScore:         int32(app.CreditScore),
			CreatedAt:           app.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:           app.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return &protos.LoanApplicationList{Applications: protoApplications}, nil
}

// UpdateApplicationStatus updates the status of a loan application
func (s *AdminService) UpdateApplicationStatus(ctx context.Context, req *protos.UpdateApplicationStatusRequest) (*protos.UpdateApplicationStatusResponse, error) {
	var application models.LoanApplication
	if err := s.DB.First(&application, req.LoanApplicationId).Error; err != nil {
		return nil, fmt.Errorf("loan application not found")
	}

	application.ApplicationStatus = req.NewStatus
	application.Reasoning = req.Reasoning
	if err := s.DB.Save(&application).Error; err != nil {
		return nil, err
	}

	return &protos.UpdateApplicationStatusResponse{Status: "Loan application status updated successfully"}, nil
}

// GetLoanStats fetches statistics about loan applications
func (s *AdminService) GetLoanStats(ctx context.Context, req *protos.Empty) (*protos.LoanStatsResponse, error) {
	var totalApplications int64
	var approved int64
	var rejected int64
	var pending int64

	s.DB.Model(&models.LoanApplication{}).Count(&totalApplications)
	s.DB.Model(&models.LoanApplication{}).Where("application_status = ?", "APPROVED").Count(&approved)
	s.DB.Model(&models.LoanApplication{}).Where("application_status = ?", "REJECTED").Count(&rejected)
	s.DB.Model(&models.LoanApplication{}).Where("application_status = ?", "PENDING").Count(&pending)

	return &protos.LoanStatsResponse{
		TotalApplications: uint32(totalApplications),
		Approved:          uint32(approved),
		Rejected:          uint32(rejected),
		Pending:           uint32(pending),
	}, nil
}
