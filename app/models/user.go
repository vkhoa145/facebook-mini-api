package models

import "time"

const TableNameUsers = "users"

type User struct {
	ID              uint       `gorm:"column:id;primaryKey;unique"`
	Name            string     `gorm:"column:name;type:varchar(255)" json:"name" validate:"required,min=3,max=255"`
	Email           string     `gorm:"column:email;unique;type:varchar(255);not null" json:"email" validate:"required,email,min=5,max=255"`
	Password        string     `gorm:"column:password;type:varchar(255);not null" json:"password" validate:"required"`
	Birthday        string     `gorm:"column:birthday;type:varchar(255);not null" json:birthday`
	Phone           string     `gorm:"column:phone;type:varchar(255)" json:phone`
	CreatedAt       time.Time  `gorm:"column:created_at"`
	UpdatedAt       time.Time  `gorm:column:updated_at`
	LoginToken      LoginToken `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	EmailVerifycode string     `gorm:"column:email_verifycode;type:varchar(255)" json:"email_verifycode"`
	PhoneVerifycode string     `gorm:"column:phone_verifycode;type:varchar(255)" json:"phone_verifycode"`
	VerifycatedAt   time.Time  `gorm:column:verifycated_at`
}

func (User) TableName() string {
	return TableNameUsers
}

type UserResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Birthday     string `json:"birthday"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignUpInput struct {
	Name       string `json:"name" validate:"required,min=3,max=255"`
	Email      string `json:"email" validate:"required,email,min=5,max=255"`
	BirthDay   int64  `json:"birthDay" validate:"gt=0"`
	BirthMonth int64  `json:"birthMonth" validate:"gt=0,lte=12"`
	BirthYear  int64  `json:"birthYear" validate:"gt=1900"`
	Password   string `json:"password" validate:"required"`
	Phone      string `json:"phone"`
}
