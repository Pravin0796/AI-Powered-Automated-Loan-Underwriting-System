package services

import (
	pb "AI-Powered-Automated-Loan-Underwriting-System/created_proto/loan"
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"context"
	"gorm.io/gorm"
	"time"
)

type LoanServiceServer struct {
	pb.UnimplementedLoanServiceServer
	DB *gorm.DB
}

// ApplyForLoan handles loan application
func (s *LoanServiceServer) ApplyForLoan(ctx context.Context, req *pb.LoanRequest) (*pb.LoanResponse, error) {
	loan := models.LoanApplication{
		UserID:           uint(req.UserId),
		LoanAmount:       req.LoanAmount,
		LoanPurpose:      req.LoanPurpose,
		EmploymentStatus: req.EmploymentStatus,
		//AnnualIncome:        req.AnnualIncome,
		ApplicationStatus:   "PENDING",
		CreditReportFetched: false,
	}

	// Create loan application in the database
	if err := s.DB.Create(&loan).Error; err != nil {
		return nil, err
	}

	// Return loan response with loan ID and status
	return &pb.LoanResponse{
		LoanId: uint64(loan.ID),
		Status: loan.ApplicationStatus,
	}, nil
}

// GetLoanStatus returns the status of a loan application
func (s *LoanServiceServer) GetLoanStatus(ctx context.Context, req *pb.LoanStatusRequest) (*pb.LoanStatusResponse, error) {
	var loan models.LoanApplication

	// Fetch the loan application by its ID
	if err := s.DB.First(&loan, req.LoanId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	// Return loan status response
	return &pb.LoanStatusResponse{
		LoanId: uint64(loan.ID),
		Status: loan.ApplicationStatus,
	}, nil
}

// GetLoanApplicationDetails returns the details of a specific loan application
func (s *LoanServiceServer) GetLoanApplicationDetails(ctx context.Context, req *pb.LoanApplicationRequest) (*pb.LoanApplicationResponse, error) {
	var loan models.LoanApplication

	// Fetch the loan application details by its ID
	if err := s.DB.First(&loan, req.LoanId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	// Prepare and return the loan application response
	return &pb.LoanApplicationResponse{
		LoanId:           uint64(loan.ID),
		UserId:           uint64(loan.UserID),
		LoanAmount:       loan.LoanAmount,
		LoanPurpose:      loan.LoanPurpose,
		EmploymentStatus: loan.EmploymentStatus,
		//AnnualIncome:        loan.AnnualIncome,
		ApplicationStatus:   loan.ApplicationStatus,
		CreditReportFetched: loan.CreditReportFetched,
		ExperianRequestId:   loan.ExperianRequestID,
		CreditScore:         int32(loan.CreditScore),
		CreatedAt:           loan.CreatedAt.Format(time.RFC3339),
		UpdatedAt:           loan.UpdatedAt.Format(time.RFC3339),
	}, nil
}
