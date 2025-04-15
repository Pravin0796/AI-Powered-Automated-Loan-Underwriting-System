package services

import (
	"context"

	protos "AI-Powered-Automated-Loan-Underwriting-System/created_proto/notification"
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	protoNotifications := make([]*protos.Notification, len(notifications))
	for i, notif := range notifications {
		protoNotifications[i] = &protos.Notification{
			Id:        uint64(notif.ID),
			UserId:    uint64(notif.UserID),
			Title:     notif.Title,
			Message:   notif.Message,
			Type:      notif.Type,
			IsRead:    notif.IsRead,
			CreatedAt: timestamppb.New(notif.CreatedAt),
			UpdatedAt: timestamppb.New(notif.UpdatedAt),
		}
	}

	return &protos.UserNotificationResponse{
		Notifications: protoNotifications,
	}, nil
}

// MarkNotificationRead marks a notification as read
func (s *NotificationService) MarkNotificationRead(ctx context.Context, req *protos.MarkReadRequest) (*protos.MarkReadResponse, error) {
	var notification models.Notification
	if err := s.DB.First(&notification, req.NotificationId).Error; err != nil {
		return nil, err
	}

	notification.IsRead = true
	if err := s.DB.Save(&notification).Error; err != nil {
		return nil, err
	}

	return &protos.MarkReadResponse{Status: "Notification marked as read"}, nil
}
