package models

import "time"

const TableNameVerificationCodes = "verification_codes"

type VerifycationType string

const (
	VerifycationTypeEmail VerifycationType = "email"
	VerifycationTypePhone VerifycationType = "phone"
)

type VerificationCode struct {
	ID               uint             `gorm:"column:id;primaryKey;unique"`
	UserID           uint             `gorm:"not null;unique"`
	VerifycationType VerifycationType `gorm:"column:verifycation_type;type:varchar(6);not null" json:"verifycation_type" validate:"required,min=6,max6"`
	Email            string           `gorm:"column:email;unique;type:varchar(255);not null" json:"email" validate:"required,email,min=5,max=255"`
	ExpiredAt        time.Time        `gorm:"column:expired_at"`
	NumsTry          int              `gorm:"nums_try;type:int" json:"nums_try" validate:"required"`
	IsLock           bool             `gorm:"column:is_lock;type:bool" json:"is_lock"`
	CreatedAt        time.Time        `gorm:"column:created_at"`
	UpdatedAt        time.Time        `gorm:column:updated_at`
	Code             string           `gorm:"code,type:varchar(6);not null" json:"code"`
}

func (VerificationCode) TableName() string {
	return TableNameVerificationCodes
}
