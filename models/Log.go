package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	ID				uint				`gorm:"primaryKey;autoIncrement;"`
	Userid			uuid.UUID			`gorm:"column:user_id"`
	UserID			User				`gorm:"foreignKey:Userid;references:ID"`
	Type			string
	Description		string
}