package users

import (
	"errors"

	"github.com/ecommerce/core/services/users"
	"github.com/gofiber/fiber/v2"
)

var (
	EmailSuccessMessage = "Email sended"
	EmailFileName       = "template/resetPassword.html"
	EmailSubject        = "Redefinição de Senha - Código de Verificação"
	ErrorMissingEmail   = errors.New("missing email parameter")
)

func SendEmailResetPassword(userAPI users.API) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userEmail := ctx.Query("email")
		if len(userEmail) == 0 {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorMissingEmail,
			})
		}

		// TODO: create kakfa topic to this
		// template, err := userAPI.SendEmailResetPassword(ctx.Context(), userEmail)
		// if err != nil {
		// 	return ctx.Status(500).JSON(&fiber.Map{
		// 		"error": err.Error(),
		// 	})
		// }

		// _, errMail := email.SendEmail(userEmail, EmailFileName, template, EmailSubject)
		// if errMail != nil {
		// 	return ctx.Status(500).JSON(&fiber.Map{
		// 		"error": ErrorMissingEmail,
		// 	})
		// }

		return ctx.Status(200).JSON(EmailSuccessMessage)
	}
}
