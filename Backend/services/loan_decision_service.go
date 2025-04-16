package services

import (
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	pb "AI-Powered-Automated-Loan-Underwriting-System/created_proto/loan_decision"
	"gorm.io/gorm"
)

type LoanDecisionServiceServer struct {
	pb.UnimplementedLoanDecisionServiceServer
	DB *gorm.DB
}

func NewLoanDecisionServiceServer(db *gorm.DB) *LoanDecisionServiceServer {
	return &LoanDecisionServiceServer{DB: db}
}

func (s *LoanDecisionServiceServer) CreateLoanDecision(ctx context.Context, req *pb.CreateLoanDecisionRequest) (*pb.CreateLoanDecisionResponse, error) {
	loanDecision := models.LoanDecision{
		LoanApplicationID: uint(req.LoanApplicationId),
		//AiDecision:        req.AiDecision,
		Reasoning: req.Reasoning,
		CreatedAt: time.Now(),
	}

	if err := s.DB.Create(&loanDecision).Error; err != nil {
		return nil, err
	}

	return &pb.CreateLoanDecisionResponse{
		LoanDecisionId: uint64(loanDecision.ID),
		Status:         "Created successfully",
	}, nil
}

func (s *LoanDecisionServiceServer) GetLoanDecision(ctx context.Context, req *pb.GetLoanDecisionRequest) (*pb.GetLoanDecisionResponse, error) {
	var decision models.LoanDecision
	if err := s.DB.First(&decision, req.LoanDecisionId).Error; err != nil {
		return nil, err
	}

	return &pb.GetLoanDecisionResponse{
		LoanDecisionId:    uint64(decision.ID),
		LoanApplicationId: uint64(decision.LoanApplicationID),
		//AiDecision:        decision.AiDecision,
		Reasoning: decision.Reasoning,
		CreatedAt: timestamppb.New(decision.CreatedAt),
	}, nil
}
