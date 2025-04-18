package services

import (
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"AI-Powered-Automated-Loan-Underwriting-System/repositories"
	"context"
	"errors"
	"fmt"
	"time"

	pb "AI-Powered-Automated-Loan-Underwriting-System/created_proto/loan_decision"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LoanDecisionServiceServer struct {
	pb.UnimplementedLoanDecisionServiceServer
	Repo *repositories.LoanDecisionRepo
}

func NewLoanDecisionServiceServer(repo *repositories.LoanDecisionRepo) *LoanDecisionServiceServer {
	return &LoanDecisionServiceServer{Repo: repo}
}

func (s *LoanDecisionServiceServer) CreateLoanDecision(ctx context.Context, req *pb.CreateLoanDecisionRequest) (*pb.CreateLoanDecisionResponse, error) {
	// Convert string decision to bool
	var decisionBool bool
	switch req.AiDecision {
	case "approved":
		decisionBool = true
	case "rejected":
		decisionBool = false
	default:
		return nil, errors.New("invalid ai_decision value; must be 'approved' or 'rejected'")
	}

	loanDecision := models.LoanDecision{
		LoanApplicationID: uint(req.LoanApplicationId),
		AiDecision:        decisionBool,
		Reasoning:         req.Reasoning,
		CreatedAt:         time.Now(),
	}

	err := s.Repo.CreateLoanDecision(ctx, loanDecision)
	if err != nil {
		return nil, fmt.Errorf("failed to create loan decision: %w", err)
	}

	return &pb.CreateLoanDecisionResponse{
		LoanDecisionId: uint64(loanDecision.ID),
		Status:         "Created successfully",
	}, nil
}

func (s *LoanDecisionServiceServer) GetLoanDecision(ctx context.Context, req *pb.GetLoanDecisionRequest) (*pb.GetLoanDecisionResponse, error) {
	var decision models.LoanDecision
	err := s.Repo.GetLoanDecisionByLoanApplicationID(ctx, uint(req.LoanDecisionId), &decision)
	if err != nil {
		return nil, fmt.Errorf("loan decision not found: %w", err)
	}

	// Convert bool back to string
	var decisionStr string
	if decision.AiDecision {
		decisionStr = "approved"
	} else {
		decisionStr = "rejected"
	}

	return &pb.GetLoanDecisionResponse{
		LoanDecisionId:    uint64(decision.ID),
		LoanApplicationId: uint64(decision.LoanApplicationID),
		AiDecision:        decisionStr,
		Reasoning:         decision.Reasoning,
		CreatedAt:         timestamppb.New(decision.CreatedAt),
	}, nil
}
