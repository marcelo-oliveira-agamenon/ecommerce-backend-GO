package utility

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

//VerifyToken return bool for jwt token
func VerifyToken(c *fiber.Ctx)  {
	rawToken := strings.Replace(string(c.Fasthttp.Request.Header.Peek("Authorization")), "Bearer ", "", 1)

	if rawToken == "" || rawToken == "null" {
		c.Status(401).JSON("Missing token")
		return
	}

	jwtKey := []byte(GetDotEnv("JWT_KEY"))
	hasToken, _ := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Error in token verify")
		}
		return jwtKey, nil
	})

	if hasToken.Valid {
		c.Next()
	} else {
		c.Status(401).JSON("Invalid Token")
	}
}