package repository

import (
	"github.com/vkhoa145/facebook-mini-api/app/models"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	CheckExistedEmail(email string) bool
	CreateUser(payload *models.User, tx *gorm.DB) (*models.User, error)
	CreateLoginToken(jwt *models.JwtResponse, UserID uint, tx *gorm.DB) (*models.JwtResponse, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}
