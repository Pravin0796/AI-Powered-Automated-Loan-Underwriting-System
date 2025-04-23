package kafka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

// KafkaProducer wraps the Kafka writer.
type KafkaProducer struct {
	writer *kafka.Writer
	topic  string
}

// NewProducer creates a new Kafka producer for the specified topic.
func NewProducer(broker, topic string) *KafkaProducer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	return &KafkaProducer{
		writer: writer,
		topic:  topic,
	}
}

// SendMessage sends a message to Kafka.
func (kp *KafkaProducer) SendMessage(v interface{}) error {
	value, err := json.Marshal(v)
	if err != nil {
		return err
	}
	msg := kafka.Message{
		Key:   []byte(time.Now().Format(time.RFC3339Nano)),
		Value: value,
	}
	return kp.writer.WriteMessages(context.Background(), msg)
}
