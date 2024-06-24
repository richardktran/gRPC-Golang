package db

import (
	"fmt"

	"github.com/richardktran/grpc-golang/protogen/golang/orders"
)

type DB struct {
	collection []*orders.Order
}

func NewDB() *DB {
	return &DB{
		collection: make([]*orders.Order, 0),
	}
}

func (db *DB) AddOrder(order *orders.Order) error {
	for _, o := range db.collection {
		if o.OrderId == order.OrderId {
			return fmt.Errorf("order with id %d already exists", order.OrderId)
		}
	}

	db.collection = append(db.collection, order)

	return nil
}
