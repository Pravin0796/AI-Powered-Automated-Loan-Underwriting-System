package routes

import (
	pbs "AI-Powered-Automated-Loan-Underwriting-System/created_proto/loan" // Update with your actual module path
	pb "AI-Powered-Automated-Loan-Underwriting-System/created_proto/user"
	"AI-Powered-Automated-Loan-Underwriting-System/middleware"
	"AI-Powered-Automated-Loan-Underwriting-System/repositories"
	"AI-Powered-Automated-Loan-Underwriting-System/services"
	"gorm.io/gorm"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// StartGRPCServer initializes and starts the gRPC server
func StartGRPCServer(db *gorm.DB) {
	listener, err := net.Listen("tcp", ":50051") // Use the correct port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.JWTAuthInterceptor))
	reflection.Register(grpcServer) // Register reflection service on gRPC server used for debugging using eg. grpcurl

	// Register gRPC services
	UserRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(UserRepo)
	pb.RegisterUserServiceServer(grpcServer, userService)

	pbs.RegisterLoanServiceServer(grpcServer, &services.LoanServiceServer{DB: db})
	//pb.RegisterUserServiceServer(grpcServer, &services.UserServer{})
	//pb.RegisterPaymentServiceServer(grpcServer, &services.PaymentServer{})

	log.Println("gRPC Server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
