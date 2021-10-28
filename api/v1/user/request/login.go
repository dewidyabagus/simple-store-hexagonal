package request

import "ApiModule/business/user"

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginUser) ToLoginUser() *user.LoginUser {
	return &user.LoginUser{
		Email:    l.Email,
		Password: l.Password,
	}
}
