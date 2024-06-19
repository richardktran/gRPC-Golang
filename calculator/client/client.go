package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/richardktran/grpc-golang/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	client := calculatorpb.NewCalculatorServiceClient(cc)

	callAverage(client)
}

func callSum(c calculatorpb.CalculatorServiceClient) {
	log.Println("Calling to do a Sum RPC...")
	resp, err := c.Sum(context.Background(), &calculatorpb.SumRequest{
		Num1: 3,
		Num2: 10,
	})

	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}

	log.Printf("Response from Sum: %v\n", resp.Result)
}

func callPND(client calculatorpb.CalculatorServiceClient) {
	stream, err := client.PrimeNumberDecomposition(context.Background(), &calculatorpb.PNRequest{
		Number: 120,
	})

	if err != nil {
		log.Fatalf("Error while calling PrimeNumberDecomposition RPC: %v", err)
	}

	for {
		resp, receiveErr := stream.Recv()

		if receiveErr == io.EOF {
			log.Println("server finished streaming")
			return
		}

		log.Printf("Prime number %v\n", resp.GetResult())
	}
}

func callAverage(client calculatorpb.CalculatorServiceClient) {
	log.Println("calling avg API")

	stream, err := client.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Average RPC: %v", err)
	}

	listReq := []calculatorpb.AverageRequest{
		{Number: 5},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	for _, req := range listReq {
		err := stream.Send(&req)

		if err != nil {
			log.Fatalf("Error while sending request to server: %v", err)
		}

		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from server: %v", err)
	}

	log.Printf("Average is %v\n", res.GetResult())
}
