package usecase

type CreateOrder interface {

	// CreateOrder defines function required to create an order
	CreateOrder(orderId string, customerId string, productId string) bool
}
