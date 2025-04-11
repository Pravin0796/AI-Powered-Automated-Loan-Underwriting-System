package services

import (
	"AI-Powered-Automated-Loan-Underwriting-System/middleware"
	"AI-Powered-Automated-Loan-Underwriting-System/repositories"
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	pb "AI-Powered-Automated-Loan-Underwriting-System/created_proto/user"
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"AI-Powered-Automated-Loan-Underwriting-System/utils"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	//DB   *gorm.DB
	repo *repositories.UserRepo
}

func NewUserService(repo *repositories.UserRepo) *UserService {
	return &UserService{repo: repo}
}

// Register user
func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := models.User{
		FullName:    req.FullName,
		Email:       req.Email,
		Password:    string(hashedPassword),
		Phone:       req.Phone,
		DateOfBirth: time.Now(), // Convert string to time if needed
		Address:     req.Address,
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, errors.New("failed to create user")
	}

	return &pb.RegisterResponse{
		Message: "User registered successfully",
		Status:  200,
	}, nil
}

// Login user
func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User
	if err := s.repo.GetUserByEmail(ctx, req, &user); err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		fmt.Println("Password mismatch:", err)
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &pb.LoginResponse{
		Token:  token,
		Status: 200,
	}, nil
}

// GetCurrentUser retrieves the current user based on the token
func (s *UserService) GetUserDetails(ctx context.Context, req *pb.UserDetailsRequest) (*pb.UserDetailsResponse, error) {
	var userID any = ctx.Value(middleware.ContextUserIDKey)

	var user models.User
	if err := s.repo.GetUserByID(ctx, userID.(uint), &user); err != nil {
		return nil, errors.New("user not found")
	}

	return &pb.UserDetailsResponse{
		FullName:    user.FullName,
		Email:       user.Email,
		Phone:       user.Phone,
		Address:     user.Address,
		DateOfBirth: user.DateOfBirth.Format("02/01/2006"),
		CreditScore: int32(user.CreditScore),
		Status:      200,
	}, nil
}

// GetCreditScore retrieves the credit score of a user
func (s *UserService) GetUserCreditScore(ctx context.Context, req *pb.UserCreditScoreRequest) (*pb.UserCreditScoreResponse, error) {
	var user models.User
	if err := s.repo.GetUserByID(ctx, uint(req.UserId), &user); err != nil {
		return nil, errors.New("user not found")
	}

	if err := s.repo.GetCreditScore(ctx, uint(req.UserId), &user); err != nil {
		return nil, errors.New("failed to retrieve credit score")
	}
	fmt.Println(user)

	return &pb.UserCreditScoreResponse{
		CreditScore: int32(user.CreditScore),
		Status:      200,
	}, nil
}
