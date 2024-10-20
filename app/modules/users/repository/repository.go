package repository

import (
	"github.com/vkhoa145/facebook-mini-api/app/models"
	"github.com/vkhoa145/facebook-mini-api/app/queries"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	CheckExistedEmail(email string) bool
	CreateUser(payload *models.User, tx *gorm.DB) (*models.User, error)
	CreateLoginToken(jwt *models.JwtResponse, UserID uint, tx *gorm.DB) (*models.JwtResponse, error)
}

type UserRepo struct {
	DB      *gorm.DB
	Queries *queries.Queries
}

func NewUserRepo(db *gorm.DB, queries *queries.Queries) *UserRepo {
	return &UserRepo{
		DB: db,
		Queries: queries,
	}
}
