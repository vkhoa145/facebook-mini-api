package db

import (
	"fmt"

	"github.com/vkhoa145/facebook-mini-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s sslmode=disable TimeZone=%s",
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.Port,
		cfg.DB.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
