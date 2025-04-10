package repositories

import (
	pb "AI-Powered-Automated-Loan-Underwriting-System/created_proto/user"
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"context"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) *UserRepo {
	return &UserRepo{DB: DB}
}

func (r *UserRepo) CreateUser(ctx context.Context, user models.User) error {
	// Implement the logic to register a user in the database
	return r.DB.WithContext(ctx).Create(&user).Error
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, req *pb.LoginRequest, user *models.User) error {
	// Implement the logic to register a user in the database
	return r.DB.WithContext(ctx).Where("email = ?", req.Email).First(&user).Error
}
