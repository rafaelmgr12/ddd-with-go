package domain

type ProductRepository interface {

	// GetProduct looks for information of a particular product
	GetProduct(productId string) Product
}
