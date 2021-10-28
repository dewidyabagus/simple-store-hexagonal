package product

import (
	"mime/multipart"
	"time"
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
}

type ProductSpec struct {
	Code        string                `validate:"required"`
	Name        string                `validate:"required"`
	Price       int                   `validate:"required"`
	Qty         int                   `validate:"required"`
	Description string                `validate:""`
	Photo       *multipart.FileHeader `validate:""`
}
