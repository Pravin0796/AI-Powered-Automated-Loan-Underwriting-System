package services

import (
	"AI-Powered-Automated-Loan-Underwriting-System/repositories"
	"context"

	protos "AI-Powered-Automated-Loan-Underwriting-System/created_proto/notification"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NotificationService struct {
	protos.UnimplementedNotificationServiceServer
	Repo *repositories.NotificationRepo
}

// GetUserNotifications fetches notifications for a user
func (s *NotificationService) GetUserNotifications(ctx context.Context, req *protos.UserNotificationRequest) (*protos.UserNotificationResponse, error) {
	notifications, err := s.Repo.GetUserNotifications(ctx, uint(req.UserId))
	if err != nil {
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

	return &protos.UserNotificationResponse{Notifications: protoNotifications}, nil
}

// MarkNotificationRead marks a notification as read
func (s *NotificationService) MarkNotificationRead(ctx context.Context, req *protos.MarkReadRequest) (*protos.MarkReadResponse, error) {
	if err := s.Repo.MarkNotificationRead(ctx, uint(req.NotificationId)); err != nil {
		return nil, err
	}
	return &protos.MarkReadResponse{Status: "Notification marked as read"}, nil
}
