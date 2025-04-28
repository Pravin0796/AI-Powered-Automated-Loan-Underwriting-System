package mockdata

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"AI-Powered-Automated-Loan-Underwriting-System/models"
)

func batchInsert[T any](db *gorm.DB, data []T, batchSize int, label string) error {
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}

		chunk := data[i:end]
		if err := db.Create(&chunk).Error; err != nil {
			return fmt.Errorf("failed to insert batch for %s: %w", label, err)
		}
	}
	return nil
}

func SeedMockData(db *gorm.DB) error {
	rand.Seed(time.Now().UnixNano())
	gofakeit.Seed(time.Now().UnixNano())

	const mockDataSize = 10000
	const batchSize = 1000

	// Step 1: Generate Users
	var users []models.User
	for i := 0; i < mockDataSize; i++ {
		users = append(users, models.User{
			FullName: gofakeit.Name(),
			Email:    gofakeit.Email(),
			Password: gofakeit.Password(true, true, true, true, false, 12),
			Phone:    gofakeit.Phone(),
			DateOfBirth: gofakeit.DateRange(
				time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2003, 12, 31, 0, 0, 0, 0, time.UTC),
			),
			Address:     gofakeit.Address().Address,
			CreditScore: gofakeit.Number(500, 850),
		})
	}
	if err := batchInsert(db, users, batchSize, "users"); err != nil {
		return err
	}

	// Re-fetch inserted Users
	var persistedUsers []models.User
	if err := db.Find(&persistedUsers).Error; err != nil {
		return fmt.Errorf("failed to fetch inserted users: %w", err)
	}

	userMap := make(map[uint]models.User)
	for _, user := range persistedUsers {
		userMap[user.ID] = user
	}

	// Step 2: Generate LoanApplications
	var applications []models.LoanApplication
	for _, user := range persistedUsers {
		loanAmount := gofakeit.Price(5000, 100000)
		grossIncome := gofakeit.Price(3000, 10000)
		debtPayment := gofakeit.Price(500, 3000)
		dti := debtPayment / grossIncome

		applications = append(applications, models.LoanApplication{
			UserID:                  user.ID,
			SSN:                     gofakeit.SSN(),
			AddressArea:             gofakeit.RandomString([]string{"urban", "rural"}),
			LoanAmount:              loanAmount,
			LoanPurpose:             gofakeit.RandomString([]string{"home", "car", "education", "business"}),
			EmploymentStatus:        gofakeit.RandomString([]string{"employed", "self-employed", "unemployed"}),
			GrossMonthlyIncome:      grossIncome,
			TotalMonthlyDebtPayment: debtPayment,
			DTIRatio:                dti,
			ApplicationStatus:       "PENDING",
			CreditReportFetched:     false,
			ExperianRequestID:       "",
			CreditScore:             0,
			Reasoning:               "",
		})
	}
	if err := batchInsert(db, applications, batchSize, "loan applications"); err != nil {
		return err
	}

	// Re-fetch LoanApplications
	var persistedApps []models.LoanApplication
	if err := db.Find(&persistedApps).Error; err != nil {
		return fmt.Errorf("failed to fetch applications: %w", err)
	}

	// Step 3: Generate CreditReports and Update LoanApplications
	reportMap := make(map[uint]models.CreditReport)
	for _, app := range persistedApps {
		reportData, _ := json.Marshal(map[string]interface{}{
			"tradelines": []map[string]interface{}{
				{
					"accountType":   gofakeit.RandomString([]string{"credit card", "auto loan", "mortgage"}),
					"balance":       gofakeit.Price(100, 10000),
					"creditLimit":   gofakeit.Price(1000, 20000),
					"paymentStatus": gofakeit.RandomString([]string{"current", "late", "default"}),
					"openedDate":    gofakeit.Date().Format("2006-01-02"),
				},
			},
		})

		fraudIndicators, _ := json.Marshal(map[string]interface{}{
			"syntheticIdentity": gofakeit.Bool(),
			"multipleSSN":       gofakeit.Bool(),
		})

		fakeCreditScore := gofakeit.Number(500, 850)

		report := models.CreditReport{
			UserID:            app.UserID,
			LoanApplicationID: app.ID,
			ReportData:        datatypes.JSON(reportData),
			CreditScore:       fakeCreditScore,
			FraudIndicators:   datatypes.JSON(fraudIndicators),
			DelinquencyFlag:   gofakeit.Bool(),
		}
		reportMap[app.ID] = report

		// Update LoanApplication with fake Experian details
		if err := db.Model(&models.LoanApplication{}).Where("id = ?", app.ID).Updates(map[string]interface{}{
			"credit_report_fetched": true,
			"experian_request_id":   gofakeit.UUID(),
			"credit_score":          fakeCreditScore,
		}).Error; err != nil {
			return fmt.Errorf("failed to update loan application after experian generation: %w", err)
		}
	}

	// Insert CreditReports
	var reports []models.CreditReport
	for _, r := range reportMap {
		reports = append(reports, r)
	}
	if err := batchInsert(db, reports, batchSize, "credit reports"); err != nil {
		return err
	}

	// Step 4: Generate LoanDecisions and Update LoanApplications
	for _, app := range persistedApps {
		user := userMap[app.UserID]
		report := reportMap[app.ID]

		approve := false
		reason := ""
		if user.CreditScore >= 650 && app.DTIRatio <= 0.35 && !report.DelinquencyFlag {
			approve = true
			reason = "Meets all criteria: High credit score, low DTI, no delinquency"
		} else {
			reason = fmt.Sprintf("CreditScore=%d, DTI=%.2f, Delinquent=%v", user.CreditScore, app.DTIRatio, report.DelinquencyFlag)
		}

		// Insert LoanDecision
		decision := models.LoanDecision{
			LoanApplicationID: app.ID,
			AiDecision:        approve,
			Reasoning:         reason,
			CreatedAt:         time.Now(),
		}
		if err := db.Create(&decision).Error; err != nil {
			return fmt.Errorf("failed to insert loan decision: %w", err)
		}

		// Update LoanApplication with Decision Status
		status := "REJECTED"
		if approve {
			status = "APPROVED"
		}

		if err := db.Model(&models.LoanApplication{}).Where("id = ?", app.ID).Updates(map[string]interface{}{
			"application_status": status,
			"reasoning":           reason,
		}).Error; err != nil {
			return fmt.Errorf("failed to update loan application after decision: %w", err)
		}
	}

	fmt.Println("✅ Successfully seeded mock data and updated loan applications.")
	return nil
}
