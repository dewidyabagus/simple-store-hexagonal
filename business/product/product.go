package product

import (
	"mime/multipart"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          string
	Code        string
	Name        string
	Price       int
	Qty         int
	Description string
	Photo       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	File        *multipart.FileHeader
}

type ModifyProduct struct {
	Code        string
	Name        string
	Price       int
	Qty         int
	Description string
	Photo       string
	UpdatedAt   time.Time
	File        *multipart.FileHeader
}

type ProductSpec struct {
	Code        string                `validate:"required,len=12"`
	Name        string                `validate:"required"`
	Price       int                   `validate:"required"`
	Qty         int                   `validate:"required"`
	Description string                `validate:""`
	Photo       *multipart.FileHeader `validate:""`
}

func toProduct(p *ProductSpec) *Product {
	uuid := uuid.NewString()
	return &Product{
		ID:          uuid,
		Code:        p.Code,
		Name:        p.Name,
		Price:       p.Price,
		Qty:         p.Qty,
		Description: p.Description,
		Photo:       uuid + strings.ReplaceAll(p.Photo.Filename, " ", "-"),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		File:        p.Photo,
	}
}

func toModifyProduct(id string, p *ProductSpec) *ModifyProduct {
	return &ModifyProduct{
		Code:        p.Code,
		Name:        p.Name,
		Price:       p.Price,
		Qty:         p.Qty,
		Description: p.Description,
		Photo:       id + strings.ReplaceAll(p.Photo.Filename, " ", "-"),
		UpdatedAt:   time.Now(),
		File:        p.Photo,
	}
}
