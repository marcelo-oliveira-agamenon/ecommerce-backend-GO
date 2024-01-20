package ports

import (
	"time"
)

type TokenService interface {
	CreateToken(data string) (*string, time.Time, error)
	VerifyToken(token string) error
}
