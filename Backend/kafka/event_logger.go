package kafka

import (
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"AI-Powered-Automated-Loan-Underwriting-System/repositories"
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

// type KafkaEvent struct {
// 	EventType string    `json:"event_type"`
// 	Payload   string    `json:"payload"` // already in JSON string format
// 	Timestamp time.Time `json:"timestamp"`
// }

func StartEventLoggerConsumer(broker, topic string, repo *repositories.EventRepo) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: "event-logger",
	})
	log.Printf("Event logger started for topic: %s", topic)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Read error:", err)
			continue
		}

		var evt models.Event
		if err := json.Unmarshal(msg.Value, &evt); err != nil {
			log.Println("Unmarshal error:", err)
			continue
		}

		err = repo.CreateEvent(context.Background(), models.Event{
			EventType: evt.EventType,
			Payload:   evt.Payload, // This can be stored as RawMessage in DB
			Timestamp: evt.Timestamp,
		})
		if err != nil {
			log.Println("Event DB save failed:", err)
		} else {
			log.Printf("Logged event: %s", evt.EventType)
		}
	}
}
