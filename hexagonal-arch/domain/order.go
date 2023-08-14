package domain

// Order defines attributes of this aggregate root
type Order struct {
	orderId            string
	customer           Customer
	product            Product
	customerRepository CustomerRepository
	productRepository  ProductRepository
	orderRepository    OrderRepository
}

// NewOrder is the only way to create this object
func NewOrder(orderId string, customerRepository CustomerRepository, productRepository ProductRepository, orderRepository OrderRepository) Order {
	return Order{
		orderId:            orderId,
		customerRepository: customerRepository,
		productRepository:  productRepository,
		orderRepository:    orderRepository,
	}
}

// LookUpCustomer looks for customer information
func (o Order) LookUpCustomer(customerId string) {
	o.customer = o.customerRepository.GetCustomer(customerId)
}

// LookUpProduct looks for product information
func (o Order) LookUpProduct(productId string) {
	o.product = o.productRepository.GetProduct(productId)
}

// CreateOrder creates an order
func (o Order) CreateOrder() bool {
	if !o.customer.IsNumberOrderAllowed() {
		return false
	}

	return o.orderRepository.CreateOrder(o.orderId, o.customer)
}
