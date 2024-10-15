package models

import "gorm.io/gorm"

const TableNameLoginTokens = "login_tokens"

type LoginToken struct {
	gorm.Model
	UserId       uint
	RefreshToken string `gorm:"type:varchar(255)" validate:"required,min=3"`
}

func (LoginToken) TableName() string {
	return TableNameLoginTokens
}
