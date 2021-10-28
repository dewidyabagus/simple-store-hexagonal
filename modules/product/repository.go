package product

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"ApiModule/business"
	"ApiModule/business/product"

	"gorm.io/gorm"
)

const (
	ImageLocation string = "./modules/product/images/"
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
	DeletedAt   time.Time `gorm:"deleted_at;default:'0001-01-01 00:00:00+00'"`
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

type ProductImage struct {
	Photo string `gorm:"photo"`
}

func insertProduct(p *product.Product) *Product {
	return &Product{
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

func createImage(flname string, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(ImageLocation + flname)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func removeImage(flname string) error {
	return os.Remove(ImageLocation + flname)
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

	err := r.DB.Where("deleted_at IS NULL").Find(products).Error
	if err != nil {
		return nil, err
	}

	return allBusinessProduct(products), nil
}

func (r *Repository) GetProductImageById(id string) (string, error) {
	var source ProductImage

	err := r.DB.Model(&Product{}).Where("ID = ? AND deleted_at IS NULL", id).Find(&source).Error
	if err != nil {
		return "", err
	}

	if len(source.Photo) == 0 {
		return "", business.ErrNotFound
	}

	_, err = os.Stat(ImageLocation + source.Photo)
	if os.IsNotExist(err) {
		return "", business.ErrNotFound
	}

	return ImageLocation + source.Photo, nil
}

func (r *Repository) FindProductById(id string) (*product.Product, error) {
	item := new(Product)

	err := r.DB.First(item, "ID = ? and deleted_at IS NULL", id).Error
	if err != nil {
		return nil, err
	}

	return item.toBusinessProduct(), nil
}

func (r *Repository) FindProductByParams(code, name, active string) (*[]product.Product, error) {
	var products []Product

	var query = "select * from products where "
	var strWhere = ""

	if code != "" {
		if strWhere != "" {
			strWhere += " and "
		}
		strWhere += fmt.Sprintf(" code = '%s'", code)
	}

	if name != "" {
		if strWhere != "" {
			strWhere += " and "
		}
		strWhere += fmt.Sprintf(" name = '%s'", name)
	}

	if active != "" {
		if strWhere != "" {
			strWhere += " and "
		}
		res, err := strconv.ParseBool(active)
		if err != nil {
			return nil, err
		}

		if res {
			strWhere += " deleted_at IS NULL"
		} else {
			strWhere += " deleted_at IS NOT NULL"
		}

	}

	err := r.DB.Raw(query + strWhere).Scan(&products).Error
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, business.ErrNotFound
	}

	return allBusinessProduct(&products), nil
}

func (r *Repository) AddNewProduct(product *product.Product) error {
	var err error

	err = r.DB.First(&Product{}, " code = ?", product.Code).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return business.ErrConflig
	}

	item := insertProduct(product)

	err = r.DB.Create(item).Error
	if err != nil {
		return err
	}

	err = createImage(product.Photo, product.File)
	if err != nil {
		r.DB.Delete(item).Where("ID", product.ID)
		return err
	}

	return nil
}

func (r *Repository) ModifyProduct(id string, product *product.ModifyProduct) error {
	var err error
	var item Product

	err = r.DB.First(&item, "ID = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return business.ErrNotFound
		}
		return err
	}

	var tmp Product
	err = r.DB.First(&tmp, " code = ? ", product.Code).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err == nil {
			if tmp.ID != id {
				return business.ErrConflig
			}
		}
	}

	_ = removeImage(item.Photo)
	err = createImage(product.Photo, product.File)
	if err != nil {
		return err
	}

	return r.DB.Model(&item).Updates(Product{
		Code:        product.Code,
		Name:        product.Name,
		Price:       product.Price,
		Qty:         product.Qty,
		Description: product.Description,
		Photo:       product.Photo,
		UpdatedAt:   product.UpdatedAt,
	}).Error
}

func (r *Repository) DeleteProduct(id string) error {
	product := new(Product)

	err := r.DB.First(product, "ID = ? and deleted_at IS NULL", id).Error
	if err != nil {
		return err
	}

	return r.DB.Model(product).Updates(Product{DeletedAt: time.Now()}).Error
}
