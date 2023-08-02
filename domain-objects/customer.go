package main

// Customer contains attributes of a customer entity
type Customer struct {
	id        int
	firstName string
	lastName  string
}

// NewCustomer is the only way to create a customer entity
func NewCustomer(id int) *Customer {
	return &Customer{
		id: id,
	}
}
