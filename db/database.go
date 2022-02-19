package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	u "github.com/ecommerce/models"
)

//DBConn connection
var (
	DBConn	*gorm.DB
)

// CreateConnection with gorm
func CreateConnection()  {
	godotenv.Load(".env")
	dbString := os.Getenv("DB_CONNECTION")

	dsn := dbString
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Print(err)
		return
	}
	DBConn = db

	db.AutoMigrate(&u.Category{})
	db.AutoMigrate(&u.Product{})
	db.AutoMigrate(&u.ProductImage{})
	db.AutoMigrate(&u.User{})
	db.AutoMigrate(&u.Order{})	
	db.AutoMigrate(&u.Favorites{})
	db.AutoMigrate(&u.Payment{})
	db.AutoMigrate(&u.Coupon{})
	db.AutoMigrate(&u.Log{})
}