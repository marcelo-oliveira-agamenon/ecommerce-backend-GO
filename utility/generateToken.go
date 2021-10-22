package utility

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

func GenerateToken(w *fiber.Ctx, userData string) (token string, expiredAt time.Time) {
	expTime := time.Now().Add(60 * time.Minute)
	claimsJwt := &claims{
		UserId: userData,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	tokenMethod := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsJwt)
	jwtKey := []byte(GetDotEnv("JWT_KEY"))
	token, err := tokenMethod.SignedString(jwtKey)
	if err != nil {
		w.Status(500).JSON("Error in jwt token")
		return "", time.Time{}
	}

	return token, expTime
}