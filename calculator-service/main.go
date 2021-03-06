package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/go-calculator/calculator-service/proto/calculator"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)


type Configuration struct {
	Port string
}


func main() {
	configuration := getConfiguration()

	lis, err := net.Listen("tcp", configuration.Port)
	if err != nil {
		log.Fatal(err)
	}

	//creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//opts := []grpc.ServerOption{grpc.Creds(creds)}
	//s := grpc.NewServer(opts...)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, new(calculatorServiceServer))

	reflection.Register(s)
	log.Println("Starting calculator server on port " + configuration.Port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type calculatorServiceServer struct {
}

func (s *calculatorServiceServer) AddNumbers(ctx context.Context, in *pb.OperationRequest) (*pb.CalculatorResponse, error) {

	sumOfValues := in.Num1 + in.Num2
	res := &pb.CalculatorResponse{}
	res.CalculatedValue = sumOfValues
	fmt.Println("AddNumbers operation called, output value: ", res.CalculatedValue)
	return res, nil
}

func (s *calculatorServiceServer) SubtractNumbers(ctx context.Context, in *pb.OperationRequest) (*pb.CalculatorResponse, error) {
	subOfValues := in.Num1 - in.Num2
	res := &pb.CalculatorResponse{CalculatedValue: subOfValues}
	fmt.Println("SubtractNumbers operation called, output value: ", res.CalculatedValue)
	return res, nil

}

func (s *calculatorServiceServer) GetAverage(ctx context.Context, in *pb.GetAverageRequest) (*pb.CalculatorResponse, error) {
	sum := float32(0)
	for _, i := range in.Numbers {
		sum += i
	}
	avg := float32(0)
	if len(in.Numbers) > 0 {
		avg = sum / float32(len(in.Numbers))
	}
	res := &pb.CalculatorResponse{CalculatedValue: avg}
	fmt.Println("GetAverage operation called, output value: ", res.CalculatedValue)
	return res, nil
}

func getConfiguration() Configuration {
	configuration := Configuration{}

	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
	}
	return configuration
}
