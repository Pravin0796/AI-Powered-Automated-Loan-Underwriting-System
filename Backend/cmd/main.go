package main

import (
	//"AI-Powered-Automated-Loan-Underwriting-System/cache"
	"AI-Powered-Automated-Loan-Underwriting-System/config"
	"AI-Powered-Automated-Loan-Underwriting-System/experian"
	"AI-Powered-Automated-Loan-Underwriting-System/kafka"
	"AI-Powered-Automated-Loan-Underwriting-System/migration"
	"AI-Powered-Automated-Loan-Underwriting-System/repositories"
	"fmt"

	//"AI-Powered-Automated-Loan-Underwriting-System/mockdata"

	"AI-Powered-Automated-Loan-Underwriting-System/routes"
	//"fmt"
)

func testMockExperianAPI() {
	fmt.Println("Calling Mock Experian API...")

	req := experian.ExperianRequest{
		SSN:              "123-45-6789",
		UserID:           "user-001",
		LoanAmount:       25000.0,
		EmploymentStatus: "full-time",
	}

	report, err := experian.FetchMockCreditReport(req, "loan-app-id-123")
	if err != nil {
		fmt.Printf("Error fetching credit report: %v\n", err)
		return
	}

	fmt.Printf("Credit Score: %d\n", report.CreditScore)
	fmt.Printf("Delinquency Flag: %v\n", report.DelinquencyFlag)
	fmt.Printf("Fraud Indicators: %s\n", report.FraudIndicators)
	fmt.Printf("Report Data: %s\n", report.ReportData)
}

func main() {
	// Initialize database connection
	config.ConnectDatabase()

	// Run migrations
	migration.MigrateDatabase(config.DB)

	go kafka.StartEventLoggerConsumer("localhost:9092", "loan-events", repositories.NewEventRepo(config.DB))

	// 🔹 Call Experian mock API test
	testMockExperianAPI()

	// err := mockdata.SeedMockData(config.DB)
	// if err != nil {
	// 	return
	// }
	//decision, err := services.GetLoanDecision()
	//if err != nil {
	//	return
	//}
	//fmt.Println(decision)

	//fmt.Println("Fetching Credit Report...")
	//experian.FetchCreditReport()

	// Initialize Redis connection
	//cache.ConnectRedis()
	routes.StartGRPCServer(config.DB)
}
