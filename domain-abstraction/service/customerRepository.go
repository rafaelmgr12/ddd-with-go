package main

// CustomerRepository defines all functionalities in a repository
type CustomerRepository interface {

	// GetCustomerInformation allows to search for a particular customer
	GetCustomerInformation(customerId string) *Customer
}
