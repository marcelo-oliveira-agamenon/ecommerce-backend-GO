package orders

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/ecommerce/core/domain/orderDetails"
	orders "github.com/ecommerce/core/services/order"
	ordersdetails "github.com/ecommerce/core/services/ordersDetails"
	"github.com/ecommerce/core/services/users"
	"github.com/ecommerce/ports"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorInvalidOrderData = errors.New("invalid order data")
)

type MapOrderEmail struct {
	Quantity    int
	Value       float64
	UserName    string
	OrderNumber string
	Email       string
}

func CreateOrder(orderAPI orders.API, userAPI users.API, orderDetailsAPI ordersdetails.API, kafka ports.KafkaService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		dec := ctx.Locals("user").(*ports.Claims)

		user, errU := userAPI.GetUserById(ctx.Context(), dec.UserId)
		if errU != nil {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": errU.Error(),
			})
		}

		var orDe []orderDetails.OrderProductData
		errPr := ctx.BodyParser(&orDe)
		if errPr != nil {
			return ctx.Status(422).JSON(&fiber.Map{
				"error": ErrorInvalidOrderData.Error(),
			})
		}

		var ordSli []orderDetails.OrderDetails
		for _, val := range orDe {
			orFor, errF := orderDetailsAPI.CheckOrderDetails(ctx.Context(), "placeholder", val)
			if errF != nil {
				return ctx.Status(422).JSON(&fiber.Map{
					"error": errF.Error(),
				})
			}

			ordSli = append(ordSli, *orFor)
		}

		toV, qtd, errDe := orderDetailsAPI.GetTotalOrderValueAndQuatity(ctx.Context(), orDe)
		if errDe != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": errDe.Error(),
			})
		}

		newO, err := orderAPI.AddOrder(ctx.Context(), dec.UserId, *qtd, *toV, ordSli, orDe[0].CouponUsed)
		if err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		resp := MapOrderEmail{
			Quantity:    newO.TotalQtd,
			Value:       newO.TotalValue,
			UserName:    user.Name,
			OrderNumber: newO.ID,
			Email:       user.Email,
		}
		body, errM := json.Marshal(resp)
		if errM == nil {
			errK := kafka.WriteMessages("newOrder", body)
			if errK != nil {
				log.Println("kafka message", errK)
			}
		} else {
			log.Println("marshall message", errM)
		}

		return ctx.Status(201).JSON(newO)
	}
}
