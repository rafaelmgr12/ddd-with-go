package domain

// Customer defines attributes of this entity
type Customer struct {
	customerId       string
	customerName     string
	ordersInProgress int
	address          Address
}

// NewCustomer is the only way to create this object
func NewCustomer(customerId string, customerName string, ordersInProgress int, streetName string, streetNumber int) Customer {
	return Customer{
		customerId:       customerId,
		customerName:     customerName,
		ordersInProgress: ordersInProgress,
		address:          NewAddress(streetName, streetNumber),
	}
}

// IsNumberOrderAllowed verifies whether customer has opened orders
func (c Customer) IsNumberOrderAllowed() bool {
	return c.ordersInProgress == 0
}

// CustomerId returns customerId value
func (c Customer) CustomerId() string {
	return c.customerId
}
