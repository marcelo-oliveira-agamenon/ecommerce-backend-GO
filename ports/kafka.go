package ports

type KafkaService interface {
	WriteMessages(body []byte) error
}
