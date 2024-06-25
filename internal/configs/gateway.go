package configs

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/richardktran/grpc-golang/protogen/golang/orders"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Gateway struct {
	Address       string `yaml:"address"`
	ServerAddress string `yaml:"server_address"`
}

func (g *Gateway) Run() error {
	grpcClientConn, err := grpc.NewClient(g.ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("could not connect to %s: %v", g.ServerAddress, err)
	}
	defer grpcClientConn.Close()

	mux := runtime.NewServeMux()

	if err = orders.RegisterOrderServiceHandler(context.Background(), mux, grpcClientConn); err != nil {
		log.Fatalf("could not register service handler: %v", err)
	}

	fmt.Println("api gateway is running on ", g.Address)

	return http.ListenAndServe(g.Address, mux)
}
