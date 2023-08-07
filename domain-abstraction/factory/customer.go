package main

// Customer is an aggregate, it has attributes of a customer
type Customer struct {
	customerId string
	name       *PersonName
}

// NewCustomer is the only way to create a customer object
func NewCustomer(customerId, firstName, lastName string) *Customer {
	return &Customer{
		customerId: customerId,
		name:       NewPersonName(firstName, lastName),
	}
}
