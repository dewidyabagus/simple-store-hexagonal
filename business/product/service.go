package product

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAllProduct() (*[]Product, error) {
	return s.repository.GetAllProduct()
}

func (s *service) AddNewProduct(product *ProductSpec) error {
	var err error

	return err
}
