package misc

import (
	"github.com/ecommerce/ports"
)

var (
	DatabaseIsHealthy = "Database is healthy"
	DatabaseIsOffline = "Database is offline, check for support"
)

type API interface {
	GetDatabaseStatus() string
}

type MiscService struct {
	miscRepository ports.MiscRepository
}

func NewMiscService(ms ports.MiscRepository) *MiscService {
	return &MiscService{
		miscRepository: ms,
	}
}
