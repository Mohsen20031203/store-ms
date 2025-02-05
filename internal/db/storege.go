package db

import (
	"fmt"
	"hello/config"
	"hello/internal/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Storege struct {
	DB *gorm.DB
}

func NewStorege(config config.Config) (*Storege, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.PgConfig.Host, config.PgConfig.User, config.PgConfig.Password, config.PgConfig.Database, config.PgConfig.Port)

	database, err := gorm.Open(postgres.Open(dsn),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		})

	database.AutoMigrate(
		&models.User{},
		&models.Brand{},
		&models.Category{},
		&models.Order{},
		&models.Payment{},
		&models.Product{},
		&models.Review{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return &Storege{
		DB: database,
	}, nil
}
