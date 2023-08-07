package main

import "fmt"

// Sells is an aggregate root, it has the attributes of Sells
type Sells struct {
	product            *Product
	customer           *Customer
	productRepository  ProductRepository
	customerRepository CustomerRepository
}

// NewSells is the only way to create a sells object
func NewSells(productId string, customerId string, productRepository ProductRepository, customerRepository CustomerRepository) Sells {
	sells := Sells{}
	sells.productRepository = productRepository
	sells.customerRepository = customerRepository
	sells.createBillInformation(productId, customerId)
	return sells
}

// GenerateBill generates a bill based on the information generated previously
func (c *Sells) GenerateBill() {
	fmt.Printf("Generating bill for customer %v whose name is %v. The product is %v and its price is %v", c.customer.customerId, c.customer.name.firstName, c.product.name, c.product.price.value)
}

// createBillInformation is a custom implementation of an aggregate
func (c *Sells) createBillInformation(productId, customerId string) {
	c.customer = c.customerRepository.GetCustomerInformation(customerId)
	c.product = c.productRepository.GetProductInformation(productId)
}
