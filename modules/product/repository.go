package product

import "time"

type Product struct {
	ID          string    `gorm:"id;type:uuid;primaryKey"`
	Code        string    `gorm:"code;type:varchar(50);index:product_code_uniq"`
	Name        string    `gorm:"name;type:varchar(100)"`
	Price       int64     `gorm:"price"`
	UnitName    string    `gorm:"unitname;type:varchar(50)"`
	Description string    `gorm:"description;type:varchar(255)"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
	DeletedAt   time.Time `gorm:"deleted_at"`
}
