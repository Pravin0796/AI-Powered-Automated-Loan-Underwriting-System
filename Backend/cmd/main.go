package main

import (
	"AI-Powered-Automated-Loan-Underwriting-System/cache"
	"AI-Powered-Automated-Loan-Underwriting-System/config"
	"AI-Powered-Automated-Loan-Underwriting-System/experian"
	"AI-Powered-Automated-Loan-Underwriting-System/migration"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	config.ConnectDatabase()

	// Fetch credit report
	creditProfile, err := experian.FetchCreditProfile(
		"123-45-6789", "Doe", "John", "123 Main St", "New York", "NY", "10001",
	)
	if err != nil {
		fmt.Println("Error fetching credit report:", err)
		return
	}
	fmt.Println("Credit Report:", creditProfile)

	// Initialize Redis connection
	cache.ConnectRedis()

	// Run migrations
	migration.MigrateDatabase(config.DB)

	r := gin.Default()

	r.Run(":8080")
}
