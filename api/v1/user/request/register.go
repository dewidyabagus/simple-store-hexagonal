package request

import (
	"ApiModule/business/user"
)

type RegisterUser struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
}

func (r *RegisterUser) ToNewUser() *user.User {
	return &user.User{
		Email:     r.Email,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Password:  r.Password,
	}
}
