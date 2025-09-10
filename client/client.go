package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "calculator-client/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the gRPC server
	log.Println("🔗 Connecting to Calculator gRPC Server...")

	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(), // Wait for connection to be established
	)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	log.Println("✅ Connected to server successfully!")

	// Create client
	client := pb.NewCalculatorClient(conn)

	// Prepare context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Test multiple calculations
	testCases := []struct {
		a, b int32
	}{
		{5, 3},
		{10, 20},
		{100, 200},
		{-5, 15},
	}

	for i, tc := range testCases {
		log.Printf("🧮 Test %d: Calculating %d + %d", i+1, tc.a, tc.b)

		// Make the gRPC call
		response, err := client.Add(ctx, &pb.AddRequest{
			A: tc.a,
			B: tc.b,
		})

		if err != nil {
			log.Fatalf("Add RPC failed: %v", err)
		}

		fmt.Printf("📊 Result: %d + %d = %d\n", tc.a, tc.b, response.Result)
		time.Sleep(500 * time.Millisecond) // Small delay between calls
	}

	log.Println("🎉 All calculations completed successfully!")
}
