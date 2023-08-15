package main

import "github.com/rafaelmgr12/ddd-with-go/hexagonal-arch/domain"

type inMemoryCustomerRepository struct{}

func NewInMemoryCustomerRepository() domain.CustomerRepository {
	return &inMemoryCustomerRepository{}
}

// GetCustomer implements interface CustomerRepository, it mocks a store system
func (i *inMemoryCustomerRepository) GetCustomer(customerId string) domain.Customer {
	return domain.NewCustomer(customerId, "DDD", 0, "educative", 1)
}
