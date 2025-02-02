package kafka_ins

import (
	"os"

	"github.com/segmentio/kafka-go"
)

var (
	address     = os.Getenv("KAFKA_ADDRESS")
	order_topic = os.Getenv("KAFKA_ORDER_TOPIC")
	user_topic  = os.Getenv("KAFKA_USER_TOPIC")
)

func initKafkaInstance() (*kafka.Writer, *kafka.Reader, error) {
	kafW := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{address},
		Topic:    order_topic,
		Balancer: &kafka.LeastBytes{},
	})

	kafR := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{address},
		Topic:       user_topic,
		Partition:   0,
		StartOffset: kafka.FirstOffset,
	})

	return kafW, kafR, nil
}

func NewKafkaRepository() (*kafka.Writer, *kafka.Reader, error) {
	kafW, kafR, err := initKafkaInstance()
	if err != nil {
		return nil, nil, err
	}

	return kafW, kafR, nil
}
