package domain

// Address defines attributes of this value object
type Address struct {
	streetName   string
	streetNumber int
}

// NewAddress is the only way to create this object
func NewAddress(streetName string, streetNumber int) Address {
	return Address{
		streetName:   streetName,
		streetNumber: streetNumber,
	}
}
