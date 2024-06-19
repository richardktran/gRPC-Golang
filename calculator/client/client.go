package main

import (
	"context"
	"io"
	"log"

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

	callPND(client)
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
