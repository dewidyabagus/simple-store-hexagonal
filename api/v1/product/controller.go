package product

import (
	"ApiModule/api/common"
	"ApiModule/api/v1/product/request"
	"ApiModule/api/v1/product/response"
	"ApiModule/business/product"

	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	service product.Service
}

func NewController(service product.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) GetAllProduct(ctx echo.Context) error {
	products, err := c.service.GetAllProduct()

	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.AllProduct(products)))
}

// func (c *Controller) FindProductBy(ctx echo.Context) error {

// }

func (c *Controller) AddNewProduct(ctx echo.Context) error {
	product := request.Product{
		Code:        "",
		Name:        "",
		Price:       0,
		Qty:         0,
		Description: "",
	}

	photo, err := ctx.FormFile("photo")
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	product.Photo = photo

	err = c.service.AddNewProduct(product.ToBusinessProduct())
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

// func (c *Controller) ModifyProduct(ctx echo.Context) error {

// }

// func (c *Controller) DeleteProduct(ctx echo.Context) error {

// }
