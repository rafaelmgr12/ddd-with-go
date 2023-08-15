package main

import (
	"fmt"

	"github.com/rafaelmgr12/ddd-with-go/hexagonal-arch/domain"
)

type inMemoryOrderRepository struct{}

func NewInMemoryOrderRepository() domain.OrderRepository {
	return &inMemoryOrderRepository{}
}

// CreateOrder implements interface OrderRepository, it mocks a store system
func (i *inMemoryOrderRepository) CreateOrder(orderId string, customer domain.Customer) bool {
	fmt.Printf("Creating order %s to customer %s", orderId, customer.CustomerId())
	return true
}
