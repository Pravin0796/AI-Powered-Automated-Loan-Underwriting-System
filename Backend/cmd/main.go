package main

import (
	"AI-Powered-Automated-Loan-Underwriting-System/cache"
	"AI-Powered-Automated-Loan-Underwriting-System/config"
	"AI-Powered-Automated-Loan-Underwriting-System/migration"
	"AI-Powered-Automated-Loan-Underwriting-System/mockdata"
	"AI-Powered-Automated-Loan-Underwriting-System/routes"
	"AI-Powered-Automated-Loan-Underwriting-System/services"
	"fmt"
	"log"
)

func main() {
	// Initialize database connection
	config.ConnectDatabase()

	// Run migrations
	migration.MigrateDatabase(config.DB)

	//
	mockdata.GenerateMockData()
	decision, err := services.GetLoanDecision()
	if err != nil {
		return
	}
	fmt.Println(decision)

	fmt.Println("Fetching Credit Report...")
	//experian.FetchCreditReport()

	// Initialize Redis connection
	cache.ConnectRedis()
	routes.StartGRPCServer(config.DB)
}
