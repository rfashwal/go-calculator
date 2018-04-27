package main

import (
	"fmt"
	"log"

	pb "go-calculator/calculator-service/proto/calculator"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9001"
)

func main() {

	client := connectToServer()

	numbers := []float32{1, 5, 6}
	addNumbers(1, 5, client)
	subtractNumbers(6, 8, client)
	calculateAverage(numbers, client)
}

func addNumbers(num1 float32, num2 float32, client pb.CalculatorServiceClient) (result float32) {
	c := connectToServer()
	req := &pb.OperationRequest{Num1: num1, Num2: num2}
	r, err := c.AddNumbers(context.Background(), req)
	if err != nil {
		fmt.Printf("Error adding numbers %v", err)
		return 0
	}

	fmt.Printf("Added numbers %v, %v. Result: %v", num1, num2, r.CalculatedValue)
	fmt.Println()
	return r.CalculatedValue
}

func subtractNumbers(num1 float32, num2 float32, client pb.CalculatorServiceClient) (result float32) {
	req := &pb.OperationRequest{Num1: num1, Num2: num2}
	r, err := client.SubtractNumbers(context.Background(), req)
	if err != nil {
		fmt.Printf("Error subtracting numbers %v", err)
		return 0
	}

	fmt.Printf("Subtracted numbers %v, %v. Result: %v", num1, num2, r.CalculatedValue)
	fmt.Println()
	return r.CalculatedValue
}

func calculateAverage(numbers []float32, client pb.CalculatorServiceClient) (result float32) {
	req := &pb.GetAverageRequest{Numbers: numbers}
	r, err := client.GetAverage(context.Background(), req)
	if err != nil {
		fmt.Printf("Error calculating average %v", err)
		return 0
	}

	fmt.Printf("average number calculated for %v. Result: %v", numbers, r.CalculatedValue)
	fmt.Println()
	return r.CalculatedValue
}

func connectToServer() pb.CalculatorServiceClient {
	//creds, err := credentials.NewClientTLSFromFile("cert.pem", "")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	// Set up a connection to the server.
	//conn, err := grpc.Dial(address, opts...)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	client := pb.NewCalculatorServiceClient(conn)

	return client
}
