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

		// ✅ Fix: Copy chunk to a temporary slice so we can take its address
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

	// Re-fetch user IDs to ensure they are populated
	var persistedUsers []models.User
	if err := db.Find(&persistedUsers).Error; err != nil {
		return fmt.Errorf("failed to fetch inserted users: %w", err)
	}

	userMap := make(map[uint]models.User)
	for _, user := range persistedUsers {
		userMap[user.ID] = user
	}

	// Step 2: LoanApplications
	var applications []models.LoanApplication
	for _, user := range persistedUsers {
		loanAmount := gofakeit.Price(5000, 100000)
		grossIncome := gofakeit.Price(3000, 10000)
		debtPayment := gofakeit.Price(500, 3000)
		dti := (debtPayment / grossIncome)

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
			Reasoning:               gofakeit.Sentence(10),
			CreditReportFetched:     true,
			ExperianRequestID:       gofakeit.UUID(),
			CreditScore:             user.CreditScore,
		})
	}
	if err := batchInsert(db, applications, batchSize, "loan applications"); err != nil {
		return err
	}

	// Fetch back applications with IDs
	var persistedApps []models.LoanApplication
	if err := db.Find(&persistedApps).Error; err != nil {
		return fmt.Errorf("failed to fetch applications: %w", err)
	}

	reportMap := make(map[uint]models.CreditReport)
	var reports []models.CreditReport
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

		report := models.CreditReport{
			UserID:            app.UserID,
			LoanApplicationID: app.ID,
			ReportData:        datatypes.JSON(reportData),
			CreditScore:       app.CreditScore,
			FraudIndicators:   datatypes.JSON(fraudIndicators),
			DelinquencyFlag:   gofakeit.Bool(),
		}
		reports = append(reports, report)
		reportMap[app.ID] = report
	}
	if err := batchInsert(db, reports, batchSize, "credit reports"); err != nil {
		return err
	}

	// Step 4: LoanPayments
	type PaymentStats struct {
		NumExistingLoans int
		NumLatePayments  int
	}

	paymentStats := make(map[uint]PaymentStats)
	var payments []models.LoanPayment

	for _, app := range persistedApps {
		numPayments := gofakeit.Number(1, 5) // Random number of payments per loan
		lateCount := 0

		for i := 0; i < numPayments; i++ {
			dueDate := gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now())
			paymentDelayDays := gofakeit.Number(-5, 15) // Can be paid early (-5 days) or late (+15 days)
			paymentDate := dueDate.AddDate(0, 0, paymentDelayDays)

			status := "successful"
			if paymentDelayDays > 0 {
				status = "failed" // late payments more likely to fail
				lateCount++
			} else if gofakeit.Bool() {
				status = "successful"
			} else {
				status = "pending"
			}

			payments = append(payments, models.LoanPayment{
				LoanApplicationID: app.ID,
				AmountPaid:        gofakeit.Price(100, 500),
				DueDate:           dueDate,
				PaymentDate:       paymentDate,
				Status:            status,
				CreatedAt:         gofakeit.Date(),
				UpdatedAt:         time.Now(),
			})
		}

		stats := paymentStats[app.UserID]
		stats.NumExistingLoans++
		stats.NumLatePayments += lateCount
		paymentStats[app.UserID] = stats
	}

	if err := batchInsert(db, payments, batchSize, "loan payments"); err != nil {
		return err
	}

	// Step 5: LoanDecisions
	var decisions []models.LoanDecision
	for _, app := range persistedApps {
		user := userMap[app.UserID]
		report := reportMap[app.ID]
		stats := paymentStats[app.UserID]

		approve := false
		reason := ""
		if user.CreditScore >= 650 && app.DTIRatio <= 0.35 && !report.DelinquencyFlag && stats.NumLatePayments <= 1 {
			approve = true
			reason = "Meets all criteria: High credit score, low DTI, good history"
		} else {
			reason = fmt.Sprintf("CreditScore=%d, DTI=%.2f, Delinquent=%v, LatePayments=%d",
				user.CreditScore, app.DTIRatio, report.DelinquencyFlag, stats.NumLatePayments)
		}

		decisions = append(decisions, models.LoanDecision{
			LoanApplicationID: app.ID,
			AiDecision:        approve,
			Reasoning:         reason,
			CreatedAt:         time.Now(),
		})
	}
	if err := batchInsert(db, decisions, batchSize, "loan decisions"); err != nil {
		return err
	}

	fmt.Println("✅ Successfully seeded 5000 mock data records.")
	return nil
}
