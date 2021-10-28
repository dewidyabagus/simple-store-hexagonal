package product

type Service interface {
	// Get all product
	GetAllProduct() (*[]Product, error)

	// Get product image by id
	GetProductImageById(id string) (string, error)

	// Add new product
	AddNewProduct(product *ProductSpec) error

	// Modify Information Product
	ModifyProduct(id string, product *ProductSpec) error
}

type Repository interface {
	// Get all product
	GetAllProduct() (*[]Product, error)

	// Get product image by id
	GetProductImageById(id string) (string, error)

	// Add new product
	AddNewProduct(product *Product) error

	// Modify Information Product
	ModifyProduct(id string, product *ModifyProduct) error
}
