package user

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string    `validate:""`
	Email     string    `validate:"required,email"`
	FirstName string    `validate:"required"`
	LastName  string    `validate:"required"`
	Password  string    `validate:"required"`
	CreatedAt time.Time `validate:""`
	UpdatedAt time.Time `validate:""`
}

func (s *User) toUser() *User {
	h := md5.New()
	h.Write([]byte(s.Password))
	s.Password = hex.EncodeToString(h.Sum(nil))
	s.ID = uuid.NewString()
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()

	return s
}

type LoginUser struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type UserToken struct {
	Token string
}
