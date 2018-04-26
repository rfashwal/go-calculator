// consignment-cli/cli.go
package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc/credentials"

	pb "go-calculator/calculator-service/proto/calculator"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9001"
)

func main() {

	creds, err := credentials.NewClientTLSFromFile("cert.pem", "")
	if err != nil {
		log.Fatal(err)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewCalculatorServiceClient(conn)

	numbers := []float32{1, 5, 6}
	addNumbers(1, 5, client)
	subtractNumbers(6, 8, client)
	calculateAverage(numbers, client)
}

func addNumbers(num1 float32, num2 float32, calculatorService pb.CalculatorServiceClient) {
	req := &pb.OperationRequest{Num1: num1, Num2: num2}
	r, err := calculatorService.AddNumbers(context.Background(), req)
	if err != nil {
		fmt.Printf("Error adding numbers %v", err)
		return
	}

	fmt.Printf("Added numbers %v, %v. Result: %v", num1, num2, r.CalculatedValue)
	fmt.Println()

}

func subtractNumbers(num1 float32, num2 float32, calculatorService pb.CalculatorServiceClient) {
	req := &pb.OperationRequest{Num1: num1, Num2: num2}
	r, err := calculatorService.SubtractNumbers(context.Background(), req)
	if err != nil {
		fmt.Printf("Error subtracting numbers %v", err)
		return
	}

	fmt.Printf("Subtracted numbers %v, %v. Result: %v", num1, num2, r.CalculatedValue)
	fmt.Println()
}

func calculateAverage(numbers []float32, calculatorService pb.CalculatorServiceClient) {
	req := &pb.GetAverageRequest{Numbers: numbers}
	r, err := calculatorService.GetAverage(context.Background(), req)
	if err != nil {
		fmt.Printf("Error calculating average %v", err)
		return
	}

	fmt.Printf("average number calculated for %v. Result: %v", numbers, r.CalculatedValue)
	fmt.Println()
}
