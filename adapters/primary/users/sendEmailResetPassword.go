package users

import (
	"errors"

	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
)

var (
	EmailSuccessMessage = "Email sended"
	EmailFileName       = "template/resetPassword.html"
	EmailSubject        = "Redefinição de Senha - Código de Verificação"
	ErrorMissingEmail   = errors.New("missing email parameter")
)

func SendEmailResetPassword(userAPI users.API, email ports.EmailService) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		userEmail := ctx.Query("email")
		if len(userEmail) == 0 {
			ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorMissingEmail,
			})
			return
		}

		template, err := userAPI.SendEmailResetPassword(ctx.Context(), userEmail)
		if err != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
			return
		}

		_, errMail := email.SendEmail(userEmail, EmailFileName, template, EmailSubject)
		if errMail != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": ErrorMissingEmail,
			})
			return
		}

		ctx.Status(200).JSON(EmailSuccessMessage)
	}
}
