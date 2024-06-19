package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/richardktran/grpc-golang/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("Received Sum RPC")
	resp := &calculatorpb.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}

	return resp, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PNRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	log.Println("Received PrimeNumberDecomposition RPC")
	num := req.GetNumber()
	divisor := int32(2)

	for num > 1 {
		if num%divisor == 0 {
			res := &calculatorpb.PNResponse{
				Result: divisor,
			}

			stream.Send(res)
			num = num / divisor
		} else {
			divisor++
			log.Printf("Divisor has increased to %v\n", divisor)
		}
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
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
