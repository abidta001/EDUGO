package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

type User struct {
	gorm.Model
	Name     string `gorm:"not null" validate:"required,min=3,max=50"`
	Email    string `gorm:"unique;not null" validate:"required,email"`
	Phone    string `gorm:"unique;not null" validate:"required,len=10,numeric"`
	Password string `gorm:"not null" validate:"required,min=6"`
	Role     string `gorm:"not null"`
}

func (u *User) Validate() error {
	return validate.Struct(u)
}
