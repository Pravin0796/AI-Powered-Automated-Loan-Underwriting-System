package migration

import (
	"log"

	"gorm.io/gorm"

	"AI-Powered-Automated-Loan-Underwriting-System/models"
)

// MigrateDatabase runs the database migrations
func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.LoanApplication{},
		&models.CreditReport{},
		&models.LoanDecision{},
		&models.LoanPayment{},
		&models.Event{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Database migrated successfully!")
}
