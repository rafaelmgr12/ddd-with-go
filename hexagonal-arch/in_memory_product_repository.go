package main

import "github.com/rafaelmgr12/ddd-with-go/hexagonal-arch/domain"

type inMemoryProductRepository struct{}

func NewInMemoryProductRepository() domain.ProductRepository {
	return &inMemoryProductRepository{}
}

// GetProduct implements interface CustomerRepository, it mocks a store system
func (i *inMemoryProductRepository) GetProduct(productId string) domain.Product {
	return domain.NewProduct(productId, "pizza")
}
