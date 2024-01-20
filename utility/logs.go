package utility

import (
	"github.com/ecommerce/db"
	"github.com/ecommerce/models"
	"github.com/gofrs/uuid"
)

// Insert a type of registry into db
func InsertLogRegistryIntoDabatase(typeOfLog string, description string, userId string) error {
	var log models.Log
	log.Userid, _ = uuid.FromString(userId)
	log.Type = typeOfLog
	log.Description = description

	result := db.DBConn.Create(&log)
	if result.Error != nil {
		return result.Error
	}

	return nil
}