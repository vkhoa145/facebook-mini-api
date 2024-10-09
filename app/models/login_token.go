package models

import "gorm.io/gorm"

const TableNameLoginTokens = "login_tokens"

type LoginToken struct {
	gorm.Model
	UserId       uint
	RefreshToken string
}

func (LoginToken) TableName() string {
	return TableNameLoginTokens
}
