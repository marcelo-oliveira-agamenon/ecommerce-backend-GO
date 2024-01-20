package ports

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

type TokenService interface {
	CreateToken(data string) (*string, time.Time, error)
	VerifyToken(token string) error
	ClaimTokenData(token string) (*Claims, error)
}
