package kafka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
	//"log"
)

type KafkaEvent struct {
	EventType string    `json:"event_type"`
	Payload   string    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

func ProduceEvent(broker, topic string, event KafkaEvent) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
		Async:   false,
	})
	defer writer.Close()

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(event.EventType),
		Value: data,
	}

	return writer.WriteMessages(context.Background(), msg)
}
