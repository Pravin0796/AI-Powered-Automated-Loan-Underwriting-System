package mockdata

import (
	"AI-Powered-Automated-Loan-Underwriting-System/config"
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"errors"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GenerateMockData() {
	var users []models.User
	const numRecords = 20

	// 1. Generate and insert mock users
	for i := 0; i < numRecords; i++ {
		// Generate a random password
		pass := gofakeit.Password(true, true, true, true, false, 10)
		// hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		if err != nil {
			errors.New("failed to hash password")
		}

		user := models.User{
			FullName:    gofakeit.Name(),
			Email:       gofakeit.Email(),
			Password:    string(hashedPassword),
			Phone:       gofakeit.Phone(),
			DateOfBirth: gofakeit.Date(),
			Address:     gofakeit.Address().Address,
			CreditScore: gofakeit.Number(300, 850),
		}
		if err := config.DB.Create(&user).Error; err != nil {
			log.Println("Error inserting user:", err)
			continue
		}
		users = append(users, user)
		fmt.Println("Inserted user:", user.FullName)
	}

	// 2. Generate and insert mock loan applications, credit reports, and loan decisions per user
	for _, user := range users {
		loanApp := models.LoanApplication{
			UserID:            user.ID,
			SSN:               gofakeit.SSN(),
			Address:           user.Address,
			LoanAmount:        gofakeit.Float64Range(1000, 50000),
			LoanPurpose:       gofakeit.Word(),
			EmploymentStatus:  gofakeit.RandomString([]string{"Employed", "Self-employed", "Unemployed"}),
			AnnualIncome:      gofakeit.Float64Range(20000, 150000),
			DTIRatio:          gofakeit.Float64Range(10, 50),
			ApplicationStatus: "PENDING",
			Reasoning:         gofakeit.Sentence(10),
			CreditScore:       user.CreditScore,
		}
		if err := config.DB.Create(&loanApp).Error; err != nil {
			log.Println("Error inserting loan application:", err)
			continue
		}
		fmt.Println("Inserted loan application for user:", user.Email)

		creditReport := models.CreditReport{
			UserID:            user.ID,
			LoanApplicationID: loanApp.ID,
			ReportData:        []byte(fmt.Sprintf(`{"credit_score": %d}`, user.CreditScore)),
			CreditScore:       user.CreditScore,
			FraudIndicators:   []byte(`{"fraud_risk": 0.2}`),
			DelinquencyFlag:   gofakeit.Bool(),
		}
		if err := config.DB.Create(&creditReport).Error; err != nil {
			log.Println("Error inserting credit report:", err)
			continue
		}
		fmt.Println("Inserted credit report for user:", user.Email)

		loanDecision := models.LoanDecision{
			LoanApplicationID: loanApp.ID,
			AIDecision:        gofakeit.RandomString([]string{"approved", "rejected"}),
			Reasoning:         gofakeit.Sentence(5),
		}
		if err := config.DB.Create(&loanDecision).Error; err != nil {
			log.Println("Error inserting loan decision:", err)
			continue
		}
		fmt.Println("Inserted loan decision for application ID:", loanApp.ID)
	}

	fmt.Println("✅ All mock data inserted successfully with correct relationships!")
}
