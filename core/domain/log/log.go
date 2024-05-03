package logs

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey;autoIncrement;"`
	Userid      uuid.UUID `gorm:"column:user_id"`
	Type        string
	Description string
}

func NewLog(data Log) (Log, error) {
	return Log{
		Userid:      data.Userid,
		Type:        data.Type,
		Description: data.Description,
	}, nil
}
