package ports

import (
	"context"

	logs "github.com/ecommerce/core/domain/log"
)

type LogRepository interface {
	AddLog(ctx context.Context, l logs.Log) (*logs.Log, error)
}
