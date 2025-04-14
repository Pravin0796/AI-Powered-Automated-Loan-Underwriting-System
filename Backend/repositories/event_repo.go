package repositories

import (
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"context"
	"gorm.io/gorm"
)

type EventRepo struct {
	DB *gorm.DB
}

func NewEventRepo(DB *gorm.DB) *EventRepo {
	return &EventRepo{DB: DB}
}

func (r *EventRepo) CreateEvent(ctx context.Context, event models.Event) error {
	return r.DB.WithContext(ctx).Create(&event).Error
}

func (r *EventRepo) GetEventsByType(ctx context.Context, eventType string, events *[]models.Event) error {
	return r.DB.WithContext(ctx).Where("event_type = ?", eventType).Find(&events).Error
}

func (r *EventRepo) GetAllEvents(ctx context.Context, events *[]models.Event) error {
	return r.DB.WithContext(ctx).Find(&events).Error
}
