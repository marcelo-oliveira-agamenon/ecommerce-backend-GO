package users

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber/v2"
)

var (
	EmailSuccessMessage = "Please, check your inbox"
	ErrorMissingEmail   = errors.New("missing email parameter")
)

func SendEmailResetPassword(userAPI users.API, kafka ports.KafkaService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userEmail := ctx.Query("email")
		if len(userEmail) == 0 {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorMissingEmail,
			})
		}

		template, err := userAPI.SendEmailResetPassword(ctx.Context(), userEmail)
		if err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		body, errM := json.Marshal(template)
		if errM == nil {
			errK := kafka.WriteMessages("resetPassword", body)
			if errK != nil {
				log.Println("kafka message", errK)
			}
		} else {
			log.Println("marshall message", errM)
		}

		return ctx.Status(200).JSON(EmailSuccessMessage)
	}
}
