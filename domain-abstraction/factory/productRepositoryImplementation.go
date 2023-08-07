package main

// ProductRepositoryImplementation is the implementation of the repository
type ProductRepositoryImplementation struct {
}

// NewProductRepositoryImplementation is the only way to create this object
func NewProductRepositoryImplementation() ProductRepository {
	return &ProductRepositoryImplementation{}
}

// GetProductInformation allows to search for a particular product
func (p *ProductRepositoryImplementation) GetProductInformation(productId string) *Product {
	return NewProduct(productId, 5644.8, "Product # "+productId)
}
