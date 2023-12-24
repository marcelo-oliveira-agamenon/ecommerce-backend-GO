package postgres

import (
	"os"

	"github.com/ecommerce/core/domain/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database connection postgres
func initDatabase() (*gorm.DB, error) {
	godotenv.Load(".env")
	dbString := os.Getenv("DB_CONNECTION")

	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{
		Logger:logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&user.User{})

	return db, nil
}

func NewPostgresRepository() (*gorm.DB, error) {
	db, err := initDatabase()
	if err != nil {
		return nil, err
	}
	
	return db, nil
}