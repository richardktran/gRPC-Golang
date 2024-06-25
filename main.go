package main

import (
	"log"

	"github.com/richardktran/grpc-golang/internal/configs"
)

func main() {
	config, err := configs.NewConfig("configs/local.yaml")
	if err != nil {
		log.Fatalf("Can not load config")
	}

	grpcServer := config.GRPC

	if err := grpcServer.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
