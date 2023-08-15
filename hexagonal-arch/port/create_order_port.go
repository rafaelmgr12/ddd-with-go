package port

type CreateOrderPort interface {

	// CreateOrderPort exposes interface to create an order
	CreateOrderPort(dto InputDTO) OutputDTO
}
