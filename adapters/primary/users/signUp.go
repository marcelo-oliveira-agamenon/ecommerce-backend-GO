package users

import (
	"encoding/json"
	"log"

	"github.com/ecommerce/core/domain/user"
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func SignUp(userAPI users.API, token ports.TokenService, storage ports.StorageService, kafka ports.KafkaService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if err := ctx.BodyParser(&user.User{}); err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		//TODO: create universal error struct with status code and text
		usrRes, err := userAPI.SignUp(ctx.Context(), user.User{
			Name:       ctx.FormValue("name"),
			Email:      ctx.FormValue("email"),
			Address:    ctx.FormValue("address"),
			Phone:      ctx.FormValue("phone"),
			Password:   ctx.FormValue("password"),
			FacebookID: ctx.FormValue("facebookID"),
			Birthday:   ctx.FormValue("birthday"),
			Gender:     ctx.FormValue("gender"),
			Roles:      pq.StringArray{"user"},
		})
		if err != nil {
			return ctx.Status(400).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		if ava, _ := ctx.FormFile("avatar"); ava != nil {
			file, err := ava.Open()
			if err != nil {
				userAPI.DeleteUser(ctx.Context(), usrRes.ID.String())
				return ctx.Status(500).JSON(&fiber.Map{
					"error": err.Error(),
				})

			}
			resp, errSto := storage.SaveFileAWS(file, ava.Filename, ava.Size, "user")
			if errSto != nil {
				userAPI.DeleteUser(ctx.Context(), usrRes.ID.String())
				return ctx.Status(500).JSON(&fiber.Map{
					"error": errSto.Error(),
				})

			}
			userAPI.UpdateUser(ctx.Context(), usrRes.ID.String(), user.User{
				ImageKey: resp.ImageKey,
				ImageURL: resp.ImageURL,
			})
		}

		token, _, errToken := token.CreateToken(usrRes.ID.String())
		if errToken != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errToken.Error(),
			})
		}

		body, errM := json.Marshal(usrRes)
		if errM == nil {
			errK := kafka.WriteMessages("signup", body)
			if errK != nil {
				log.Println("kafka message", errK)
			}
		} else {
			log.Println("marshall message", errM)
		}

		return ctx.Status(201).JSON(&fiber.Map{
			"user":  usrRes,
			"token": token,
		})
	}
}
