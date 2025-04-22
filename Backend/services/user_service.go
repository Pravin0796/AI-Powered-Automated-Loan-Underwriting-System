package services

import (
	"AI-Powered-Automated-Loan-Underwriting-System/repositories"
	"context"
	"errors"
	"fmt"
	_ "time"

	"google.golang.org/protobuf/types/known/timestamppb"

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
		DateOfBirth: req.DateOfBirth.AsTime(), // Convert string to time if needed
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
	// var userID any = ctx.Value(middleware.ContextUserIDKey)
	// fmt.Println(userID)

	var userID any = req.UserId

	var user models.User
	if err := s.repo.GetUserByID(ctx, userID.(uint), &user); err != nil {
		return nil, errors.New("user not found")
	}

	var dob *timestamppb.Timestamp
	if !user.DateOfBirth.IsZero() {
		dob = timestamppb.New(user.DateOfBirth)
	} else {
		dob = nil
	}
	fmt.Println(user.DateOfBirth)

	return &pb.UserDetailsResponse{
		FullName:    user.FullName,
		Email:       user.Email,
		Phone:       user.Phone,
		Address:     user.Address,
		DateOfBirth: dob,
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

// UpdateUserDetails updates the user info
func (s *UserService) UpdateUserDetails(ctx context.Context, req *pb.UserUpdateRequest) (*pb.UserUpdateResponse, error) {
	var user models.User
	if err := s.repo.GetUserByID(ctx, uint(req.UserId), &user); err != nil {
		return nil, errors.New("user not found")
	}

	user.FullName = req.FullName
	user.Phone = req.Phone
	user.DateOfBirth = req.DateOfBirth.AsTime()
	user.Address = req.Address

	if err := s.repo.UpdateUser(ctx, &user); err != nil {
		return nil, errors.New("failed to update user details")
	}

	return &pb.UserUpdateResponse{
		Message: "User details updated successfully",
		Status:  200,
	}, nil
}
