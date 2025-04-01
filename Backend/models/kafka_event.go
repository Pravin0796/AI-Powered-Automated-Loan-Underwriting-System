package models

import "time"

// Event represents a Kafka event in the system
type Event struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	EventType string    `gorm:"type:varchar(50);not null" json:"event_type"`
	Payload   string    `gorm:"type:jsonb;not null" json:"payload"`
	Timestamp time.Time `gorm:"autoCreateTime" json:"timestamp"`
}
