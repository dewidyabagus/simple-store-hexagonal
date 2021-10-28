package request

import (
	"ApiModule/business/product"
	"mime/multipart"
)

type Product struct {
	Code        string                `json:"code"`
	Name        string                `json:"name"`
	Price       int                   `json:"price"`
	Qty         int                   `json:"qty"`
	Description string                `json:"description"`
	Photo       *multipart.FileHeader `json:"photo"`
}

func (p *Product) ToBusinessProduct() *product.ProductSpec {
	return &product.ProductSpec{
		Code:        p.Code,
		Name:        p.Name,
		Price:       p.Price,
		Qty:         p.Qty,
		Description: p.Description,
		Photo:       p.Photo,
	}
}
