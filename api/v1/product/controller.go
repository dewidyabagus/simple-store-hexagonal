package product

import (
	"ApiModule/api/common"
	"ApiModule/api/v1/product/request"
	"ApiModule/api/v1/product/response"
	"ApiModule/business/product"
	"mime/multipart"

	"strconv"

	"github.com/google/uuid"
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

func (c *Controller) GetProductImageById(ctx echo.Context) error {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	source, err := c.service.GetProductImageById(id)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.File(source)
}

func (c *Controller) FindProductBy(ctx echo.Context) error {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	product, err := c.service.FindProductById(id)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.GetProduct(*product)))
}

func (c *Controller) FindProductByParams(ctx echo.Context) error {
	code := ctx.QueryParam("code")
	name := ctx.QueryParam("name")
	active := ctx.QueryParam("active")

	if len(code)+len(name)+len(active) == 0 {
		return ctx.JSON(common.BadRequestResponse())
	}

	products, err := c.service.FindProductByParams(code, name, active)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.AllProduct(products)))
}

func (c *Controller) AddNewProduct(ctx echo.Context) error {
	qty, _ := strconv.Atoi(ctx.FormValue("qty"))
	price, _ := strconv.Atoi(ctx.FormValue("price"))

	product := request.Product{
		Code:        ctx.FormValue("code"),
		Name:        ctx.FormValue("name"),
		Price:       price,
		Qty:         qty,
		Description: ctx.FormValue("description"),
	}

	photo, err := ctx.FormFile("photo")
	if err != nil {
		if err.Error() != "http: no such file" {
			return ctx.JSON(common.BadRequestResponse())
		}
		product.Photo = &multipart.FileHeader{}

	} else {
		product.Photo = photo

	}

	err = c.service.AddNewProduct(product.ToBusinessProduct())
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) ModifyProduct(ctx echo.Context) error {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	qty, _ := strconv.Atoi(ctx.FormValue("qty"))
	price, _ := strconv.Atoi(ctx.FormValue("price"))

	product := request.Product{
		Code:        ctx.FormValue("code"),
		Name:        ctx.FormValue("name"),
		Price:       price,
		Qty:         qty,
		Description: ctx.FormValue("description"),
	}

	photo, err := ctx.FormFile("photo")
	if err != nil {
		if err.Error() != "http: no such file" {
			return ctx.JSON(common.BadRequestResponse())
		}
		product.Photo = &multipart.FileHeader{}

	} else {
		product.Photo = photo

	}

	err = c.service.ModifyProduct(id, product.ToBusinessProduct())
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) DeleteProduct(ctx echo.Context) error {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	err := c.service.DeleteProduct(id)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}
