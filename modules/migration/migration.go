package migration

import (
	"ApiModule/modules/product"
	"ApiModule/modules/user"

	"gorm.io/gorm"
)

func TableMigration(db *gorm.DB) {
	db.AutoMigrate(&user.User{}, &product.Product{})
}
