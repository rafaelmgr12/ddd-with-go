package domain

// Product defines attributes of this entity
type Product struct {
	productId   string
	productName string
}

// NewProduct is the only way to create this object
func NewProduct(productId, productName string) Product {
	return Product{
		productId:   productId,
		productName: productName,
	}
}
