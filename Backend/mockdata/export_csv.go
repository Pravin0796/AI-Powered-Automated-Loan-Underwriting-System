package mockdata

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"time"

	"AI-Powered-Automated-Loan-Underwriting-System/models"

	"gorm.io/gorm"
)

// ExportToCSV is a utility function to export data to a CSV file
func ExportToCSV(filename string, data interface{}, headers []string) error {
	dir := "../../ml_model/csv_files"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	pathToCSV := path.Join(dir, filename)
	file, err := os.Create(pathToCSV)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("failed to write headers to CSV: %w", err)
	}

	// Write data
	switch v := data.(type) {
	case []models.User:
		for _, record := range v {
			if err := writer.Write([]string{
				fmt.Sprintf("%d", record.ID),
				record.FullName,
				record.Email,
				record.Phone,
				record.DateOfBirth.Format(time.RFC3339),
				record.Address,
				fmt.Sprintf("%d", record.CreditScore),
			}); err != nil {
				return fmt.Errorf("failed to write user record to CSV: %w", err)
			}
		}
	case []models.LoanApplication:
		for _, record := range v {
			if err := writer.Write([]string{
				fmt.Sprintf("%d", record.ID),
				fmt.Sprintf("%d", record.UserID),
				record.SSN,
				record.AddressArea,
				fmt.Sprintf("%f", record.LoanAmount),
				record.LoanPurpose,
				record.EmploymentStatus,
				fmt.Sprintf("%f", record.GrossMonthlyIncome),
				fmt.Sprintf("%f", record.TotalMonthlyDebtPayment),
				fmt.Sprintf("%f", record.DTIRatio),
				record.ApplicationStatus,
				record.Reasoning,
				fmt.Sprintf("%t", record.CreditReportFetched),
				record.ExperianRequestID,
				fmt.Sprintf("%d", record.CreditScore),
			}); err != nil {
				return fmt.Errorf("failed to write loan application record to CSV: %w", err)
			}
		}
	case []models.CreditReport:
		for _, record := range v {
			reportData := string(record.ReportData)
			fraudIndicators := string(record.FraudIndicators)
			if err := writer.Write([]string{
				fmt.Sprintf("%d", record.ID),
				fmt.Sprintf("%d", record.UserID),
				fmt.Sprintf("%d", record.LoanApplicationID),
				reportData,
				fmt.Sprintf("%d", record.CreditScore),
				fraudIndicators,
				fmt.Sprintf("%t", record.DelinquencyFlag),
			}); err != nil {
				return fmt.Errorf("failed to write credit report record to CSV: %w", err)
			}
		}
	case []models.LoanPayment:
		for _, record := range v {
			if err := writer.Write([]string{
				fmt.Sprintf("%d", record.ID),
				fmt.Sprintf("%d", record.LoanApplicationID),
				fmt.Sprintf("%.2f", record.AmountPaid),
				record.PaymentDate.Format(time.RFC3339),
				record.DueDate.Format(time.RFC3339),
				record.Status,
			}); err != nil {
				return fmt.Errorf("failed to write loan payment record to CSV: %w", err)
			}
		}
	case []models.LoanDecision:
		for _, record := range v {
			if err := writer.Write([]string{
				fmt.Sprintf("%d", record.ID),
				fmt.Sprintf("%d", record.LoanApplicationID),
				fmt.Sprintf("%t", record.AiDecision),
				record.Reasoning,
				record.CreatedAt.Format(time.RFC3339),
			}); err != nil {
				return fmt.Errorf("failed to write loan decision record to CSV: %w", err)
			}
		}
	default:
		return fmt.Errorf("unsupported data type %T", v)
	}

	return nil
}

// ExportMockDataToCSV will export the mock data to CSV files
func ExportMockDataToCSV(db *gorm.DB) error {
	// Step 1: Fetch data from the database
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return fmt.Errorf("failed to fetch users: %w", err)
	}
	if err := ExportToCSV("users.csv", users, []string{"ID", "FullName", "Email", "Phone", "DateOfBirth", "Address", "CreditScore"}); err != nil {
		return err
	}

	var loanApplications []models.LoanApplication
	if err := db.Find(&loanApplications).Error; err != nil {
		return fmt.Errorf("failed to fetch loan applications: %w", err)
	}
	if err := ExportToCSV("loan_applications.csv", loanApplications, []string{"ID", "UserID", "SSN", "AddressArea", "LoanAmount", "LoanPurpose", "EmploymentStatus", "GrossMonthlyIncome", "TotalMonthlyDebtPayment", "DTIRatio", "ApplicationStatus", "Reasoning", "CreditReportFetched", "ExperianRequestID", "CreditScore"}); err != nil {
		return err
	}

	var creditReports []models.CreditReport
	if err := db.Find(&creditReports).Error; err != nil {
		return fmt.Errorf("failed to fetch credit reports: %w", err)
	}
	if err := ExportToCSV("credit_reports.csv", creditReports, []string{"ID", "UserID", "LoanApplicationID", "ReportData", "CreditScore", "FraudIndicators", "DelinquencyFlag"}); err != nil {
		return err
	}

	var loanPayments []models.LoanPayment
	if err := db.Find(&loanPayments).Error; err != nil {
		return fmt.Errorf("failed to fetch loan payments: %w", err)
	}

	// Now also export DueDate
	if err := ExportToCSV("loan_payments.csv", loanPayments, []string{
		"ID", "LoanApplicationID", "AmountPaid", "PaymentDate", "DueDate", "Status",
	}); err != nil {
		return err
	}

	var loanDecisions []models.LoanDecision
	if err := db.Find(&loanDecisions).Error; err != nil {
		return fmt.Errorf("failed to fetch loan decisions: %w", err)
	}
	if err := ExportToCSV("loan_decisions.csv", loanDecisions, []string{"ID", "LoanApplicationID", "AiDecision", "Reasoning", "CreatedAt"}); err != nil {
		return err
	}

	fmt.Println("✅ Successfully exported mock data to CSV files.")
	return nil
}
