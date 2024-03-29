package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ecommerce/ports"
)

var (
	ErrorToken           = errors.New("invalid generated token")
	ErrorInvalidPassword = errors.New("invalid password")
	ErrorInvalidToken    = errors.New("invalid token")
	ErrorParseToken      = errors.New("parse token")
)

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
	claimsJwt := &ports.Claims{
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

func (jt *JWTToken) VerifyToken(token string) error {
	jwtKey := []byte(os.Getenv("JWT_KEY"))
	hasToken, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrorParseToken
		}
		return jwtKey, nil
	})

	if hasToken.Valid {
		return nil
	}
	return ErrorInvalidToken
}

func (jt *JWTToken) ClaimTokenData(token string) (*ports.Claims, error) {
	jwtSecret := []byte(os.Getenv("JWT_KEY"))
	decodeClaim := &ports.Claims{}

	_, err := jwt.ParseWithClaims(token, decodeClaim, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, ErrorParseToken
	}

	return decodeClaim, nil
}
