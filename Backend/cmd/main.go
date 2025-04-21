package main

import (
	//"AI-Powered-Automated-Loan-Underwriting-System/cache"
	"AI-Powered-Automated-Loan-Underwriting-System/config"
	"AI-Powered-Automated-Loan-Underwriting-System/migration"

	//"AI-Powered-Automated-Loan-Underwriting-System/mockdata"

	"AI-Powered-Automated-Loan-Underwriting-System/routes"
	//"fmt"
)

func main() {
	// Initialize database connection
	config.ConnectDatabase()

	// Run migrations
	migration.MigrateDatabase(config.DB)

	//eventRepo := repositories.NewEventRepo(db)

	// Start Kafka consumer in a goroutine
	//go kafka.StartEventConsumer("localhost:9092", "loan-events", eventRepo)
	//
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
