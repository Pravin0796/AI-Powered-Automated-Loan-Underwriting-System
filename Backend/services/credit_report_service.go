package services

import (
	"AI-Powered-Automated-Loan-Underwriting-System/created_proto/credit_report"
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type CreditReportService struct {
	DB *gorm.DB
}

func NewCreditReportService(db *gorm.DB) *CreditReportService {
	return &CreditReportService{DB: db}
}

func (s *CreditReportService) CreateCreditReport(ctx context.Context, req *credit_report.CreateCreditReportRequest) (*credit_report.CreateCreditReportResponse, error) {
	// Validate JSON
	var rawData map[string]interface{}
	if err := json.Unmarshal([]byte(req.ReportData), &rawData); err != nil {
		return nil, errors.New("invalid report_data JSON")
	}

	var fraudData map[string]interface{}
	if err := json.Unmarshal([]byte(req.FraudIndicators), &fraudData); err != nil {
		return nil, errors.New("invalid fraud_indicators JSON")
	}

	report := models.CreditReport{
		UserID:            uint(req.UserId),
		LoanApplicationID: uint(req.LoanApplicationId),
		ReportData:        datatypes.JSON(req.ReportData),
		CreditScore:       int(req.CreditScore),
		FraudIndicators:   datatypes.JSON(req.FraudIndicators),
		DelinquencyFlag:   false, // You can add logic to set this
	}

	if err := s.DB.Create(&report).Error; err != nil {
		return nil, err
	}

	return &credit_report.CreateCreditReportResponse{
		CreditReportId: uint64(report.ID),
		Status:         "Created successfully",
	}, nil
}

func (s *CreditReportService) GetCreditReport(ctx context.Context, req *credit_report.GetCreditReportRequest) (*credit_report.GetCreditReportResponse, error) {
	var report models.CreditReport
	if err := s.DB.First(&report, req.CreditReportId).Error; err != nil {
		return nil, err
	}

	return &credit_report.GetCreditReportResponse{
		Id:                uint64(report.ID),
		UserId:            uint64(report.UserID),
		LoanApplicationId: uint64(report.LoanApplicationID),
		ReportData:        string(report.ReportData),
		CreditScore:       int32(report.CreditScore),
		FraudIndicators:   string(report.FraudIndicators),
		GeneratedAt:       timestamppb.New(report.CreatedAt),
	}, nil
}
