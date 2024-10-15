package repository

import (
	"github.com/vkhoa145/facebook-mini-api/app/models"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	CheckExistedEmail(email string) bool
	CreateUser(payload *models.User, tx *gorm.DB) (*models.User, error)
	CreateLoginToken(jwt *models.JwtResponse, tx *gorm.DB) (*models.JwtResponse, error)
}

type UserRepo struct {
	DB *gorm.DB
	Tx *gorm.DB
}

func NewUserRepo(db *gorm.DB, tx *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
		Tx: tx,
	}
}
