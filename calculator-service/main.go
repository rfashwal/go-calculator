package main

import (
	"fmt"
	"log"
	"net"

	pb "go-calculator/calculator-service/proto/calculator"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":9001"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, new(calculatorServiceServer))

	reflection.Register(s)
	log.Println("Starting calculator server on port " + port)

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
	fmt.Println(sumOfValues)
	return res, nil
}

func (s *calculatorServiceServer) SubtractNumbers(ctx context.Context, in *pb.OperationRequest) (*pb.CalculatorResponse, error) {
	subOfValues := in.Num1 - in.Num2
	res := &pb.CalculatorResponse{CalculatedValue: subOfValues}
	fmt.Println(subOfValues)
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
	return res, nil
}
