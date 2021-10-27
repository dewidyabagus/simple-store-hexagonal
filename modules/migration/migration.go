package migration

import (
	"ApiModule/modules/user"

	"gorm.io/gorm"
)

func TableMigration(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
}
