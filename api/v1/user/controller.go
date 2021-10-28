package user

import (
	"ApiModule/api/common"
	"ApiModule/api/v1/user/request"
	"ApiModule/business/user"

	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	service user.Service
}

func NewController(service user.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) AddNewUser(ctx echo.Context) error {
	user := &request.RegisterUser{
		Email:     ctx.FormValue("email"),
		FirstName: ctx.FormValue("firstname"),
		LastName:  ctx.FormValue("lastname"),
		Password:  ctx.FormValue("password"),
	}

	err := c.service.AddNewUser(user.ToNewUser())
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) LoginUser(ctx echo.Context) error {
	user := &request.LoginUser{
		Email:    ctx.FormValue("email"),
		Password: ctx.FormValue("password"),
	}

	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	token, err := c.service.GetLoginUser(user.ToLoginUser())
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(
		map[string]string{"token": token},
	))
}
