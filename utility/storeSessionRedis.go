package utility

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ecommerce/db"
	"github.com/gofiber/fiber"
)

var ctx = context.Background()

type redisStore struct {
	UserId	string
	AccessAt	string
	ExpiresAt	string
}

func StoreSessionRedis(w *fiber.Ctx, userData string, expTime string) {
	store := redisStore{
		UserId: userData,
		AccessAt: time.Now().String(),
		ExpiresAt: expTime,
	}
	marshalStore, err := json.Marshal(store)
	if err != nil {
		w.Status(500).JSON("Error marshall redis store")
		return
	}
	errRedis := db.RedisServer.Set(ctx, userData, marshalStore, 0).Err()
	if errRedis != nil {
		fmt.Print(errRedis)
		w.Status(500).JSON("Error in redis store session")
		return
	}
}