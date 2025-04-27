package services

import (
	//"AI-Powered-Automated-Loan-Underwriting-System/config"
	"AI-Powered-Automated-Loan-Underwriting-System/config"
	pb "AI-Powered-Automated-Loan-Underwriting-System/created_proto/loan"
	"AI-Powered-Automated-Loan-Underwriting-System/kafka"
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"AI-Powered-Automated-Loan-Underwriting-System/repositories"
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

type LoanServiceServer struct {
	pb.UnimplementedLoanServiceServer
	repo *repositories.LoanApplicationRepo
}

func NewLoanServiceServer(repo *repositories.LoanApplicationRepo) *LoanServiceServer {
	return &LoanServiceServer{repo: repo}
}

// ApplyForLoan handles loan application
func (s *LoanServiceServer) ApplyForLoan(ctx context.Context, req *pb.LoanRequest) (*pb.LoanResponse, error) {
	dti := 0.0
	if req.GrossMonthlyIncome > 0 {
		dti = req.TotalMonthlyDebtPayment / req.GrossMonthlyIncome
	}

	loan := models.LoanApplication{
		UserID:                  uint(req.UserId),
		SSN:                     req.Ssn,
		AddressArea:             req.AddressArea,
		LoanAmount:              req.LoanAmount,
		LoanPurpose:             req.LoanPurpose,
		EmploymentStatus:        req.EmploymentStatus,
		GrossMonthlyIncome:      req.GrossMonthlyIncome,
		TotalMonthlyDebtPayment: req.TotalMonthlyDebtPayment,
		DTIRatio:                dti,
		ApplicationStatus:       "PENDING",
		CreditReportFetched:     false,
	}

	if err := s.repo.CreateLoanApplication(ctx, loan); err != nil {
		return nil, errors.New("failed to create Loan application")
	}

	if err := s.repo.GetLoanApplicationBySSN(ctx, loan.SSN, &loan); err != nil {
		return nil, errors.New("failed to fetch Loan application")
	}

	event := models.Event{
		EventType: "LoanApplicationSubmitted",
		Payload:   fmt.Sprintf(`{"loan_id":%d,"user_id":%d,"status":"%s"}`, loan.ID, loan.UserID, loan.ApplicationStatus),
		Timestamp: time.Now(),
	}
	kafkaServer := config.GetKafkaServer()
	producer := kafka.NewProducer(kafkaServer, "LoanApplicationSubmitted")
	if err := producer.SendMessage(event); err != nil {
		log.Printf("Kafka produce error: %v", err)
	}

	return &pb.LoanResponse{
		LoanId: uint64(loan.ID),
		Status: loan.ApplicationStatus,
	}, nil
}

// GetLoanStatus returns the status of a loan application
func (s *LoanServiceServer) GetLoanStatus(ctx context.Context, req *pb.LoanStatusRequest) (*pb.LoanStatusResponse, error) {
	var loan models.LoanApplication

	// Fetch the loan application by its ID
	if err := s.repo.GetLoanApplicationByID(ctx, uint(req.LoanId), &loan); err != nil {
		return nil, errors.New("failed to fetch Loan application")
	}

	//kafka.ProduceEvent("LoanStatusChecked", fmt.Sprintf(`{"loan_id":%d,"status":"%s"}`, loan.ID, loan.ApplicationStatus))

	// Return loan status response
	return &pb.LoanStatusResponse{
		LoanId: uint64(loan.ID),
		Status: loan.ApplicationStatus,
	}, nil
}

// GetLoanApplicationDetails returns the details of a specific loan application
func (s *LoanServiceServer) GetLoanApplicationDetails(ctx context.Context, req *pb.LoanApplicationRequest) (*pb.LoanApplicationResponse, error) {
	var loan models.LoanApplication

	// Fetch the loan application by its ID
	if err := s.repo.GetLoanApplicationByID(ctx, uint(req.LoanId), &loan); err != nil {
		return nil, errors.New("failed to fetch Loan application")
	}

	//kafka.ProduceEvent("LoanStatusChecked", fmt.Sprintf(`{"loan_id":%d,"status":"%s"}`, loan.ID, loan.ApplicationStatus))

	// Prepare and return the loan application response
	return &pb.LoanApplicationResponse{
		LoanId:                  uint64(loan.ID),
		UserName:                loan.User.FullName,
		Ssn:                     loan.SSN,
		AddressArea:             loan.AddressArea,
		LoanAmount:              loan.LoanAmount,
		LoanPurpose:             loan.LoanPurpose,
		EmploymentStatus:        loan.EmploymentStatus,
		GrossMonthlyIncome:      loan.GrossMonthlyIncome,
		TotalMonthlyDebtPayment: loan.TotalMonthlyDebtPayment,
		DtiRatio:                loan.DTIRatio,
		ApplicationStatus:       loan.ApplicationStatus,
		CreditReportFetched:     loan.CreditReportFetched,
		ExperianRequestId:       loan.ExperianRequestID,
		CreditScore:             int32(loan.CreditScore),
		Reasoning:               loan.Reasoning,
		CreatedAt:               loan.CreatedAt.Format(time.RFC3339),
		UpdatedAt:               loan.UpdatedAt.Format(time.RFC3339),
		DeletedAt:               loan.DeletedAt.Time.Format(time.RFC3339),
	}, nil
}

// UpdateLoanStatus updates the status of a loan application
func (s *LoanServiceServer) UpdateApplicationStatus(ctx context.Context, req *pb.UpdateApplicationStatusRequest) (*pb.UpdateApplicationStatusResponse, error) {
	var loan models.LoanApplication

	// Fetch the loan application by its ID
	if err := s.repo.GetLoanApplicationByID(ctx, uint(req.LoanApplicationId), &loan); err != nil {
		return nil, errors.New("failed to fetch Loan application")
	}

	// Update the status and reasoning
	loan.ApplicationStatus = req.NewStatus
	loan.Reasoning = req.Reasoning

	if err := s.repo.UpdateLoanApplication(ctx, &loan); err != nil {
		return nil, errors.New("failed to Update Loan application")
	}

	return &pb.UpdateApplicationStatusResponse{
		Status: "Loan status updated successfully",
	}, nil
}

// GetAllLoanApplications retrieves all loan applications
func (s *LoanServiceServer) GetAllLoanApplications(ctx context.Context, req *pb.Empty) (*pb.LoanApplicationList, error) {
	var loans []models.LoanApplication

	// Fetch all loan applications
	if err := s.repo.GetAllLoanApplications(ctx, &loans); err != nil {
		return nil, errors.New("failed to fetch all Loan application")
	}

	var responses []*pb.LoanApplicationResponse
	for _, loan := range loans {
		responses = append(responses, &pb.LoanApplicationResponse{
			LoanId:                  uint64(loan.ID),
			UserName:                loan.User.FullName,
			Ssn:                     loan.SSN,
			AddressArea:             loan.AddressArea,
			LoanAmount:              loan.LoanAmount,
			LoanPurpose:             loan.LoanPurpose,
			EmploymentStatus:        loan.EmploymentStatus,
			GrossMonthlyIncome:      loan.GrossMonthlyIncome,
			TotalMonthlyDebtPayment: loan.TotalMonthlyDebtPayment,
			DtiRatio:                loan.DTIRatio,
			ApplicationStatus:       loan.ApplicationStatus,
			CreditReportFetched:     loan.CreditReportFetched,
			ExperianRequestId:       loan.ExperianRequestID,
			CreditScore:             int32(loan.CreditScore),
			Reasoning:               loan.Reasoning,
			CreatedAt:               loan.CreatedAt.Format(time.RFC3339),
			UpdatedAt:               loan.UpdatedAt.Format(time.RFC3339),
			DeletedAt:               loan.DeletedAt.Time.Format(time.RFC3339),
		})
	}

	return &pb.LoanApplicationList{
		Applications: responses,
	}, nil
}
