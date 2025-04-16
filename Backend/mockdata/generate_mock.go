package mockdata

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"AI-Powered-Automated-Loan-Underwriting-System/models" // Replace with your actual module path
)

func SeedMockData(db *gorm.DB) error {
	rand.Seed(time.Now().UnixNano())
	gofakeit.Seed(0)

	var mockdata = 5000
	// Step 1: Generate Users
	var users []models.User
	for i := 0; i < mockdata; i++ {
		user := models.User{
			FullName:    gofakeit.Name(),
			Email:       gofakeit.Email(),
			Password:    gofakeit.Password(true, true, true, true, false, 12),
			Phone:       gofakeit.Phone(),
			DateOfBirth: gofakeit.DateRange(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2003, 12, 31, 0, 0, 0, 0, time.UTC)),
			Address:     gofakeit.Address().Address,
			CreditScore: gofakeit.Number(300, 850),
		}
		users = append(users, user)
	}
	if err := db.Create(&users).Error; err != nil {
		return err
	}

	userMap := make(map[uint]models.User)
	for _, user := range users {
		userMap[user.ID] = user
	}

	// Step 2: Generate LoanApplications
	var applications []models.LoanApplication
	for _, user := range users {
		loanAmount := gofakeit.Price(5000, 100000)
		grossIncome := gofakeit.Price(3000, 10000)
		debtPayment := gofakeit.Price(500, 3000)
		dti := (debtPayment / grossIncome)

		app := models.LoanApplication{
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
		}
		applications = append(applications, app)
	}
	if err := db.Create(&applications).Error; err != nil {
		return err
	}

	// Step 3: Generate CreditReports
	var reports []models.CreditReport
	reportMap := make(map[uint]models.CreditReport)

	for _, app := range applications {
		reportData := map[string]interface{}{
			"tradelines": []map[string]interface{}{
				{
					"accountType":   gofakeit.RandomString([]string{"credit card", "auto loan", "mortgage"}),
					"balance":       gofakeit.Price(100, 10000),
					"creditLimit":   gofakeit.Price(1000, 20000),
					"paymentStatus": gofakeit.RandomString([]string{"current", "late", "default"}),
					"openedDate":    gofakeit.Date().Format("2006-01-02"),
				},
			},
			"publicRecords": []map[string]interface{}{
				{
					"type":      gofakeit.RandomString([]string{"bankruptcy", "tax lien", "none"}),
					"filedDate": gofakeit.Date().Format("2006-01-02"),
					"status":    gofakeit.RandomString([]string{"active", "discharged", "released"}),
				},
			},
			"inquiries": []map[string]interface{}{
				{
					"inquiredDate": gofakeit.Date().Format("2006-01-02"),
					"inquirer":     gofakeit.Company(),
					"type":         gofakeit.RandomString([]string{"hard", "soft"}),
				},
			},
			"utilization": map[string]interface{}{
				"totalLimit": gofakeit.Price(10000, 50000),
				"totalUsed":  gofakeit.Price(1000, 25000),
			},
		}

		reportJSON, _ := json.Marshal(reportData)
		fraudIndicators, _ := json.Marshal(map[string]interface{}{
			"syntheticIdentity": gofakeit.Bool(),
			"multipleSSN":       gofakeit.Bool(),
		})

		report := models.CreditReport{
			UserID:            app.UserID,
			LoanApplicationID: app.ID,
			ReportData:        datatypes.JSON(reportJSON),
			CreditScore:       app.CreditScore,
			FraudIndicators:   datatypes.JSON(fraudIndicators),
			DelinquencyFlag:   gofakeit.Bool(),
		}
		reports = append(reports, report)
		reportMap[app.ID] = report
	}
	if err := db.Create(&reports).Error; err != nil {
		return err
	}

	// Step 4: Generate LoanPayments and compute stats
	type PaymentStats struct {
		NumExistingLoans int
		NumLatePayments  int
	}
	paymentStats := make(map[uint]PaymentStats)
	var payments []models.LoanPayment

	for _, app := range applications {
		numPayments := gofakeit.Number(1, 5)
		lateCount := 0

		for i := 0; i < numPayments; i++ {
			status := gofakeit.RandomString([]string{"on-time", "late", "missed"})
			if status == "late" {
				lateCount++
			}
			payments = append(payments, models.LoanPayment{
				LoanApplicationID: app.ID,
				AmountPaid:        gofakeit.Price(100, 500),
				PaymentDate:       gofakeit.Date(),
				Status:            status,
			})
		}

		stats := paymentStats[app.UserID]
		stats.NumExistingLoans += 1
		stats.NumLatePayments += lateCount
		paymentStats[app.UserID] = stats
	}
	if err := db.Create(&payments).Error; err != nil {
		return err
	}

	// Step 5: Generate LoanDecisions using prediction logic
	var decisions []models.LoanDecision

	for _, app := range applications {
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
	if err := db.Create(&decisions).Error; err != nil {
		return err
	}

	fmt.Println("✅ Mock data seeded successfully.")
	return nil
}
