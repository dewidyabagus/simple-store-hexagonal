package response

import (
	"ApiModule/business/product"
	"time"
)

type DetailProduct struct {
	ID          string
	Code        string
	Name        string
	Price       int
	Qty         int
	Description string
	Photo       string
	UpdatedAt   time.Time
}

func GetProduct(p product.Product) DetailProduct {
	return DetailProduct{
		ID:          p.ID,
		Code:        p.Code,
		Name:        p.Name,
		Price:       p.Price,
		Qty:         p.Qty,
		Description: p.Description,
		Photo:       p.Photo,
		UpdatedAt:   p.UpdatedAt,
	}
}

func AllProduct(products *[]product.Product) *[]DetailProduct {
	var listProduct []DetailProduct

	for _, item := range *products {
		listProduct = append(listProduct, GetProduct(item))
	}

	if listProduct == nil {
		listProduct = []DetailProduct{}
	}

	return &listProduct
}
