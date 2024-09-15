package models

import "gorm.io/gorm"

const TableNameSchemaMigration = "schema_migrations"

type SchemaMigration struct {
	gorm.Model
	Version string `gorm:"type:varchar(255)"`
}

func (SchemaMigration) TableName() string {
	return TableNameSchemaMigration
}
