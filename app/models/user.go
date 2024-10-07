package models

import (
	"gorm.io/gorm"
)

const TableNameUsers = "users"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255)" json:"name" validate:"required,min=3,max=255"`
	Email    string `gorm:"type:varchar(255)" json:"email" validate:"required,email,min=5,max=255"`
	Password string `gorm:"type:varchar(255)" json:"password" validate:"required"`
	Birthday string `gorm:"type:varchar(255)" json:birthday`
	Phone    string `gorm:"type:varchar(255)" json:phone`
}

func (User) TableName() string {
	return TableNameUsers
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignUpInput struct {
	Name       string `json:"name" validate:"required,min=3,max=255"`
	Email      string `json:"email" validate:"required,email,min=5,max=255"`
	BirthDay   int64  `json:"birth_day" validate:"gt=0"`
	BirthMonth int64  `json:"birth_month" validate:"gt=0,lte=12"`
	BirthYear  int64  `json:"birth_year" validate:"gt=1900"`
	Password   string `json:"password" validate:"required"`
}
