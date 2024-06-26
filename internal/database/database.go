package database

import (
	"fmt"
	"log"

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

	log.Printf("Adding order with id %d\n", order.GetOrderId())

	db.collection = append(db.collection, order)

	return nil
}

func (db *DB) GetOrderByID(orderId uint64) *orders.Order {
	for _, o := range db.collection {
		if o.OrderId == orderId {
			return o
		}
	}

	return nil
}

func (db *DB) UpdateOrder(order *orders.Order) {
	log.Printf("Updating order with id %d\n", order.GetOrderId())
	for i, o := range db.collection {
		if o.OrderId == order.OrderId {
			db.collection[i] = order
		}
	}
}

func (db *DB) DeleteOrder(orderId uint64) {
	log.Printf("Deleting order with id %d\n", orderId)
	newCollection := make([]*orders.Order, 0)
	for i, o := range db.collection {
		if o.OrderId != orderId {
			newCollection = append(newCollection, db.collection[i])
		}
	}

	db.collection = newCollection
}
