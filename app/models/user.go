package models

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUsers = "users"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"type:varchar(255)" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	Birthday string `gorm:"type:varchar(255)" json:birthday`
	Phone    string `gorm:"type:varchar(255)" json:phone`
}

func (User) TableName() string {
	return TableNameUsers
}

type UserResponse struct {
	ID        uint      `json:"id,omitemtpty"`
	Name      string    `json:"name,omitemtpty"`
	Email     string    `json:"email,omitemtpty"`
	Birthday  string    `json:"birthday,omitemtpty"`
	Phone     string    `json:"phone,omitemtpty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SignUpInput struct {
	Name       string `json:"name" validate:"required,min=3,max=255"`
	Email      string `json:"email" validate:"required"`
	BirthDay   int64  `json:"birth_day" validate:"required"`
	BirthMonth int64  `json:"birth_month" validate:"required"`
	BirthYear  int64  `json:"birth_year" validate:"required"`
}
