package product

import (
	"time"

	"ApiModule/business/product"

	"gorm.io/gorm"
)

type Product struct {
	ID          string    `gorm:"id;type:uuid;primaryKey"`
	Code        string    `gorm:"code;type:varchar(12);index:product_code_uniq;unique"`
	Name        string    `gorm:"name;type:varchar(100)"`
	Price       int       `gorm:"price;default:0"`
	Qty         int       `gorm:"qty;default:0"`
	Description string    `gorm:"description;type:varchar(150)"`
	Photo       string    `gorm:"photo;type:varchar(100)"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
	DeletedAt   time.Time `gorm:"deleted_at"`
}

func (p *Product) toBusinessProduct() *product.Product {
	return &product.Product{
		ID:          p.ID,
		Code:        p.Code,
		Name:        p.Name,
		Price:       p.Price,
		Qty:         p.Qty,
		Description: p.Description,
		Photo:       p.Photo,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func allBusinessProduct(products *[]Product) *[]product.Product {
	var items []product.Product

	for _, item := range *products {
		items = append(items, *item.toBusinessProduct())
	}

	if items == nil {
		items = []product.Product{}
	}

	return &items
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAllProduct() (*[]product.Product, error) {
	products := new([]Product)

	err := r.DB.Find(products).Error
	if err != nil {
		return nil, err
	}

	return allBusinessProduct(products), nil
}

func (r *Repository) AddNewProduct(product *product.Product) error {
	var err error

	return err
}
