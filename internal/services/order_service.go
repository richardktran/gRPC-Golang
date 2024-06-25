package services

import (
	"context"
	"fmt"
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

func (s *OrderService) CreateOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Printf("Received order request")

	err := s.db.AddOrder(req.GetOrder())

	return &orders.Empty{}, err
}

func (s *OrderService) GetOrder(_ context.Context, req *orders.PayloadWithOrderId) (*orders.PayloadWithSingleOrder, error) {
	log.Println("Received get order request")

	order := s.db.GetOrderByID(req.GetOrderId())

	if order == nil {
		return nil, fmt.Errorf("order with id %d not found", req.GetOrderId())
	}

	return &orders.PayloadWithSingleOrder{Order: order}, nil
}
