package postgres

import (
	"context"

	logs "github.com/ecommerce/core/domain/log"
	"gorm.io/gorm"
)

type LogRepository struct {
	db *gorm.DB
}

func NewLogRepository(dbConn *gorm.DB) *LogRepository {
	return &LogRepository{
		db: dbConn,
	}
}

func (lg *LogRepository) AddLog(ctx context.Context, l logs.Log) (*logs.Log, error) {
	result := lg.db.Create(&l)
	if result.Error != nil {
		return nil, result.Error
	}

	return &l, nil
}
