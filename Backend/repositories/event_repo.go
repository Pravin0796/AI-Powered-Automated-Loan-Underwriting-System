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

func (r *EventRepo) GetAllEvents(ctx context.Context, events *[]models.Event) error {
	return r.DB.WithContext(ctx).Find(&events).Error
}

func (r *EventRepo) GetEventByID(ctx context.Context, eventID uint, event *models.Event) error {
	return r.DB.WithContext(ctx).First(&event, eventID).Error
}

func (r *EventRepo) UpdateEvent(ctx context.Context, event models.Event) error {
	return r.DB.WithContext(ctx).Save(&event).Error
}

func (r *EventRepo) DeleteEvent(ctx context.Context, eventID uint) error {
	return r.DB.WithContext(ctx).Delete(&models.Event{}, eventID).Error
}
