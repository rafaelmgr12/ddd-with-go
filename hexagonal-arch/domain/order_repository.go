package domain

type OrderRepository interface {

	// CreateOrder saves an order
	CreateOrder(orderId string, customer Customer) bool
}
