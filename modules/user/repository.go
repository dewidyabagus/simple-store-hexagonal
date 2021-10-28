package user

import (
	"ApiModule/business"
	"ApiModule/business/user"
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"type:uuid;primary_key"`
	Email     string    `gorm:"email;index:user_email_uniq,unique;type:varchar(50)"`
	FirstName string    `gorm:"firstname;type:varchar(50)"`
	LastName  string    `gorm:"lastname;type:varchar(50)"`
	Password  string    `gorm:"password;type:varchar(100)"`
	CreatedAt time.Time `gorm:"created_at;type:timestamp"`
	UpdatedAt time.Time `gorm:"updated_at;type:timestamp"`
	DeletedAt time.Time `gorm:"deleted_at;type:timestamp"`
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func insertNewUser(u *user.User) *User {
	return &User{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Password:  u.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
func (r *Repository) AddNewUser(user *user.User) error {
	err := r.DB.First(&User{}, "email = ?", user.Email).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return business.ErrConflig
	}

	return r.DB.Create(insertNewUser(user)).Error
}

func (r *Repository) GetLoginUser(email string, password string) (string, error) {
	var user = new(User)

	err := r.DB.First(user, " email = ? and password = md5(?) ", email, password).Error
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
