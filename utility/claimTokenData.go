package utility

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

type claims struct {
	userId	string
	jwt.StandardClaims
}

//Get data from claim of JWT token
func ClaimTokenData(w *fiber.Ctx) (claims, string)  {
	token := strings.Replace(string(w.Fasthttp.Request.Header.Peek("Authorization")), "Bearer ", "", 1)
	jwtSecret := []byte(GetDotEnv("JWT_KEY"))

	localToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Error in token verify")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return claims{}, "Error parsing token"
	}

	cl := localToken.Claims.(jwt.MapClaims)
	user := cl["userId"].(string)
	
	fmt.Print("aa ",user)

	return claims{}, ""
}