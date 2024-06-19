package main

import (
	"context"
	"fmt"
	"io"
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

func (*server) Average(stream calculatorpb.CalculatorService_AverageServer) error {
	total := int32(0)
	cnt := int32(0)
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			log.Println("Server finished streaming")
			resp := &calculatorpb.AverageResponse{
				Result: float32(total) / float32(cnt),
			}
			stream.SendAndClose(resp)
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		n := req.GetNumber()
		total += n
		cnt++

		log.Printf("Received number: %v\n", n)
	}
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
