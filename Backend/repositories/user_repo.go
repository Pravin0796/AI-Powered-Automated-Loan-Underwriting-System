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

func (r *UserRepo) GetUserByID(ctx context.Context, userID uint, user *models.User) error {
	return r.DB.WithContext(ctx).First(&user, userID).Error
}

func (r *UserRepo) UpdateUser(ctx context.Context, user *models.User) error {
	return r.DB.WithContext(ctx).Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (r *UserRepo) DeleteUser(ctx context.Context, userID uint) error {
	return r.DB.WithContext(ctx).Delete(&models.User{}, userID).Error
}

func (r *UserRepo) GetAllUsers(ctx context.Context, users *[]models.User) error {
	return r.DB.WithContext(ctx).Find(&users).Error
}

func (r *UserRepo) GetCreditScore(ctx context.Context, userID uint, user *models.User) error {
	return r.DB.WithContext(ctx).Raw("select credit_score from users where id=?", userID).Scan(&user.CreditScore).Error
}
