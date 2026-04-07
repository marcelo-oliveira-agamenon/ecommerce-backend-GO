package postgres

import (
	"log"
	"os"
	"time"

	"github.com/ecommerce/adapters/secondary/postgres/seeds"
	"github.com/ecommerce/core/domain/category"
	"github.com/ecommerce/core/domain/coupon"
	"github.com/ecommerce/core/domain/favorite"
	logs "github.com/ecommerce/core/domain/log"
	"github.com/ecommerce/core/domain/order"
	"github.com/ecommerce/core/domain/orderDetails"
	"github.com/ecommerce/core/domain/payment"
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

type QueryParamsUsers struct {
	CreatedAtStart string
	CreatedAtEnd   string
	Gender         string
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

	sqlDb, err1 := db.DB()
	if err1 != nil {
		return nil, err
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Second * 30)

	if err := db.AutoMigrate(&user.User{}); err != nil {
		log.Printf("Migration failed for User migration: %v", err)
	}
	if err := db.AutoMigrate(&product.Product{}); err != nil {
		log.Printf("Migration failed for Product migration: %v", err)
	}
	if err := db.AutoMigrate(&category.Category{}); err != nil {
		log.Printf("Migration failed for Category migration: %v", err)
	}
	if err := db.AutoMigrate(&productImage.ProductImage{}); err != nil {
		log.Printf("Migration failed for ProductImage migration: %v", err)
	}
	if err := db.AutoMigrate(&favorite.Favorite{}); err != nil {
		log.Printf("Migration failed for Favorite migration: %v", err)
	}
	if err := db.AutoMigrate(&coupon.Coupon{}); err != nil {
		log.Printf("Migration failed for Coupon migration: %v", err)
	}
	if err := db.AutoMigrate(&order.Order{}); err != nil {
		log.Printf("Migration failed for Order migration: %v", err)
	}
	if err := db.AutoMigrate(&orderDetails.OrderDetails{}); err != nil {
		log.Printf("Migration failed for OrderDetails migration: %v", err)
	}
	if err := db.AutoMigrate(&payment.Payment{}); err != nil {
		log.Printf("Migration failed for Payment migration: %v", err)
	}
	if err := db.AutoMigrate(&logs.Log{}); err != nil {
		log.Printf("Migration failed for Log migration: %v", err)
	}

	// Populate database with products, categories
	errSe := seeds.SeedInitialData(db)
	if errSe != nil {
		log.Println("Error in seeding database")
	}

	return db, nil
}

func NewPostgresRepository() (*gorm.DB, error) {
	db, err := initDatabase()
	if err != nil {
		return nil, err
	}

	return db, nil
}
