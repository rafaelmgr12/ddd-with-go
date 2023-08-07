package main

// Product is an aggregate, it has attributes of a product
type Product struct {
	id    string
	price *Price
	name  string
}

// NewProduct is the only way to create a product object
func NewProduct(id string, price float64, name string) *Product {
	return &Product{
		id:    id,
		name:  name,
		price: NewPrice(price, "USD"),
	}
}
