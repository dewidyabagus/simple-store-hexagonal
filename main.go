package main

import (
	"ApiModule/api"
	productController "ApiModule/api/v1/product"
	userController "ApiModule/api/v1/user"

	productService "ApiModule/business/product"
	userService "ApiModule/business/user"

	"ApiModule/config"
	"ApiModule/modules/migration"

	productRepository "ApiModule/modules/product"
	userRepository "ApiModule/modules/user"

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	echo "github.com/labstack/echo/v4"
)

func newDatabaseConnection(config *config.ConfigApp) *gorm.DB {
	stringConfig := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		config.DbHost, config.DbPort, config.DbUsername, config.DbPassword, config.DbName,
	)
	db, err := gorm.Open(postgres.Open(stringConfig), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	migration.TableMigration(db)

	return db
}

func main() {
	// Geting config app
	config := config.GetConfig()

	// Initiate database connection
	dbConnection := newDatabaseConnection(config)

	// Initiate User Repository
	userRepo := userRepository.NewRepository(dbConnection)

	// Initiate User Service
	userServc := userService.NewService(userRepo)

	// Initiate User Controller
	userHandler := userController.NewController(userServc)

	// Initiate Product Repository
	productRepo := productRepository.NewRepository(dbConnection)

	// Initiate Product Service
	productServc := productService.NewService(productRepo)

	// Initiate Product Controller
	productHandler := productController.NewController(productServc)

	// Initiate Echo Web Service
	e := echo.New()

	// Add routing
	api.AddRoute(e, userHandler, productHandler)

	// Start service
	e.Start(fmt.Sprintf("%s:%d", config.AppHost, config.AppPort))
}
