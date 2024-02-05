package postgres

import (
	"os"

	"github.com/ecommerce/core/domain/product"
	"github.com/ecommerce/core/domain/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database connection postgres
func initDatabase() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_CONNECTION_HOST")
	dbUser := os.Getenv("DB_CONNECTION_USER")
	dbPassword := os.Getenv("DB_CONNECTION_PASSWORD")
	dbName := os.Getenv("DB_CONNECTION_DBNAME")
	dbPort := os.Getenv("DB_CONNECTION_PORT")
	dbConn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&product.Product{})

	return db, nil
}

func NewPostgresRepository() (*gorm.DB, error) {
	db, err := initDatabase()
	if err != nil {
		return nil, err
	}

	return db, nil
}
