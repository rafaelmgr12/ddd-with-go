package usecase

import "github.com/rafaelmgr12/ddd-with-go/hexagonal-arch/domain"

// createOrderImpl defines attributes of this use case
type createOrderImpl struct {
	customerRepository domain.CustomerRepository
	productRepository  domain.ProductRepository
	orderRepository    domain.OrderRepository
}

// NewCreateOrderImpl implements interface CreateOrder
func NewCreateOrderImpl(customerRepository domain.CustomerRepository, productRepository domain.ProductRepository, orderRepository domain.OrderRepository) CreateOrder {
	return createOrderImpl{
		customerRepository: customerRepository,
		productRepository:  productRepository,
		orderRepository:    orderRepository,
	}
}

// CreateOrder executes logic required to create an order
func (c createOrderImpl) CreateOrder(orderId string, customerId string, productId string) bool {
	order := domain.NewOrder(orderId, c.customerRepository, c.productRepository, c.orderRepository)
	order.LookUpCustomer(customerId)
	order.LookUpProduct(productId)
	return order.CreateOrder()
}
