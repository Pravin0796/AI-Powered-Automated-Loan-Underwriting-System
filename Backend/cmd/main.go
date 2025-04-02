package main

import (
	"AI-Powered-Automated-Loan-Underwriting-System/config"
	"AI-Powered-Automated-Loan-Underwriting-System/migration"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	config.ConnectDatabase()

	// Run migrations
	migration.MigrateDatabase(config.DB)

	r := gin.Default()

	r.Run(":8080")
}
