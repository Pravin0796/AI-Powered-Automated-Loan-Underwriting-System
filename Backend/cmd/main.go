package main

import (
	"AI-Powered-Automated-Loan-Underwriting-System/cache"
	"AI-Powered-Automated-Loan-Underwriting-System/config"
	"AI-Powered-Automated-Loan-Underwriting-System/experian"
	"AI-Powered-Automated-Loan-Underwriting-System/migration"
	"AI-Powered-Automated-Loan-Underwriting-System/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	config.ConnectDatabase()

	// Fetch credit report
	//creditProfile, err := experian.FetchCreditProfile(
	//	"123-45-6789", "Doe", "John", "123 Main St", "New York", "NY", "10001",
	//)
	//if err != nil {
	//	fmt.Println("Error fetching credit report:", err)
	//	return
	//}
	//fmt.Println("Credit Report:", creditProfile)

	//token, err := experian.GetAccessToken()
	//if err != nil {
	//	fmt.Println("Error fetching access token:", err)
	//} else {
	//	fmt.Println("Access Token:", token)
	//}

	routes.StartGRPCServer()

	fmt.Println("Fetching Credit Report...")
	experian.FetchCreditReport()

	// Initialize Redis connection
	cache.ConnectRedis()

	// Run migrations
	migration.MigrateDatabase(config.DB)

	r := gin.Default()

	r.Run(":8080")
}
