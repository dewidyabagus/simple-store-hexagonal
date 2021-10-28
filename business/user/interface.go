package user

type Service interface {
	AddNewUser(user *User) error
	GetLoginUser(user *LoginUser) (string, error)
}

type Repository interface {
	AddNewUser(user *User) error
	GetLoginUser(email string, password string) (string, error)
}
