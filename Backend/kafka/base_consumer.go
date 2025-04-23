package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

// KafkaConsumer wraps the Kafka reader.
type KafkaConsumer struct {
	Reader *kafka.Reader
}

// NewConsumer creates a new Kafka consumer for the specified topic and group.
func NewConsumer(broker, topic, groupID string) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	return &KafkaConsumer{Reader: reader}
}

// StartListening starts the consumer and processes messages with a custom handler.
func (kc *KafkaConsumer) StartListening(handler func([]byte)) {
	for {
		m, err := kc.Reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Kafka read error: %v", err)
			continue
		}
		handler(m.Value)
	}
}
