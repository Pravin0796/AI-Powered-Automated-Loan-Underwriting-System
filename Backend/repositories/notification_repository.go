package repositories

import (
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"context"

	"gorm.io/gorm"
)

type NotificationRepo struct {
	DB *gorm.DB
}

func NewNotificationRepo(db *gorm.DB) *NotificationRepo {
	return &NotificationRepo{DB: db}
}

// GetUserNotifications fetches all notifications for a specific user
func (r *NotificationRepo) GetUserNotifications(ctx context.Context, userID uint) ([]models.Notification, error) {
	var notifications []models.Notification
	err := r.DB.WithContext(ctx).Where("user_id = ?", userID).Order("created_at DESC").Find(&notifications).Error
	return notifications, err
}

// MarkNotificationRead sets IsRead to true for a specific notification
func (r *NotificationRepo) MarkNotificationRead(ctx context.Context, notificationID uint) error {
	return r.DB.WithContext(ctx).Model(&models.Notification{}).
		Where("id = ?", notificationID).
		Update("is_read", true).Error
}

// CreateNotification adds a new notification
func (r *NotificationRepo) CreateNotification(ctx context.Context, notification *models.Notification) error {
	return r.DB.WithContext(ctx).Create(notification).Error
}

// DeleteNotification deletes a notification (soft delete)
func (r *NotificationRepo) DeleteNotification(ctx context.Context, notificationID uint) error {
	return r.DB.WithContext(ctx).Delete(&models.Notification{}, notificationID).Error
}
