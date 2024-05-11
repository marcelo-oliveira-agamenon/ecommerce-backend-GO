package ports

type KafkaService interface {
	WriteMessages(key []byte, body []byte) error
}
