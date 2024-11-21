package ports

type KafkaService interface {
	WriteMessages(typ []byte, body []byte) error
}
