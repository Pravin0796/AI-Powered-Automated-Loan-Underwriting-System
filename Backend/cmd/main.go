package main

import (
	//"AI-Powered-Automated-Loan-Underwriting-System/cache"
	"AI-Powered-Automated-Loan-Underwriting-System/config"
	"AI-Powered-Automated-Loan-Underwriting-System/kafka"
	"AI-Powered-Automated-Loan-Underwriting-System/migration"
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"AI-Powered-Automated-Loan-Underwriting-System/repositories"

	//"AI-Powered-Automated-Loan-Underwriting-System/mockdata"

	"AI-Powered-Automated-Loan-Underwriting-System/routes"
	//"fmt"
)

func main() {
	// Initialize database connection
	config.ConnectDatabase()

	// Run migrations
	migration.MigrateDatabase(config.DB)

	// Define Kafka broker address
	broker := "localhost:9092" // Update this with the actual broker URL

	// Initialize EventRepo (you need to initialize db connection properly)
	// Assuming you have a `db` variable or connection to your database
	eventRepo := repositories.NewEventRepo(config.DB) // Ensure db is defined somewhere

	go func() {
		appConsumer := kafka.NewLoanApplicationConsumer(broker, "loan-app-group")
		appConsumer.ListenLoanApplication(func(app models.LoanApplication) {
			// trigger scoring, Experian mock, etc.
		})
	}()

	go kafka.StartEventLoggerConsumer(broker, "loan_application_submitted", eventRepo)
	go kafka.StartEventLoggerConsumer(broker, "experian_report_generated", eventRepo)
	go kafka.StartEventLoggerConsumer(broker, "loan_application_processed", eventRepo)
	go kafka.StartEventLoggerConsumer(broker, "loan_notification_triggered", eventRepo)

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
