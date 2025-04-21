package kafka

import (
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"AI-Powered-Automated-Loan-Underwriting-System/repositories"
	"context"
	"encoding/json"
	"log"
	//"time"

	"github.com/segmentio/kafka-go"
)

func StartEventConsumer(broker, topic string, repo *repositories.EventRepo) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		GroupID:  "event-consumer-group",
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	log.Println("Kafka event consumer started...")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			continue
		}

		var event KafkaEvent
		if err := json.Unmarshal(m.Value, &event); err != nil {
			log.Printf("Error unmarshaling Kafka message: %v\n", err)
			continue
		}

		// Save event to DB using repo
		err = repo.CreateEvent(context.Background(), models.Event{
			EventType: event.EventType,
			Payload:   event.Payload,
			Timestamp: event.Timestamp,
		})
		if err != nil {
			log.Printf("Failed to persist event: %v\n", err)
		} else {
			log.Printf("Event persisted: %s\n", event.EventType)
		}
	}
}
