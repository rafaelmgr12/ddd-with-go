package main

// Customer is an aggregate, it has attributes of a customer
type Customer struct {
	customerId   string
	name         *PersonName
	emailAddress string
}

// NewCustomer is the only way to create a customer object
func NewCustomer(customerId, firstName, lastName, emailAddress string) *Customer {
	return &Customer{
		customerId:   customerId,
		name:         NewPersonName(firstName, lastName),
		emailAddress: emailAddress,
	}
}
