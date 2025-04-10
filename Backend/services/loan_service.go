package services

import (
	pb "AI-Powered-Automated-Loan-Underwriting-System/created_proto/loan"
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"context"

	"gorm.io/gorm"
)

type LoanServiceServer struct {
	pb.UnimplementedLoanServiceServer
	DB *gorm.DB
}

// Create/Apply Loan
func (s *LoanServiceServer) ApplyForLoan(ctx context.Context, req *pb.LoanRequest) (*pb.LoanResponse, error) {
	loan := models.LoanApplication{
		UserID:          uint(req.UserId), ///modifly later
		AmountRequested: req.AmountRequested,
		LoanTerm:        int(req.LoanTerm),
		Status:          "pending",
	}

	if err := s.DB.Create(&loan).Error; err != nil {
		return nil, err
	}

	return &pb.LoanResponse{
		LoanId: uint64(loan.ID),
		Status: loan.Status,
	}, nil
}

// Get Loan Status
func (s *LoanServiceServer) GetLoanStatus(ctx context.Context, req *pb.LoanStatusRequest) (*pb.LoanStatusResponse, error) {
	var loan models.LoanApplication

	if err := s.DB.First(&loan, req.LoanId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	return &pb.LoanStatusResponse{
		LoanId: uint64(loan.ID),
		Status: loan.Status,
	}, nil
}
