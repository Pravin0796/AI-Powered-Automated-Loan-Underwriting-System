package kafka

import (
	"context"
	"encoding/json"
	"log"

	"AI-Powered-Automated-Loan-Underwriting-System/models" // replace with your actual path

	"github.com/segmentio/kafka-go"
)

const (
	kafkaBroker = "localhost:9092"           // your Kafka broker
	kafkaTopic  = "LoanApplicationSubmitted" // your topic name
	kafkaGroup  = "loan-app-consumers"       // consumer group ID
)

func ConsumeLoanApplications() {
	// Create a Kafka reader with consumer group support
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		Topic:   kafkaTopic,
		GroupID: kafkaGroup,
		Logger:  kafka.LoggerFunc(log.Printf),
	})

	defer func() {
		if err := reader.Close(); err != nil {
			log.Printf("Error closing Kafka reader: %v", err)
		}
	}()

	log.Println("Listening for LoanApplicationSubmitted events using kafka-go...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading Kafka message: %v", err)
			continue
		}

		var event models.Event
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("Error unmarshalling Kafka message: %v", err)
			continue
		}

		if event.EventType == "LoanApplicationSubmitted" {
			handleLoanApplicationSubmitted(event)
		}
	}
}
