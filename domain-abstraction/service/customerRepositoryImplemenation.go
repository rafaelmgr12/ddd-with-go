package main

// CustomerRepositoryImplementation is the implementation of the repository
type CustomerRepositoryImplementation struct {
}

// NewCustomerRepositoryImplementation is the only way to create this object
func NewCustomerRepositoryImplementation() CustomerRepository {
	return CustomerRepositoryImplementation{}
}

// GetCustomerInformation allows to search for a particular customer
func (c CustomerRepositoryImplementation) GetCustomerInformation(customerId string) *Customer {
	return NewCustomer(customerId, "DDD", "Test", "dddTest@educative.com")
}
