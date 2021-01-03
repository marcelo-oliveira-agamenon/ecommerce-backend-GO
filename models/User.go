package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

//JSONB type postgres
type JSONB map[string]interface{}

//User model struct
type User struct {
	gorm.Model
	ID			uuid.UUID			`gorm:"type:uuid"`
	Name			string
	Email			string			
	Address		string
	Avatar		pq.StringArray		`gorm:"type:text[]"`			
	Phone			string
	Password		string			
	FacebookID		string			
	Birthday		string			
	Gender		string
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		gorm.DeletedAt
	Favorite		[]Favorites			`gorm:"foreignKey:UserID"`
	Order			[]Order			`gorm:"foreignKey:UserID"`
}