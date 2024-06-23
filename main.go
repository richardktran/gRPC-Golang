package main

import (
	"fmt"
	"log"

	"github.com/richardktran/grpc-golang/protogen/golang/orders"
	"github.com/richardktran/grpc-golang/protogen/golang/product"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	orderItem := orders.Order{
		OrderId:    1,
		CustomerId: 1,
		IsActive:   true,
		OrderDate:  &date.Date{Year: 2021, Month: 1, Day: 1},
		Products: []*product.Product{
			{ProductId: 1, ProductName: "Product 1", ProductType: product.ProductType_DRINK},
		},
	}

	// Marshal the orderItem to JSON
	bytes, err := protojson.Marshal(&orderItem)
	if err != nil {
		log.Fatal("deserialize error: ", err)
	}

	fmt.Println(string(bytes))
}
