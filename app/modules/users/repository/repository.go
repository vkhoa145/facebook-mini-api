package repository

import "gorm.io/gorm"

type UserRepoInterface interface {
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}
