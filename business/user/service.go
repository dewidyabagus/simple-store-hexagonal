package user

import (
	"ApiModule/business"
	"ApiModule/config"
	"ApiModule/util/validator"
	"time"

	"github.com/golang-jwt/jwt"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) AddNewUser(user *User) error {
	err := validator.GetValidator().Struct(user)
	if err != nil {
		return business.ErrInvalidSpec
	}
	return s.repository.AddNewUser(user.toUser())
}

func (s *service) GetLoginUser(user *LoginUser) (string, error) {
	err := validator.GetValidator().Struct(user)
	if err != nil {
		return "", business.ErrInvalidSpec
	}

	uuid, err := s.repository.GetLoginUser(user.Email, user.Password)
	if err != nil {
		return "", business.ErrNotFound
	}

	token, err := s.createToken(uuid)
	if err != nil {
		return "", nil
	}

	return token, nil

}

func (s *service) createToken(uuid string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized":  true,
		"access_uuid": uuid,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := at.SignedString([]byte(config.GetConfig().JwtSecretKey))

	if err != nil {
		return "", err
	}

	return token, nil
}
