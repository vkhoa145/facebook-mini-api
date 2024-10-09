package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/vkhoa145/facebook-mini-api/app/models"
	"github.com/vkhoa145/facebook-mini-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb(cfg *config.Config) *gorm.DB {
	if err := createDatabaseIfNotExists(cfg); err != nil {
		log.Fatalf("Failed to create or verify database: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.Port,
		cfg.DB.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(
		&models.User{},
		&models.SchemaMigration{},
		&models.LoginToken{},
	)

	initSchemaMigration(db)
	runMigration(db)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func createDatabaseIfNotExists(cfg *config.Config) error {
	// Connect to the default 'postgres' database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable TimeZone=%s",
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Port,
		cfg.DB.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Check if the database exists
	var exists bool
	err = db.Raw("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = ?)", cfg.DB.Name).Scan(&exists).Error
	if err != nil {
		return fmt.Errorf("failed to check if database exists: %v", err)
	}

	// Create the database if it doesn't exist
	if !exists {
		if err := db.Exec("CREATE DATABASE " + cfg.DB.Name).Error; err != nil {
			return fmt.Errorf("failed to create database: %v", err)
		}
		fmt.Printf("Database %s created successfully.\n", cfg.DB.Name)
	} else {
		fmt.Printf("Database %s already exists.\n", cfg.DB.Name)
	}

	return nil
}

func initSchemaMigration(db *gorm.DB) {
	sqlFile, err := os.ReadFile(filepath.Join("sql", "schema_migration.sql"))
	if err != nil {
		log.Fatal("Failed to read SQL file:", err)
	}

	fmt.Println("Initializing schema_migrations table")
	if err := db.Exec(string(sqlFile)).Error; err != nil {
		log.Fatal("Failed to initialize schema migrations table:", err)
	}
}

func runMigration(db *gorm.DB) {
	files, err := ioutil.ReadDir("sql/migrations")
	if err != nil {
		log.Fatal("Failed to read migration files:", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			var count int64
			version := file.Name()

			db.Raw("SELECT COUNT(*) FROM schema_migrations WHERE version = ?", version).Scan(&count)
			if count > 0 {
				fmt.Printf("Skipping migration: %s (already applied)\n", version)
				continue
			}

			sqlFile, err := os.ReadFile(filepath.Join("sql/migrations", version))
			if err != nil {
				log.Fatal("Failed to read SQL file:", err)
			}

			fmt.Printf("Applying migration: %s\n", version)
			if err := db.Exec(string(sqlFile)).Error; err != nil {
				log.Fatal("Failed to execute SQL migration:", err)
			}

			schema_version := &models.SchemaMigration{Version: version}
			if err := db.Table(models.SchemaMigration{}.TableName()).Create(schema_version).Error; err != nil {
				log.Fatal("Failed to insert migration version into schema_migrations table")
			}
		}
	}
}
