package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

//User model struct
type User struct {
	gorm.Model
	ID			uuid.UUID			`gorm:"type:uuid"`
	Name			string
	Email			string			
	Address		string
	ImageKey		string
	ImageURL		string			
	Phone			string
	Password		string			
	FacebookID		string			
	Birthday		string			
	Gender		string
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		gorm.DeletedAt
	Favorite		[]Favorites			`gorm:"foreignKey:UserID"`
	Order			[]Order			`gorm:"foreignKey:Userid"`
}