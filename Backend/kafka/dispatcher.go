package kafka

import (
	"context"
	"log"
	"encoding/json"
	"AI-Powered-Automated-Loan-Underwriting-System/models"
)

// Dispatch is a generic method for dispatching messages to handlers.
func (kc *KafkaConsumer) Dispatch(ctx context.Context, handler func([]byte)) {
	for {
		m, err := kc.Reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Kafka read error: %v", err)
			continue
		}
		handler(m.Value)
	}
}

// ListenLoanApplication listens for LoanApplications and triggers the handler.
func (kc *KafkaConsumer) ListenLoanApplication(handler func(models.LoanApplication)) {
	kc.Dispatch(context.Background(), func(value []byte) {
		var app models.LoanApplication
		if err := json.Unmarshal(value, &app); err != nil {
			log.Println("Failed to parse loan application:", err)
			return
		}
		handler(app)
	})
}

// ListenLoanProcessed listens for processed loan decisions and triggers the handler.
func (kc *KafkaConsumer) ListenLoanProcessed(handler func(models.LoanDecision)) {
	kc.Dispatch(context.Background(), func(value []byte) {
		var decision models.LoanDecision
		if err := json.Unmarshal(value, &decision); err != nil {
			log.Println("Failed to parse loan decision:", err)
			return
		}
		handler(decision)
	})
}
