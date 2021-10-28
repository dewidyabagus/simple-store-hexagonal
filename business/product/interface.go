package product

type Service interface {
	// Get all product
	GetAllProduct() (*[]Product, error)

	// Get product image by id
	GetProductImageById(id string) (string, error)

	// Get product detail by id
	FindProductById(id string) (*Product, error)

	// Add new product
	AddNewProduct(product *ProductSpec) error

	// Modify Information Product
	ModifyProduct(id string, product *ProductSpec) error

	// Delete product
	DeleteProduct(id string) error
}

type Repository interface {
	// Get all product
	GetAllProduct() (*[]Product, error)

	// Get product image by id
	GetProductImageById(id string) (string, error)

	// Get product detail by id
	FindProductById(id string) (*Product, error)

	// Add new product
	AddNewProduct(product *Product) error

	// Modify Information Product
	ModifyProduct(id string, product *ModifyProduct) error

	// Delete product
	DeleteProduct(id string) error
}
