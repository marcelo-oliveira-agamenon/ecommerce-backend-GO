package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
		return
	}
	DBConn = db

	db.AutoMigrate(&u.User{}, &u.Product{}, &u.Category{}, &u.Favorites{}, &u.Order{})
}