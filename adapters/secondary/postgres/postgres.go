package postgres

import (
	"os"

	"github.com/ecommerce/core/domain/category"
	"github.com/ecommerce/core/domain/favorite"
	"github.com/ecommerce/core/domain/product"
	"github.com/ecommerce/core/domain/productImage"
	"github.com/ecommerce/core/domain/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type QueryParamsProduct struct {
	Limit          int
	Offset         int
	GetByCategory  string
	GetByPromotion string
	GetRecentOnes  string
	GetByName      string
}

type QueryParams struct {
	Limit  int
	Offset int
}

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
	db.AutoMigrate(&category.Category{})
	db.AutoMigrate(&productImage.ProductImage{})
	db.AutoMigrate(&favorite.Favorite{})

	return db, nil
}

func NewPostgresRepository() (*gorm.DB, error) {
	db, err := initDatabase()
	if err != nil {
		return nil, err
	}

	return db, nil
}
