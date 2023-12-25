package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ecommerce/ports"
)

var (
	ErrorToken           = errors.New("invalid generated token")
	ErrorInvalidPassword = errors.New("invalid password")
)

type claims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

type JWTToken struct {
	jwyKey string
}

func NewToken(jwtKey string) ports.TokenService {
	return &JWTToken{
		jwyKey: jwtKey,
	}
}

func (jt *JWTToken) CreateToken(userID string) (*string, time.Time, error) {
	expTime := time.Now().Add(60 * time.Minute)
	claimsJwt := &claims{
		UserId: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	tokenMethod := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsJwt)
	token, err := tokenMethod.SignedString([]byte(jt.jwyKey))
	if err != nil {
		return nil, time.Time{}, ErrorToken
	}

	return &token, expTime, nil
}

func (jt *JWTToken) VerifyToken() error {
	return nil
}
