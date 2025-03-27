package kafka_ins

import (
	"context"
	"log"

	"github.com/ecommerce/adapters/secondary/postgres"
	"github.com/segmentio/kafka-go"
)

type KafkaRepository struct {
	kw *kafka.Writer
	kr *kafka.Reader
}

func NewKafkaSessionRepository(writer *kafka.Writer, reader *kafka.Reader) *KafkaRepository {
	return &KafkaRepository{
		kw: writer,
		kr: reader,
	}
}

func (kr *KafkaRepository) WriteMessages(typ string, body []byte) error {
	err := kr.kw.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(typ),
		Value: body,
	})
	if err != nil {
		return err
	}

	return nil
}

// TODO: try to send email again, if possible
func (kr *KafkaRepository) ExecuteMessageReceived(ur *postgres.UserRepository) {
	for {
		msg, errR := kr.kr.ReadMessage(context.Background())
		if errR != nil {
			log.Println(errR)
			return
		}

		switch string(msg.Key) {
		case "welcomeEmailSended":
			mail := string(msg.Value)
			ur.WelcomeEmailReceived(mail)
			log.Println("user welcome email received: ", mail)
		default:
			log.Println("default kafka message")
		}
	}
}
