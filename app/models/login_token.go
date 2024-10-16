package models

import "gorm.io/gorm"

const TableNameLoginTokens = "login_tokens"

type LoginToken struct {
	gorm.Model
	UserID       uint   `gorm:"not null;unique"`
	RefreshToken string `gorm:"type:varchar(255);not null" validate:"required,min=3"`
}

func (LoginToken) TableName() string {
	return TableNameLoginTokens
}
