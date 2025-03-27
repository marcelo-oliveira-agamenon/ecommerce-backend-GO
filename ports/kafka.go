package ports

import "github.com/ecommerce/adapters/secondary/postgres"

type KafkaService interface {
	WriteMessages(typ string, body []byte) error
	ExecuteMessageReceived(ur *postgres.UserRepository)
}
