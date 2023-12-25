package users

import (
	"github.com/ecommerce/core/domain/user"
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber"
	"github.com/lib/pq"
)

func SignUp(userAPI users.API, token ports.TokenService, storage ports.StorageService) fiber.Handler {
	return func(ctx *fiber.Ctx) {
		if err := ctx.BodyParser(&user.User{}); err != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
			return
		}

		UserResponse, err := userAPI.SignUp(ctx.Context(), user.User{
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
			ctx.Status(400).JSON(&fiber.Map{
				"error": err.Error(),
			})
			return
		}
		//todo: change to a transaction
		if ava, _ := ctx.FormFile("avatar"); ava != nil {
			file, err := ava.Open()
			if err != nil {
				userAPI.DeleteUser(ctx.Context(), UserResponse.ID.String())
				ctx.Status(500).JSON(&fiber.Map{
					"error": err.Error(),
				})
				return
			}
			resp, errSto := storage.SaveFileAWS(file, ava.Filename, ava.Size, "user")
			if errSto != nil {
				userAPI.DeleteUser(ctx.Context(), UserResponse.ID.String())
				ctx.Status(500).JSON(&fiber.Map{
					"error": errSto.Error(),
				})
				return
			}
			userAPI.UpdateUser(ctx.Context(), UserResponse.ID.String(), user.User{
				ImageKey: resp.ImageKey,
				ImageURL: resp.ImageURL,
			})
		}

		token, _, errToken := token.CreateToken(UserResponse.ID.String())
		if errToken != nil {
			ctx.Status(500).JSON(&fiber.Map{
				"error": errToken.Error(),
			})
			return
		}

		ctx.Status(201).JSON(&fiber.Map{
			"user":  UserResponse,
			"token": token,
		})
	}
}
