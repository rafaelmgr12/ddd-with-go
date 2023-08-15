package port

import (
	"github.com/rafaelmgr12/ddd-with-go/hexagonal-arch/domain"
	"github.com/rafaelmgr12/ddd-with-go/hexagonal-arch/usecase"
)

// createOrderPortImpl defines attributes of this object
type createOrderPortImpl struct {
	customerRepository domain.CustomerRepository
	productRepository  domain.ProductRepository
	orderRepository    domain.OrderRepository
}

func NewCreateOrderPortImpl(customerRepository domain.CustomerRepository, productRepository domain.ProductRepository, orderRepository domain.OrderRepository) CreateOrderPort {
	return createOrderPortImpl{
		customerRepository: customerRepository,
		productRepository:  productRepository,
		orderRepository:    orderRepository,
	}
}

// CreateOrderPort executes an order creation
func (c createOrderPortImpl) CreateOrderPort(dto InputDTO) OutputDTO {
	createOrderUseCase := usecase.NewCreateOrderImpl(c.customerRepository, c.productRepository, c.orderRepository)
	result := createOrderUseCase.CreateOrder(dto.OrderId, dto.CustomerId, dto.ProductId)
	res := OutputDTO{}
	if result {
		res.Result = "Creation successful"
	} else {
		res.Result = "Creation failed"
	}

	return res
}
