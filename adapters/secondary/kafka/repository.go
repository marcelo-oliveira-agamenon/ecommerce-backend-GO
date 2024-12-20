package kafka_ins

import (
	"github.com/segmentio/kafka-go"
)

type KafkaRepository struct {
	kf *kafka.Conn
}

func NewKafkaSessionRepository(conn *kafka.Conn) *KafkaRepository {
	return &KafkaRepository{
		kf: conn,
	}
}

func (kr *KafkaRepository) WriteMessages(typ []byte, body []byte) error {
	_, err := kr.kf.WriteMessages(kafka.Message{
		Key:   typ,
		Value: body,
	})
	if err != nil {
		return err
	}

	return nil
}
