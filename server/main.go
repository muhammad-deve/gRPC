package main

import (
	"context"
	"log"
	"net"

	pb "calculator-server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// CalculatorServer implements the Calculator service
type CalculatorServer struct {
	pb.UnimplementedCalculatorServer
}

// Add implements the Add RPC method
func (s *CalculatorServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Received Add request: a=%d, b=%d", req.A, req.B)

	result := req.A + req.B

	log.Printf("Sending response: result=%d", result)
	return &pb.AddResponse{Result: result}, nil
}

func main() {
	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register our service
	pb.RegisterCalculatorServer(grpcServer, &CalculatorServer{})

	// Enable reflection (useful for testing with tools like grpcurl)
	reflection.Register(grpcServer)

	log.Println("🚀 gRPC Calculator Server starting on port 50051...")
	log.Println("📡 Server is ready to accept connections")

	// Start serving
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
