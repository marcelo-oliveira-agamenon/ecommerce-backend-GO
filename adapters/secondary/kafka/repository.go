package kafka_ins

import (
	"errors"

	"github.com/segmentio/kafka-go"
)

var (
	ErrorTopic           = errors.New("failed kafka topic")
	ErrorGenerateMessage = errors.New("failed message")
	ErrorCloseMessage    = errors.New("failed close kafka connection")
)

type KafkaRepository struct {
	kf *kafka.Conn
}

func NewKafkaSessionRepository(conn *kafka.Conn) *KafkaRepository {
	return &KafkaRepository{
		kf: conn,
	}
}

func (kr *KafkaRepository) WriteMessages(key []byte, body []byte) error {
	_, err := kr.kf.WriteMessages(kafka.Message{
		Key:   key,
		Value: body,
	})
	if err != nil {
		return ErrorGenerateMessage
	}

	if errC := kr.kf.Close(); errC != nil {
		return ErrorCloseMessage
	}

	return nil
}
