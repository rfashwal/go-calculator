// consignment-cli/cli.go
package main

import (
	"log"

	pb "go-calculator/calculator-service/proto/calculator"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9001"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewCalculatorServiceClient(conn)

	req := &pb.OperationRequest{Num1: 10, Num2: 5}

	r, err := client.AddNumbers(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Value: %t", r.CalculatedValue)

	r2, err2 := client.SubtractNumbers(context.Background(), req)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Value: %t", r2.CalculatedValue)

	numbers := []float32{1, 5, 6}
	avgReq := &pb.GetAverageRequest{Numbers: numbers}

	r3, err3 := client.GetAverage(context.Background(), avgReq)
	if err3 != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Value: %t", r3.CalculatedValue)
}
