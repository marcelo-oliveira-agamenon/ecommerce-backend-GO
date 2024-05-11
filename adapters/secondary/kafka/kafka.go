package kafka_ins

import (
	"context"
	"os"

	"github.com/segmentio/kafka-go"
)

func initKafkaInstance() (*kafka.Conn, error) {
	address := os.Getenv("KAFKA_ADDRESS")
	network := os.Getenv("KAFKA_NETWORK")
	topic := os.Getenv("KAFKA_TOPIC")

	kfCon, errK := kafka.DialLeader(context.Background(), network, address, topic, 0)
	if errK != nil {
		return nil, errK
	}

	return kfCon, nil
}

func NewKafkaRepository() (*kafka.Conn, error) {
	kf, err := initKafkaInstance()
	if err != nil {
		return nil, err
	}

	return kf, nil
}
