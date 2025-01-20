package postgres

import (
	"gorm.io/gorm"
)

type MiscRepository struct {
	db *gorm.DB
}

func NewMiscRepository(dbConn *gorm.DB) *MiscRepository {
	return &MiscRepository{
		db: dbConn,
	}
}

func (ms *MiscRepository) GetDatabaseStatus() bool {
	dbC, errA := ms.db.DB()
	if errA != nil {
		return false
	}

	if err := dbC.Ping(); err != nil {
		return false
	}

	return true
}
