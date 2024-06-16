package main

import (
	"fmt"
	"log"
	"net"

	"github.com/richardktran/grpc-golang/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	fmt.Println("Calculator is running...")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
