package logs

import (
	"context"
	"errors"

	"github.com/ecommerce/ports"
	"github.com/gofrs/uuid"
)

var (
	ErrorProductDoesntExist = errors.New("product doenst exist with this id")
)

type API interface {
	AddLog(context context.Context, typ string, description string, userId uuid.UUID) error
}

type LogService struct {
	logRepository ports.LogRepository
}

func NewLogService(lg ports.LogRepository) *LogService {
	return &LogService{
		logRepository: lg,
	}
}
