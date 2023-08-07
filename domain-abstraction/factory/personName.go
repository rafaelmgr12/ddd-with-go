package main

// PersonName is a value object, it has attributes of a name
type PersonName struct {
	firstName string
	lastName  string
}

// NewPersonName is the only way to create a person name object
func NewPersonName(firstName, lastName string) *PersonName {
	return &PersonName{
		firstName: firstName,
		lastName:  lastName,
	}
}
