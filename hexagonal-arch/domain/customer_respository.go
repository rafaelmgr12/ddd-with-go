package domain

type CustomerRepository interface {

	// GetCustomer looks for information of a particular customer
	GetCustomer(customerId string) Customer
}
