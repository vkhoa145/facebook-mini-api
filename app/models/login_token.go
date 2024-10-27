package models

import "time"

const TableNameLoginTokens = "login_tokens"

type LoginToken struct {
	ID           uint      `gorm:"column:id;primaryKey;unique"`
	UserID       uint      `gorm:"not null;unique"`
	RefreshToken string    `gorm:"type:varchar(255);not null" validate:"required,min=3"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:column:updated_at`
}

func (LoginToken) TableName() string {
	return TableNameLoginTokens
}
