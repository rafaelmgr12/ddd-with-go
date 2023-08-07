package main

// Price contains attributes of price value object
type Price struct {
	value    float64
	currency string
}

// NewPrice is the only way to create a price object
func NewPrice(value float64, currency string) *Price {
	return &Price{
		value:    value,
		currency: currency,
	}
}
