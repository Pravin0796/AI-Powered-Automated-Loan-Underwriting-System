package services

import (
	pb "AI-Powered-Automated-Loan-Underwriting-System/created_proto/loan"
	//"AI-Powered-Automated-Loan-Underwriting-System/kafka"
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	//"encoding/json"
	//"fmt"
	"time"

	"gorm.io/gorm"
)

type LoanServiceServer struct {
	pb.UnimplementedLoanServiceServer
	DB *gorm.DB
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

	if err := s.DB.Create(&loan).Error; err != nil {
		return nil, err
	}

	// Create Kafka event payload
	// eventPayload := map[string]interface{}{
	// 	"loan_id":      loan.ID,
	// 	"user_id":      loan.UserID,
	// 	"loan_amount":  loan.LoanAmount,
	// 	"loan_purpose": loan.LoanPurpose,
	// 	"status":       loan.ApplicationStatus,
	// 	"timestamp":    time.Now().Format(time.RFC3339),
	// }
	//payloadJSON, _ := json.Marshal(eventPayload)

	// Send Kafka event
	//kafka.ProduceEvent("LoanApplicationSubmitted", string(payloadJSON))

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

	// Fetch the loan application details by its ID
	if err := s.DB.First(&loan, req.LoanId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	//kafka.ProduceEvent("LoanStatusChecked", fmt.Sprintf(`{"loan_id":%d,"status":"%s"}`, loan.ID, loan.ApplicationStatus))

	// Prepare and return the loan application response
	return &pb.LoanApplicationResponse{
		LoanId:                  uint64(loan.ID),
		UserId:                  uint64(loan.UserID),
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
	if err := s.DB.First(&loan, req.LoanApplicationId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "Loan application not found")
		}
		return nil, err
	}

	// Update the status and reasoning
	loan.ApplicationStatus = req.NewStatus
	loan.Reasoning = req.Reasoning

	if err := s.DB.Save(&loan).Error; err != nil {
		return nil, err
	}

	return &pb.UpdateApplicationStatusResponse{
		Status: "Loan status updated successfully",
	}, nil
}

// GetAllLoanApplications retrieves all loan applications
func (s *LoanServiceServer) GetAllLoanApplications(ctx context.Context, req *pb.Empty) (*pb.LoanApplicationList, error) {
	var loans []models.LoanApplication

	// Fetch all loan applications
	if err := s.DB.Find(&loans).Error; err != nil {
		return nil, err
	}

	var responses []*pb.LoanApplicationResponse
	for _, loan := range loans {
		responses = append(responses, &pb.LoanApplicationResponse{
			LoanId:                  uint64(loan.ID),
			UserId:                  uint64(loan.UserID),
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
