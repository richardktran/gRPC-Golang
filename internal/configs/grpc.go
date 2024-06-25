package configs

import (
	"log"
	"net"

	"github.com/richardktran/grpc-golang/internal/database"
	"github.com/richardktran/grpc-golang/internal/services"
	"github.com/richardktran/grpc-golang/protogen/golang/orders"
	"google.golang.org/grpc"
)

type GRPC struct {
	Address string `yaml:"address"`
}

func (s *GRPC) Run() error {
	listener, err := net.Listen("tcp", s.Address)

	if err != nil {
		log.Fatalf("failed to listen grpc protocol with address %s: %v", s.Address, err)
	}

	grpcServer := grpc.NewServer()
	db := database.NewDB()

	orderService := services.NewOrderService(db)

	orders.RegisterOrderServiceServer(grpcServer, &orderService)

	log.Printf("server listening at %v", listener.Addr())

	return grpcServer.Serve(listener)
}
