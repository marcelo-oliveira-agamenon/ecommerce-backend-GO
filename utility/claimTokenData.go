package utility

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

type claims struct {
	UserId	string	`json:"userId"`
	jwt.StandardClaims
}

//Get data from claim of JWT token
func ClaimTokenData(w *fiber.Ctx) claims  {
	token := strings.Replace(string(w.Fasthttp.Request.Header.Peek("Authorization")), "Bearer ", "", 1)
	jwtSecret := []byte(GetDotEnv("JWT_KEY"))
	decodeClaim := &claims{}

	_, err := jwt.ParseWithClaims(token, decodeClaim, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		w.Status(500).JSON("Error token decode")
		return claims{}
	}

	return *decodeClaim
}