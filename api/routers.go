package api

import (
	"ApiModule/api/v1/user"

	echo "github.com/labstack/echo/v4"
)

func AddRoute(e *echo.Echo, user *user.Controller) {
	if user == nil {
		panic("Invalid route parameters")
	}

	e.POST("/v1/register", user.AddNewUser)
	e.POST("/v1/login", user.LoginUser)
}
