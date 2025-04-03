package routes

import (
	"AI-Powered-Automated-Loan-Underwriting-System/config"
	pb "AI-Powered-Automated-Loan-Underwriting-System/created_proto/pb" // Update with your actual module path
	"AI-Powered-Automated-Loan-Underwriting-System/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

// StartGRPCServer initializes and starts the gRPC server
func StartGRPCServer() {
	listener, err := net.Listen("tcp", ":50051") // Use the correct port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register gRPC services
	pb.RegisterUserServiceServer(grpcServer, &services.UserService{DB: config.DB})
	//pb.RegisterUserServiceServer(grpcServer, &services.UserServer{})
	//pb.RegisterPaymentServiceServer(grpcServer, &services.PaymentServer{})

	log.Println("gRPC Server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
