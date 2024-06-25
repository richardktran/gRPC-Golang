package services

import (
	"context"
	"log"

	"github.com/richardktran/grpc-golang/internal/database"
	"github.com/richardktran/grpc-golang/protogen/golang/orders"
)

type OrderService struct {
	db *database.DB
	orders.UnimplementedOrderServiceServer
}

func NewOrderService(db *database.DB) OrderService {
	return OrderService{
		db: db,
	}
}

func (s *OrderService) CreateOrder(_ context.Context, req *orders.Order) (*orders.Empty, error) {
	log.Printf("Received order request")

	err := s.db.AddOrder(req)

	return nil, err
}
