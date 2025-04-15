package main

import (
	"AI-Powered-Automated-Loan-Underwriting-System/cache"
	"AI-Powered-Automated-Loan-Underwriting-System/config"
	"AI-Powered-Automated-Loan-Underwriting-System/migration"
	"AI-Powered-Automated-Loan-Underwriting-System/mockdata"
	"AI-Powered-Automated-Loan-Underwriting-System/routes"
	"fmt"
	"log"
)

func main() {
	// Initialize database connection
	config.ConnectDatabase()

	// Run migrations
	migration.MigrateDatabase(config.DB)

	// Step 2: Backfill NULL SSNs before running migration
	if err := config.DB.Exec(`UPDATE loan_applications SET ssn = '000-00-0000' WHERE ssn IS NULL`).Error; err != nil {
		log.Fatalf("❌ Failed to backfill SSN: %v", err)
	} else {
		fmt.Println("✅ Backfilled NULL SSNs successfully")
	}

	//
	mockdata.GenerateMockData()

	fmt.Println("Fetching Credit Report...")
	//experian.FetchCreditReport()

	// Initialize Redis connection
	cache.ConnectRedis()
	routes.StartGRPCServer(config.DB)
}
