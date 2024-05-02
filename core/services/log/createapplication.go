package logs

import (
	"context"

	logs "github.com/ecommerce/core/domain/log"
	"github.com/gofrs/uuid"
)

func (l *LogService) AddLog(context context.Context, typ string, description string, userId uuid.UUID) error {
	ty, errT := logs.NewType(typ)
	if errT != nil {
		return errT
	}

	de, errD := logs.NewDescription(description)
	if errD != nil {
		return errD
	}

	lo, errN := logs.NewLog(logs.Log{
		Userid:      userId,
		Type:        *ty,
		Description: *de,
	})
	if errN != nil {
		return errN
	}

	_, err := l.logRepository.AddLog(context, lo)
	if err != nil {
		return err
	}

	return nil
}
