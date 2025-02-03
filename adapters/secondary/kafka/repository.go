package kafka_ins

import (
	"context"
	"log"

	"github.com/ecommerce/core/domain/user"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
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

func (kr *KafkaRepository) WriteMessages(typ []byte, body []byte) error {
	err := kr.kw.WriteMessages(context.Background(), kafka.Message{
		Key:   typ,
		Value: body,
	})
	if err != nil {
		return err
	}

	return nil
}

// TODO: use a function to handle logic
func (kr *KafkaRepository) ExecuteMessageReceived(ps *gorm.DB) {
	for {
		msg, errR := kr.kr.ReadMessage(context.Background())
		if errR != nil {
			log.Println(errR)
			return
		}

		switch string(msg.Key) {
		case "welcomeEmailSended":
			var us user.User
			us.WelcomeEmailSended = true
			usEmail := string(msg.Value)
			ps.Where("email = ?", usEmail).Updates(&us)
			log.Println("user welcome email received: ", usEmail)
		default:
			log.Println("default kafka message")
		}
	}
}
