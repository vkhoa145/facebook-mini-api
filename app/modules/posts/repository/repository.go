package repository

import (
	"github.com/vkhoa145/facebook-mini-api/app/queries"
	"gorm.io/gorm"
)

type PostRepoInterface interface {
	GetPosts()
}

type PostRepo struct {
	DB      *gorm.DB
	Queries *queries.Queries
}

func NewPostRepo(db *gorm.DB, queries *queries.Queries) *PostRepo {
	return &PostRepo{
		DB:      db,
		Queries: queries,
	}
}
