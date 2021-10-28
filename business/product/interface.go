package product

type Service interface {
	// Get all product
	GetAllProduct() (*[]Product, error)

	// Get product image by id
	GetProductImageById(id string) (string, error)

	// Add new product
	AddNewProduct(product *ProductSpec) error
}

type Repository interface {
	// Get all product
	GetAllProduct() (*[]Product, error)

	// Get product image by id
	GetProductImageById(id string) (string, error)

	// Add new product
	AddNewProduct(product *Product) error
}
