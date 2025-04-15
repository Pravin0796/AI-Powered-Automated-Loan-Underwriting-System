package services

import (
	"context"
	"time"

	"backend/models"
	"backend/protos"
	"gorm.io/gorm"
)

type NotificationService struct {
	protos.UnimplementedNotificationServiceServer
	DB *gorm.DB
}

// GetUserNotifications fetches notifications for a user
func (s *NotificationService) GetUserNotifications(ctx context.Context, req *protos.UserNotificationRequest) (*protos.UserNotificationResponse, error) {
	var notifications []models.Notification
	if err := s.DB.Where("user_id = ?", req.UserId).Find(&notifications).Error; err != nil {
		return nil, err
	}

	// Map models to proto message
	protoNotifications := make([]*protos.Notification, len(notifications))
	for i, notif := range notifications {
		protoNotifications[i] = &protos.Notification{
			Id:        uint64(notif.ID),
			EventType: notif.EventType,
			Payload:   notif.Payload,
			Timestamp: notif.Timestamp.Format(time.RFC3339),
			Read:      notif.Read,
		}
	}

	return &protos.UserNotificationResponse{Notifications: protoNotifications}, nil
}

// MarkNotificationRead marks a notification as read
func (s *NotificationService) MarkNotificationRead(ctx context.Context, req *protos.MarkReadRequest) (*protos.MarkReadResponse, error) {
	var notification models.Notification
	if err := s.DB.First(&notification, req.NotificationId).Error; err != nil {
		return nil, err
	}

	notification.Read = true
	if err := s.DB.Save(&notification).Error; err != nil {
		return nil, err
	}

	return &protos.MarkReadResponse{Status: "Notification marked as read"}, nil
}
