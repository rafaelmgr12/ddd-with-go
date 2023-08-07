package main

// ProductRepository defines all functionalities in a repository
type ProductRepository interface {

	// GetProductInformation allows to search for a particular product
	GetProductInformation(productId string) *Product
}
