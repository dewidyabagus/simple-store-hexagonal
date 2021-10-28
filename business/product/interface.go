package product

type Service interface {
	GetAllProduct() (*[]Product, error)
	AddNewProduct(product *ProductSpec) error
}

type Repository interface {
	GetAllProduct() (*[]Product, error)
	AddNewProduct(product *Product) error
}
