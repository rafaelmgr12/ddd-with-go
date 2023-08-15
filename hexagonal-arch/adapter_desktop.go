package main

import (
	"fmt"

	"github.com/rafaelmgr12/ddd-with-go/hexagonal-arch/domain"
	"github.com/rafaelmgr12/ddd-with-go/hexagonal-arch/port"
)

type AdapterDesktop struct {
	customerRepository domain.CustomerRepository
	orderRepository    domain.OrderRepository
	productRepository  domain.ProductRepository
}

func NewAdapterDesktop(customerRepository domain.CustomerRepository, orderRepository domain.OrderRepository, productRepository domain.ProductRepository) AdapterDesktop {
	return AdapterDesktop{
		customerRepository: customerRepository,
		orderRepository:    orderRepository,
		productRepository:  productRepository,
	}
}

func (ad *AdapterDesktop) CreateOrder(req Request) Response {
	inpDto := port.InputDTO{
		OrderId:    req.OrderId,
		CustomerId: req.CustomerId,
		ProductId:  req.ProductId,
	}
	adapter := port.NewCreateOrderPortImpl(ad.customerRepository, ad.productRepository, ad.orderRepository)
	res := Response{
		Result: adapter.CreateOrderPort(inpDto).Result,
	}

	return res
}

func main() {
	customerRepository := NewInMemoryCustomerRepository()
	orderRepository := NewInMemoryOrderRepository()
	productRepository := NewInMemoryProductRepository()
	req := Request{
		OrderId:    "O-1234",
		ProductId:  "P-24244",
		CustomerId: "1234",
	}

	fmt.Printf("Before creating an order \n")
	adapterDesktop := NewAdapterDesktop(customerRepository, orderRepository, productRepository)
	res := adapterDesktop.CreateOrder(req)

	fmt.Printf("\n\nAfter creating an order the result is: %s", res.Result)

}
