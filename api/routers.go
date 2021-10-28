package api

import (
	"ApiModule/api/middleware"
	"ApiModule/api/v1/product"
	"ApiModule/api/v1/user"

	echo "github.com/labstack/echo/v4"
)

func AddRoute(e *echo.Echo, user *user.Controller, product *product.Controller) {
	if user == nil || product == nil {
		panic("Invalid route parameters")
	}

	e.POST("/v1/register", user.AddNewUser)
	e.POST("/v1/login", user.LoginUser)

	eproduct := e.Group("/v1/products")
	eproduct.Use(middleware.JWTMiddleware())
	eproduct.GET("", product.GetAllProduct)
	eproduct.POST("", product.AddNewProduct)
	eproduct.GET("/:id", product.FindProductBy)
	eproduct.PUT("/:id", product.ModifyProduct)
	eproduct.GET("/:id/photo", product.GetProductImageById)
}
