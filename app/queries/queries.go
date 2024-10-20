package queries

import "gorm.io/gorm"

type QueriesInterface interface {
	IsForeignKeyExisted(value interface{}, field string, table string) (bool, error)
}

type Queries struct {
	db *gorm.DB
}

func NewQueries(db *gorm.DB) *Queries {
	return &Queries{
		db: db,
	}
}
