package product

import (
	"ApiModule/business"
	"ApiModule/util/validator"
)

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
	err := validator.GetValidator().Struct(product)
	if err != nil {
		return business.ErrInvalidSpec
	}
	return s.repository.AddNewProduct(toProduct(product))
}

func (s *service) GetProductImageById(id string) (string, error) {
	return s.repository.GetProductImageById(id)
}
