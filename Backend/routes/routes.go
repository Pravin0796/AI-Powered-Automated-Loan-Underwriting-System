package routes

import (
	"AI-Powered-Automated-Loan-Underwriting-System/config"
	pb "AI-Powered-Automated-Loan-Underwriting-System/created_proto/pb" // Update with your actual module path
	"AI-Powered-Automated-Loan-Underwriting-System/services"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// StartGRPCServer initializes and starts the gRPC server
func StartGRPCServer() {
	listener, err := net.Listen("tcp", ":50051") // Use the correct port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer) // Register reflection service on gRPC server used for debugging using eg. grpcurl

	// Register gRPC services
	pb.RegisterUserServiceServer(grpcServer, &services.UserService{DB: config.DB})
	//pb.RegisterUserServiceServer(grpcServer, &services.UserServer{})
	//pb.RegisterPaymentServiceServer(grpcServer, &services.PaymentServer{})

	log.Println("gRPC Server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
