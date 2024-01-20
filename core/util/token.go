package util

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber"
)

var (
	ErrorMissingToken = errors.New("missing token")
)

func GetToken(ctx *fiber.Ctx, header string) (*string, error) {
	rawToken := strings.Replace(string(ctx.Fasthttp.Request.Header.Peek(header)), "Bearer ", "", 1)
	if rawToken == "" || rawToken == "null" {
		return nil, ErrorMissingToken
	}
	return &rawToken, nil
}
