package services

import (
	pb "AI-Powered-Automated-Loan-Underwriting-System/created_proto/loan_payment"
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"AI-Powered-Automated-Loan-Underwriting-System/repositories"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LoanPaymentServiceServer struct {
	pb.UnimplementedLoanPaymentServiceServer
	Repo *repositories.LoanPaymentRepo
}

func NewLoanPaymentServiceServer(repo *repositories.LoanPaymentRepo) *LoanPaymentServiceServer {
	return &LoanPaymentServiceServer{
		Repo: repo,
	}
}

// CreateLoanPayment implements the CreateLoanPayment RPC method
func (s *LoanPaymentServiceServer) CreateLoanPayment(ctx context.Context, req *pb.CreateLoanPaymentRequest) (*pb.CreateLoanPaymentResponse, error) {
	loanPayment := models.LoanPayment{
		LoanApplicationID: uint(req.LoanApplicationId),
		AmountPaid:        req.AmountPaid,
		Status:            req.Status,
		PaymentDate:       req.PaymentDate.AsTime(),
	}

	// Call repository to save loan payment
	if err := s.Repo.CreateLoanPayment(ctx, loanPayment); err != nil {
		return nil, err
	}

	return &pb.CreateLoanPaymentResponse{
		LoanPaymentId: uint64(loanPayment.ID),
		Status:        "Payment successful",
	}, nil
}

// GetLoanPayment implements the GetLoanPayment RPC method
func (s *LoanPaymentServiceServer) GetLoanPayment(ctx context.Context, req *pb.GetLoanPaymentRequest) (*pb.GetLoanPaymentResponse, error) {
	var loanPayment models.LoanPayment

	// Fetch loan payment using the repository
	err := s.Repo.DB.WithContext(ctx).First(&loanPayment, req.LoanPaymentId).Error
	if err != nil {
		return nil, err
	}

	// Return the loan payment details as response
	return &pb.GetLoanPaymentResponse{
		LoanPaymentId:     uint64(loanPayment.ID),
		LoanApplicationId: uint64(loanPayment.LoanApplicationID),
		AmountPaid:        loanPayment.AmountPaid,
		Status:            loanPayment.Status,
		PaymentDate:       timestamppb.New(loanPayment.PaymentDate),
	}, nil
}
